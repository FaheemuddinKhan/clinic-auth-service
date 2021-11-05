package repo

import (
	"context"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/domain/dao"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/errors"
	"github.com/faheemuddinkhan/clinic-auth-service/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepo struct {
	client  *mongo.Client
	dbName  string
	timeout time.Duration
}

func (r UserRepo) Store(user dao.User) (*primitive.ObjectID, *errors.AppError) {
	collection := r.client.Database(r.dbName).Collection("User")
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)

	defer cancel()
	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		logger.Error("Error while creating new user: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected Database error")
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return &user.ID, nil
}

func (r UserRepo) Find(ID primitive.ObjectID) (*dao.User, *errors.AppError) {
	collection := r.client.Database(r.dbName).Collection("User")
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)

	defer cancel()
	filter := bson.M{"_id": ID}
	var user dao.User

	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		logger.Error("Error while retrieving user: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected Database error")
	}
	return &user, nil
}

func (r UserRepo) Modify(user dao.User) *errors.AppError {
	collection := r.client.Database(r.dbName).Collection("User")
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)

	defer cancel()
	filter := bson.M{"_id": user.ID}
	result, err := collection.UpdateOne(
		ctx,
		filter,
		user,
	)

	if err != nil {
		logger.Error("Error while modifying user: " + err.Error())
		return errors.NewUnexpectedError("Unexpected Database error")
	}
	if result.ModifiedCount == 0 {
		logger.Info("Record not found in MomgoDB while modifying User Collection")
		return errors.NewUnexpectedError("Unexpected Database error")
	}
	return nil
}

func (r UserRepo) CheckExists(fieldName string, value interface{}) (bool, *errors.AppError) {
	collection := r.client.Database(r.dbName).Collection("User")
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)

	defer cancel()
	filter := bson.M{fieldName: value}
	itemCount, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		logger.Error("Error while modifying user: " + err.Error())
		return false, errors.NewUnexpectedError("Unexpected Database error")
	}

	if itemCount > 0 {
		return true, nil
	}
	return false, nil

}

func NewUserRepo(client *mongo.Client, dbName string, timeout time.Duration) UserRepo {
	return UserRepo{
		client:  client,
		dbName:  dbName,
		timeout: timeout,
	}
}
