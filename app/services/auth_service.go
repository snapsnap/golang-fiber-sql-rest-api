package services

import (
	"context"
	"database/sql"
	"errors"
	mdldomain "rest-project/app/models/mdl-domain"
	"rest-project/app/models/request"
	"rest-project/app/repositories"
	"time"
)

type AuthService interface {
	Register(ctx context.Context, req *request.ReqRegister) error
}

type AuthServiceImpl struct {
	DB       *sql.DB
	userRepo repositories.UserRepo
}

// Register implements AuthService.
func (a *AuthServiceImpl) Register(ctx context.Context, req *request.ReqRegister) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Cek email sudah terdaftar
	foundUser, err := a.userRepo.FindByEmail(ctx, tx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if foundUser != nil {
		return errors.New("registered email")
	}

	user := mdldomain.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	err = a.userRepo.Save(tx, &user)
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err
}

func NewAuthServiceImpl(db *sql.DB, userRepository repositories.UserRepo) AuthService {
	return &AuthServiceImpl{
		DB:       db,
		userRepo: userRepository,
	}
}
