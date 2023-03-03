package usecase

import (
	"errors"
	jwt2 "github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"go-gorm/entity"
	"go-gorm/internal/repository/model"
	"go-gorm/pkg/jwt"
	"go-gorm/pkg/util"
	"time"
)

func (u *Usecase) CreateUser(user *entity.UserDTO) error {
	if user == nil {
		return errors.New("nil user")
	}

	if err := user.Validate(entity.ValidateCreate); err != nil {
		return err
	}

	return u.User.Create(model.UserToModel(u.uuidfn(), user))
}

func (u *Usecase) Login(username, password string) (*entity.UserDTO, map[string]interface{}, *util.UsecaseError) {
	var (
		user         *entity.UserDTO
		usecaseError = util.NewUsecaseError()
		err          error
		TokenMap     map[string]interface{}
	)

	if password == "" {
		usecaseError.Add(errors.New("empty password"), 400)
	}
	if username == "" {
		usecaseError.Add(errors.New("empty username"), 400)
	}
	if usecaseError.HasError() {
		return nil, nil, usecaseError
	}

	if username != "" {
		user, err = u.User.GetUser(username, "username")
		if err != nil {
			usecaseError.Add(err, 400)
			return nil, nil, usecaseError
		}
	}

	err = util.ComparePassword(user.Password, password)
	if err != nil {
		err = errors.New("wrong password")
		return nil, nil, usecaseError.Add(err, 400)
	}

	accessToken, errJwt := u.Jwt.GenerateToken(&jwt.Payload{
		jwt2.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        user.ID.String(),
		},
	})
	if errJwt != nil {
		return nil, nil, usecaseError.Add(errJwt, 400)
	}

	refreshToken, errJwt := u.Jwt.GenerateToken(&jwt.Payload{
		jwt2.StandardClaims{
			ExpiresAt: time.Now().Add(20000 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        user.ID.String(),
		},
	})
	if errJwt != nil {
		return nil, nil, usecaseError.Add(errJwt, 400)
	}

	TokenMap = map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return user, TokenMap, nil
}

func (u *Usecase) GetUserByID(id uuid.UUID) (*entity.UserDTO, error) {
	if id == uuid.Nil {
		return nil, errors.New("empty id")
	}

	user, err := u.User.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Usecase) UpdateUser(newData *entity.UserDTO) error {
	if newData == nil {
		return errors.New("nil user")
	}

	oldData, err := u.User.GetByID(newData.ID)
	if err != nil {
		return err
	}

	oldData.Update(newData)

	if err := oldData.Validate(entity.ValidateUpdate); err != nil {
		return err
	}

	return u.User.Update(model.UserToModel(u.uuidfn(), oldData))
}

func (u *Usecase) DeleteUser(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("empty id")
	}

	return u.User.Delete(id)
}
