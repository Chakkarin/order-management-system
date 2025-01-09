package controllers

import (
	"fmt"
	"services/auth-service/shared/utils"
)

func authValidate(email, pass *string) error {
	if *email == "" || *pass == "" {
		return fmt.Errorf("username or password not empty")
	}

	if !utils.IsEmail(email) {
		return fmt.Errorf("invalid email format")
	}

	return nil
}
