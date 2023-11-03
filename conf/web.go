package conf

import (
	"fmt"
	"os"
	"strconv"
)

type WebConf struct {
	Port int
}

func EnvWebConf() WebConf {
	port := os.Getenv("APPLICATION_PORT")
	if port == "" {
		port = "8000"
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	return WebConf{
		Port: portInt,
	}
}

func (w WebConf) ListenerAddr() string {
	return fmt.Sprintf(":%d", w.Port)
}
