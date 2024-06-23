package requests

type ConfigRequest struct {
	Endpoints []EndpointConfig
}

type EndpointConfig struct {
	Port     string
	Percents int
}
