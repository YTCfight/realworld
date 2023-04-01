package biz

import (
	"context"
	"fmt"
	"kratos/internal/conf"
	"kratos/internal/pkg/middleware/auth"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email        string
	UserName     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	UserName string
	Token    string
	Bio      string
	Image    string
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", b)
	return string(b)

}

func verifyPassword(hashed, input string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input)); err != nil {
		return false
	}
	return true
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type ProfileRepo interface {
}

type UserUseCase struct {
	ur   UserRepo
	pr   ProfileRepo
	jwtc *conf.JWT

	log *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, jwtc *conf.JWT, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		ur:   ur,
		pr:   pr,
		jwtc: jwtc,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUseCase) generateToken(username string) string {
	return auth.GenerateToken(uc.jwtc.Token, username)
}

func (uc *UserUseCase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Email:        email,
		UserName:     username,
		PasswordHash: hashPassword(password),
	}
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}

	return &UserLogin{
		Email:    email,
		UserName: username,
		Token:    uc.generateToken(username),
	}, nil
}

func (uc *UserUseCase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	if len(email) == 0 {
		return nil, errors.New(422, "email", "cannot be empty")
	}
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPassword(u.PasswordHash, password) {
		return nil, errors.Unauthorized("user", "login failed")
	}
	return &UserLogin{
		Email:    u.Email,
		UserName: u.UserName,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    uc.generateToken(u.UserName),
	}, nil
}
