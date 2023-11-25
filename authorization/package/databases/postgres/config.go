package postgres

type ConnectionString struct {
	user         string
	password     string
	hostname     string
	port         string
	databaseName string
}

func (str *ConnectionString) SetUser(user string) *ConnectionString {
	str.user = user
	return str
}

func (str *ConnectionString) SetPassword(password string) *ConnectionString {
	str.password = password
	return str
}

func (str *ConnectionString) SetHostname(hostname string) *ConnectionString {
	str.hostname = hostname
	return str
}

func (str *ConnectionString) SetPort(port string) *ConnectionString {
	str.port = port
	return str
}

func (str *ConnectionString) SetDatabaseName(databaseName string) *ConnectionString {
	str.databaseName = databaseName
	return str
}

func (str *ConnectionString) Build() string {
	return "postgresClient://" + str.user + ":" + str.password + "@" + str.hostname + ":" + str.port + "/" + str.databaseName + "?sslmode=disable"
}
