package usecase

import (
	"go-gorm/domain/repository"
	"go-gorm/pkg/jwt"
	"go-gorm/pkg/util"
)

type Usecase struct {
	User     repository.User
	Room     repository.Room
	Ticket   repository.Ticket
	UserRoom repository.UserRoom
	Jwt      jwt.Jwt
	uuidfn   util.UUIDfn
}

func NewUsecase(Jwt jwt.Jwt, user repository.User, room repository.Room, ticket repository.Ticket, userroom repository.UserRoom, uuidfn util.UUIDfn) *Usecase {
	return &Usecase{
		User:     user,
		Room:     room,
		Ticket:   ticket,
		UserRoom: userroom,
		Jwt:      Jwt,
		uuidfn:   uuidfn,
	}
}
