package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"vlad_xp/internal/models"
	"vlad_xp/internal/repository"
)

type userService struct {
	userRepo repository.UserRepo
}

func InitUserService(userRepo repository.UserRepo) User {
	return userService{
		userRepo: userRepo,
	}
}

func (u userService) GetMe(ctx context.Context, id int) (*models.GetUser, error) {
	user, err := u.userRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userService) DeleteMe(ctx context.Context, id int) error {
	err := u.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u userService) Create(ctx context.Context, user models.CreateUser) (int, error) {
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), 11)
	user.HashedPassword = string(hashedPwd)

	id, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u userService) Login(ctx context.Context, user models.CreateUser) error {
	hashedPwd, err := u.userRepo.GetPwdByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(user.HashedPassword))
	if err != nil {
		return err
	}

	return nil
}
