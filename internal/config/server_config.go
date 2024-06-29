package config

type configServer struct {
	ADDR  string
	PORT  string
	Proto string
}

var ServerConfig *configServer

func init() {
	ServerConfig = &configServer{
		ADDR:  getEnv("ADDR"),
		PORT:  getEnv("PORT"),
		Proto: getEnv("Proto"),
	}
}
