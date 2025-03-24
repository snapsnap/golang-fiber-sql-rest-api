package services

import (
	"context"
	"database/sql"
	"errors"
	"rest-project/app/config"
	mdldomain "rest-project/app/models/mdl-domain"
	"rest-project/app/models/request"
	"rest-project/app/models/response"
	"rest-project/app/repositories"
	"rest-project/app/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Register(ctx context.Context, req *request.ReqRegister) error
	Login(ctx context.Context, req *request.ReqLogin) (*response.ResLogin, error)
}

type AuthServiceImpl struct {
	DB       *sql.DB
	userRepo repositories.UserRepo
}

// Login implements AuthService.
func (a *AuthServiceImpl) Login(ctx context.Context, req *request.ReqLogin) (*response.ResLogin, error) {
	result := response.ResLogin{}

	tx, err := a.DB.Begin()
	if err != nil {
		return &result, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Cek email sudah terdaftar
	foundUser, err := a.userRepo.FindByEmail(ctx, tx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if foundUser == nil {
		return &result, errors.New("not found user")
	}

	// compare password
	errVerify := utils.VerifyPassword(foundUser.Password, req.Password)
	if !errVerify {
		return &result, errors.New("invalid credentials")
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": foundUser.Id,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	cnf := config.Get()
	tokenString, err := token.SignedString([]byte(cnf.Jwt.Key))
	if err != nil {
		return &result, errors.New("failed to generate token")
	}

	result = response.ResLogin{
		Token: tokenString,
	}

	return &result, nil
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
