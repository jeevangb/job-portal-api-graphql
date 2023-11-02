package service

import (
	"fmt"
	"graphql/graph/model"
	"graphql/models"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SignUp(userData model.NewUser) (*model.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return &model.User{}, fmt.Errorf("generating password hash: %w", err)
	}

	userDetails := models.User{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: string(hashPass),
	}
	userDetails, err = s.R.CreateUser(userDetails)
	if err != nil {
		return nil, err
	}
	uid := strconv.FormatUint(uint64(userDetails.ID), 10)

	return &model.User{
		ID:    uid,
		Name:  userDetails.Name,
		Email: userDetails.Email,
	}, nil

}
