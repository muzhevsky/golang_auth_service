package databases

type Database interface {
	Connect() error
	Disconnect() error
}
