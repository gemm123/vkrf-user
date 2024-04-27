package service

import (
	"github.com/gemm123/vkrf-user/helper"
	"github.com/gemm123/vkrf-user/internal/model"
	"github.com/gemm123/vkrf-user/internal/repository"
	"github.com/google/uuid"
	"mime/multipart"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	CreateUser(user model.UserRequest, image *multipart.FileHeader) error
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) CreateUser(user model.UserRequest, image *multipart.FileHeader) error {
	imageUrl, err := helper.UploadToS3(image)
	if err != nil {
		return err
	}

	newUser := model.User{
		Id:         uuid.New(),
		Name:       user.Name,
		Email:      user.Email,
		ProfilePic: imageUrl,
	}

	err = s.userRepository.CreateUser(newUser)
	if err != nil {
		return err
	}

	return nil
}
