package repository

import (
	"echo-server/internal/model"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (r *userRepo) Create(user *model.User) error {
	user.HashPassword()
	err := mgm.Coll(user).Create(user)
	return err
}

func (r *userRepo) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepo) FindById(id string) (*model.User, error) {
	user := &model.User{}
	err := mgm.Coll(user).First(bson.M{"_id": id}, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
