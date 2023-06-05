package service

import (
	"ToDoLIst/repository"
	"crypto/sha1"
	"fmt"
	"github.com/zhashkevych/todo-app"
)

const salt = "dagdfsgsdfg13e42r5fsd"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}

}
func (s *AuthService) CreatUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreatUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
