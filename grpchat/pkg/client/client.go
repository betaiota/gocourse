package client

import (
	"bufio"
	"context"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	pb "github.com/betaiota/grpchat/proto/chatpb"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
	colorGray   = "\033[90m"
)

var usernameColors = []string{
	colorRed,
	colorGreen,
	colorYellow,
	colorBlue,
	colorPurple,
	colorCyan,
	colorWhite,
	colorGray,
}

func getColorForUsername(username string) string {
	h := fnv.New32a()
	h.Write([]byte(username))
	hash := h.Sum32()
	return usernameColors[hash%uint32(len(usernameColors))]
}

func centerText(text string, width int) string {
	textLen := len([]rune(text))
	if textLen >= width {
		return text
	}
	padding := (width - textLen) / 2
	return strings.Repeat(" ", padding) + text
}

type OfflineMessage struct {
	Body      string
	Sender    string
	Timestamp int64
}

var user *pb.UserCredientials
var client pb.ChatClient
var wait *sync.WaitGroup

func connect(creds *pb.UserCredientials) error {
	var streamerror error
	stream, err := client.CreateStream(context.Background(), &pb.Connect{
		Creds:  creds,
		Active: true,
	})

	if err != nil {
		return fmt.Errorf("Connection failed: %v", err)
	}

	wait.Add(1)
	go func(str pb.Chat_CreateStreamClient) {
		defer wait.Done()
		for {
			response, err := str.Recv()
			if err != nil {
				streamerror = fmt.Errorf("Error reading message: %v", err)
				break
			}
			timestampStr := time.Unix(response.Timestamp, 0).Format("15:04:05")
			usernameColor := getColorForUsername(response.Username)
			fmt.Printf("\r[%s] %s%s%s: %s\n> ",
				timestampStr,
				usernameColor,
				response.Username,
				colorReset,
				response.Content)
		}
	}(stream)

	return streamerror
}

type ClientConfig struct {
	ServerURL  string
	ServerPort string
	Username   string
	Password   string
}

func ShowMenu() (*ClientConfig, error) {
	viper.SetConfigName("client_config")
	viper.SetConfigType("yaml")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}
	configDir := filepath.Join(homeDir, ".gprchat")
	viper.AddConfigPath(configDir)
	viper.AddConfigPath(".")

	viper.SetDefault("server.url", "localhost")
	viper.SetDefault("server.port", "5001")
	viper.SetDefault("user.username", "")
	viper.SetDefault("user.password", "")

	viper.ReadInConfig()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[2J\033[H")
	fmt.Println(strings.Repeat("=", 90))
	fmt.Printf("%s%s%s\n", colorCyan, centerText("Добро пожаловать в gRPC-чат (клиент)!", 90), colorReset)
	fmt.Printf("%s%s%s\n", colorCyan, centerText("Разработчик: Илья Бабак (github.com/betaiota)", 90), colorReset)
	fmt.Printf("%s%s%s\n", colorCyan, centerText("Дипломный проект по курсу 'Разработчик Go', ЦДПО ИТМО", 90), colorReset)
	fmt.Printf("%s%s%s\n", colorPurple, centerText("2025 год", 90), colorReset)
	fmt.Println(strings.Repeat("=", 90))
	fmt.Println()

	config := &ClientConfig{}

	defaultURL := viper.GetString("server.url")
	fmt.Printf("Укажите адрес сервера [%s%s%s]: ", colorYellow, defaultURL, colorReset)
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			config.ServerURL = defaultURL
		} else {
			config.ServerURL = input
		}
	}

	defaultPort := viper.GetString("server.port")
	fmt.Printf("Укажите порт [%s%s%s]: ", colorYellow, defaultPort, colorReset)
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			config.ServerPort = defaultPort
		} else {
			config.ServerPort = input
		}
	}

	fmt.Println()

	defaultUsername := viper.GetString("user.username")
	fmt.Printf("Укажите свой никнейм [%s%s%s]: ", colorYellow, defaultUsername, colorReset)
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			if defaultUsername == "" {
				fmt.Printf("%sUsername cannot be empty!%s\n", colorRed, colorReset)
				return nil, fmt.Errorf("username is required")
			}
			config.Username = defaultUsername
		} else {
			config.Username = input
		}
	}

	defaultPassword := viper.GetString("user.password")
	if defaultPassword != "" {
		fmt.Printf("Введите пароль [%s***%s (либо используйте сохраненный)]: ", colorYellow, colorReset)
	} else {
		fmt.Printf("Введите пароль: ")
	}
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			config.Password = defaultPassword
		} else {
			config.Password = input
		}
	}

	viper.Set("server.url", config.ServerURL)
	viper.Set("server.port", config.ServerPort)
	viper.Set("user.username", config.Username)
	viper.Set("user.password", config.Password)

	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Printf("Warning: Could not create config directory: %v", err)
	} else {
		viper.SetConfigFile(filepath.Join(configDir, "client_config.yaml"))
		if err := viper.WriteConfig(); err != nil {
			if err := viper.SafeWriteConfig(); err != nil {
				log.Printf("Warning: Could not save configuration: %v", err)
			}
		}
	}

	if config.ServerURL == "" {
		return nil, fmt.Errorf("server URL cannot be empty")
	}
	if config.ServerPort == "" {
		return nil, fmt.Errorf("server port cannot be empty")
	}
	if config.Username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}

	fmt.Println()
	fmt.Printf("%sConnecting to %s:%s as %s...%s\n", colorGreen, config.ServerURL, config.ServerPort, config.Username, colorReset)
	fmt.Println()

	return config, nil
}

func CreateChatClient(username, password, url, port string) {
	if url == "" {
		log.Fatalf("%sError: Server URL cannot be empty%s\n", colorRed, colorReset)
	}
	if port == "" {
		log.Fatalf("%sError: Server port cannot be empty%s\n", colorRed, colorReset)
	}
	if username == "" {
		log.Fatalf("%sError: Username cannot be empty%s\n", colorRed, colorReset)
	}

	wait = &sync.WaitGroup{}
	done := make(chan int)

	address := url + ":" + port
	fmt.Printf("%sConnecting to %s...%s\n", colorCyan, address, colorReset)

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("%sError: Could not create a client connection to %s: %v%s\n", colorRed, address, err, colorReset)
	}
	defer conn.Close()

	fmt.Printf("%sConnected successfully!%s\n", colorGreen, colorReset)

	client = pb.NewChatClient(conn)
	userCreds := &pb.UserCredientials{
		Username: username,
		Password: password,
	}

	if err := connect(userCreds); err != nil {
		log.Fatalf("%sError: Failed to establish stream: %v%s\n", colorRed, err, colorReset)
	}

	wait.Add(1)
	go func() {
		defer wait.Done()

		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("> ")
			os.Stdout.Sync() // Ensure prompt is displayed
			if !scanner.Scan() {
				break
			}

			messageText := scanner.Text()
			if len(messageText) == 0 {
				continue
			}

			fmt.Print("\033[1A\r\033[2K")
			os.Stdout.Sync() // Ensure clearing is applied immediately

			msg := &pb.ChatMessage{
				Username:  userCreds.Username,
				Content:   messageText,
				Timestamp: time.Now().Unix(),
			}
			_, err := client.BroadcastMessage(context.Background(), msg)
			if err != nil {
				fmt.Printf("Error sending message: %v\n> ", err)
				os.Stdout.Sync()
				continue
			}
		}
	}()

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
}
