// added by xiaonan

package lmfService

import (
	"fmt"
	liblog "github.com/wangxn2015/onos-lib-go/pkg/logging"
	"google.golang.org/grpc"
	"net"
)

var log = liblog.GetLogger()

type ServerConfig struct {
	Port int
}

type Service interface {
	Register(s *grpc.Server)
}

type Server struct {
	cfg      *ServerConfig
	services []Service
	server   *grpc.Server
	//Uestore  stores.UeStore
}

func NewServer(configPort int) *Server {
	return &Server{
		cfg: &ServerConfig{
			Port: configPort,
		},
		services: []Service{},
		//Uestore:  stores.NewInMemoryUeStore(),
	}

}

func (s *Server) AddService(r Service) {
	s.services = append(s.services, r)
}

func (s *Server) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	log.Infof("lmf listens on %v", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		return err
	}
	s.server = grpc.NewServer()
	for i := range s.services {
		s.services[i].Register(s.server)
	}
	log.Info("lmf begin to server")
	return s.server.Serve(lis)
}
