package usecases

import (
	"context"
	"errors"
	"fmt"
	"services/auth-service/shared/constants"
	"services/auth-service/shared/models"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserUsecase) Register(ctx context.Context, user *models.User) error {

	// check email exists
	isDupEmail, err := u.UserRepo.HasEmail(ctx, &user.Email)
	if err != nil {
		return err
	}

	saveUser := func() error {

		// Hash Password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password : %v", err.Error())
		}
		user.Password = string(hashedPassword)

		// save user
		if err := u.UserRepo.SaveUser(ctx, user); err != nil {
			return err
		}

		// ส่ง email ใหม่
		if err = u.sendEmailUsecase(&user.Email, &constants.NAME_VERIFIER_TYPE); err != nil {
			return err
		}

		return nil
	}

	if *isDupEmail {
		// check email verified
		isEmailVerified, err := u.UserRepo.HasEmailVerified(ctx, &user.Email)
		if err != nil {
			return err
		}

		if !*isEmailVerified {

			if err := saveUser(); err != nil {
				return err
			}

			return nil
		}

		return errors.New("email already exists")
	}

	if err := saveUser(); err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) Verify(ctx context.Context, text *string) error {

	/*
		รับ text base64 มา decode email|<string ที่ถูก hash ไว้ใน redis>
		<string ที่ถูก hash ไว้ใน redis> เมื่อถอดออกมาจะได้ email ของ user ที่ต้องการ verify
	*/

	email := text

	// update verified email = true
	return u.UserRepo.EmailVerified(ctx, email)
}
