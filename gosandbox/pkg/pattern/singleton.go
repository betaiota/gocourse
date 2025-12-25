package pattern

var instance *dbConnection

type dbConnection struct {
	host    string
	port    int
	dbName  string
	dbLogin string
	dbPass  string
}

func init() {
	GetInstance()
}

func GetInstance() *dbConnection {
	if instance == nil {
		instance = &dbConnection{"localhost", 5436, "testDb", "testLogin", "1234"}
		return instance
	}
	return instance
}
