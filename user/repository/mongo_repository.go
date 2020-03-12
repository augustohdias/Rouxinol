package repository

import (
	"context"
	"github.com/augustohdias/rouxinol/models"
	"github.com/augustohdias/rouxinol/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDBUserRepo struct {
	Collection *mongo.Collection
}

func (m *mongoDBUserRepo) GetById(ctx context.Context, id string) (*models.User, error) {
	filter := bson.D{{"user_id", id}}
	return m.findOne(ctx, filter)
}

func NewMongoDBUserRepository(db *mongo.Database) user.Repository {
	return &mongoDBUserRepo{db.Collection("users")}
}

func (m *mongoDBUserRepo) findOne(ctx context.Context, filter bson.D) (*models.User, error) {
	result := m.Collection.FindOne(ctx, filter)
	var usr models.User
	err := result.Decode(&usr)
	return &usr, err
}
