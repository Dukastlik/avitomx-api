package apiserver


type APIServer struct {
	config *Config
}

func New(config *Config) *APIServer{
	return &APIServer{
		config: config,
	}
}

func Start(s *APIServer)  error{
	return nil
}