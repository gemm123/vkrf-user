package grpc

import (
	"context"
	"github.com/gemm123/vkrf-user/internal/repository"
	"log"
)

type Server struct {
	UserRepository repository.UserRepository
	UnimplementedUserServiceServer
}

func (g *Server) GetUserByEmail(ctx context.Context, req *GetUserByEmailRequest) (*GetUserByEmailResponse, error) {
	log.Println("received GetUserByEmail request")

	user, err := g.UserRepository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	return &GetUserByEmailResponse{
		User: &UserProto{
			Id:         user.Id.String(),
			Name:       user.Name,
			Email:      user.Email,
			ProfilePic: user.ProfilePic,
		},
	}, nil
}
