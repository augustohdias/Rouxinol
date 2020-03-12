package repository

import (
	"context"
	"github.com/augustohdias/rouxinol/models"
	"github.com/augustohdias/rouxinol/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDBUserRepo struct {
	DB *mongo.Database
}

func (m *mongoDBUserRepo) GetById(ctx context.Context, id string) (*models.User, error) {
	filter := bson.D{{"user_id", id}}
	return m.findOne(ctx, filter)
}

func NewMongoDBRepository(db *mongo.Database) user.Repository {
	return &mongoDBUserRepo{
		DB: db,
	}
}

func (m *mongoDBUserRepo) findOne(ctx context.Context, filter bson.D) (*models.User, error) {
	collection := m.DB.Collection("user")
	result := collection.FindOne(ctx, filter)
	var usr models.User
	err := result.Decode(&usr)
	return &usr, err
}
