package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	SetID(string)
	GetID() string

	GetJsonValue() (string, error)

	EncryptPassWord() error
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}