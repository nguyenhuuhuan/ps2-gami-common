package configs

import (
	"fmt"
)

type GRPC struct {
	GamiServiceHost string `default:"127.0.0.1" envconfig:"GAMI_SERVICE_GRPC_HOST"`
	GamiServicePort int    `default:"8888" envconfig:"GAMI_SERVICE_GRPC_PORT"`
}

func (g *GRPC) GamiServiceAddress() string {
	return fmt.Sprintf("%v:%v", g.GamiServiceHost, g.GamiServicePort)
}
