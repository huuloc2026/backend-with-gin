package service

import "github.com/huuloc2026/go-backend/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// user service: us
func (us *UserService) GetInforUserService() string {
	return us.userRepo.GetInforUser()
}
