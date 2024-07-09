package server

import (
	"context"

	"UserService/proto" // Update this import path to match your project
)

type Server struct {
	proto.UnimplementedUserServiceServer
}

func (s *Server) GetUserByID(ctx context.Context, req *proto.GetUserByIDRequest) (*proto.GetUserByIDResponse, error) {
	// Implement your logic here
	return &proto.GetUserByIDResponse{
		User: &proto.User{
			Id:      1,
			Fname:   "Steve",
			City:    "LA",
			Phone:   1234567890,
			Height:  5.8,
			Married: true,
		},
	}, nil
}

func (s *Server) GetUsersByIDs(ctx context.Context, req *proto.GetUsersByIDsRequest) (*proto.GetUsersByIDsResponse, error) {
	// Implement your logic here
	return &proto.GetUsersByIDsResponse{
		Users: []*proto.User{
			{
				Id:      1,
				Fname:   "Steve",
				City:    "LA",
				Phone:   1234567890,
				Height:  5.8,
				Married: true,
			},
		},
	}, nil
}

func (s *Server) SearchUsers(ctx context.Context, req *proto.SearchUsersRequest) (*proto.SearchUsersResponse, error) {
	// Implement your logic here
	return &proto.SearchUsersResponse{
		Users: []*proto.User{
			{
				Id:      1,
				Fname:   "Steve",
				City:    "LA",
				Phone:   1234567890,
				Height:  5.8,
				Married: true,
			},
		},
	}, nil
}
