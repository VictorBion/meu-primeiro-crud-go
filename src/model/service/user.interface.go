package service

import (
	"github.com/VictorBion/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/VictorBion/meu-primeiro-crud-go/src/model"
	"github.com/VictorBion/meu-primeiro-crud-go/src/model/repository"
)

func NewUserDomainService(repo repository.UserRepository) UserDomainService {
	return &userDomainService{
		userRepository: repo,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository	
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}