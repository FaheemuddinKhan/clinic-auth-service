package port

import (
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/domain/dao"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserRepo interface {
	Store(dao.User) (*primitive.ObjectID, *errors.AppError)
	Find(primitive.ObjectID) (*dao.User, *errors.AppError)
	Modify(dao.User) *errors.AppError
	CheckExists(string, interface{}) (bool, *errors.AppError)
}
