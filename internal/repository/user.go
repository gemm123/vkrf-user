package repository

import (
	"context"
	"github.com/gemm123/vkrf-user/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type userRepository struct {
	db *pgxpool.Pool
}

type UserRepository interface {
	CreateUser(user model.User) error
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user model.User) error {
	query := "INSERT INTO users (id, name, email, profile_pic) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(context.Background(), query, user.Id, user.Name, user.Email, user.ProfilePic)
	if err != nil {
		log.Printf("REPO: Error creating user: %v\n", err)
		return err
	}

	return nil
}
