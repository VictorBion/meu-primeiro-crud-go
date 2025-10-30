package converter

import (
	"github.com/VictorBion/meu-primeiro-crud-go/src/model"
	"github.com/VictorBion/meu-primeiro-crud-go/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity.UserEntity,
) model.UserDomainInterface {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}