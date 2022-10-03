package uesposition

import (
	"fmt"
	"github.com/wangxn2015/ran-simulator/pkg/lmfService"
	"github.com/wangxn2015/ran-simulator/pkg/proto/track_msg"
	"github.com/wangxn2015/ran-simulator/pkg/store/ues"
	//"github.com/wangxn2015/ran-simulator/pkg/myfmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewService() 返回一个接口（支持Register）， 这个接口实际是一个Service struct
func NewService(store ues.Store) lmfService.Service {
	return &Service{
		uestore: store,
	}
}

type Service struct {
	//lmfService.Service //explain here
	uestore ues.Store
}

// Register registers the TrafficSim Service with the gRPC server.
//
func (s *Service) Register(r *grpc.Server) {
	server := &Server{
		ueStore: s.uestore,
	}
	track_msg.RegisterUEsServer(r, server)
}

//---------------------------------------------------------------

// Server implements the Routes gRPC service for administrative facilities.
type Server struct {
	track_msg.UnimplementedUEsServer
	ueStore ues.Store
}

func (server *Server) GetUEs(request *track_msg.UERequest, stream track_msg.UEs_GetUEsServer) error {
	//myfmt.Logger.Println("Get UEs request: %+v", request)
	fmt.Println("Get UEs request: %+v", request)

	err := server.ueStore.Search(
		stream.Context(),
		func(ue *track_msg.UEInfo) error {
			//res := &track_msg.UEInfo{Laptop: ue}
			res := ue
			err := stream.Send(res)
			if err != nil {
				return err
			}

			fmt.Printf("sent ue with id: %d\n", ue.GetImsi())
			return nil
		},
	)

	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}

	return nil
	//return status.Errorf(codes.Unimplemented, "method GetUEs not implemented")
}
