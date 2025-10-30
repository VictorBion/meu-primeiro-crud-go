package model

import "golang.org/x/crypto/bcrypt"

func (ud *userDomain) EncryptPassWord() error {
	hashedPassWord, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	ud.Password = string(hashedPassWord)
	return nil
}