package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/betaiota/grpchat/pkg/storage"
	pb "github.com/betaiota/grpchat/proto/chatpb"

	"google.golang.org/grpc"
)

type Connection struct {
	stream pb.Chat_CreateStreamServer
	id     string
	active bool
	err    chan error
}

type server struct {
	pb.UnimplementedChatServer
	Connection []*Connection
	storage    *storage.RedisStorage
}

func (s *server) CreateStream(pconn *pb.Connect, stream pb.Chat_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     pconn.Creds.Username,
		active: true,
		err:    make(chan error),
	}

	if s.storage != nil {
		messages, err := s.storage.GetRecentMessages(50)
		if err != nil {
			log.Printf("Warning: failed to retrieve recent messages for new user %s: %v", conn.id, err)
		} else {
			for i := len(messages) - 1; i >= 0; i-- {
				if err := stream.Send(messages[i]); err != nil {
					log.Printf("Error sending history message to %s: %v", conn.id, err)
					return err
				}
			}
			log.Printf("Sent %d recent messages to new user %s", len(messages), conn.id)
		}
	}

	s.Connection = append(s.Connection, conn)
	return <-conn.err
}

func (s *server) BroadcastMessage(ctx context.Context, message *pb.ChatMessage) (*pb.Close, error) {
	if s.storage != nil {
		if err := s.storage.StoreMessage(message); err != nil {
			log.Printf("Warning: failed to store message in Redis: %v", err)
		}
	}

	wait := sync.WaitGroup{}
	done := make(chan int)

	for _, conn := range s.Connection {
		wait.Add(1)

		go func(msg *pb.ChatMessage, conn *Connection) {
			defer wait.Done()

			if conn.active {
				if err := conn.stream.Send(msg); err != nil {
					log.Printf("Could not broadcast to %v, %v", conn.id, err)
					conn.active = false
					conn.err <- err
					return
				}
			}
			log.Printf("Sending message to: %v", conn.stream)
		}(message, conn)
	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return &pb.Close{}, nil
}

func CreateChatServer(port string, redisAddr string) {
	var connections []*Connection
	var redisStorage *storage.RedisStorage

	if redisAddr != "" {
		var err error
		redisStorage, err = storage.NewRedisStorage(redisAddr)
		if err != nil {
			log.Printf("Warning: Failed to connect to Redis at %s: %v. Continuing without Redis storage.", redisAddr, err)
			redisStorage = nil
		}
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error creating TCP listener: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterChatServer(srv, &server{
		Connection: connections,
		storage:    redisStorage,
	})

	fmt.Printf("Server listening on %v...\n", port)
	if redisStorage != nil {
		fmt.Printf("Redis storage enabled at %s\n", redisAddr)
	}

	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Error: Server could not initialize: %v", err)
	}

	if redisStorage != nil {
		redisStorage.Close()
	}
}
