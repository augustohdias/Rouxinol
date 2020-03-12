package repository

import (
	"context"
	"github.com/augustohdias/rouxinol/models"
	"github.com/augustohdias/rouxinol/post"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type mongoDBPostRepository struct {
	Collection *mongo.Collection
}

func (m *mongoDBPostRepository) GetByID(ctx context.Context, id string) (*models.Post, error) {
	result := m.Collection.FindOne(ctx, bson.D{{"post_id", id}})
	var p models.Post
	err := result.Decode(&p)
	return &p, err
}

func (m *mongoDBPostRepository) GetAllByUserID(ctx context.Context, id string) ([]*models.Post, error) {
	return m.find(ctx, bson.D{{"user_id", id}})
}

func (m *mongoDBPostRepository) GetAllByUserIDs(ctx context.Context, ids []string) ([]*models.Post, error) {
	return m.find(ctx, bson.D{{"user_id", bson.D{{"$in", ids}}}})
}

func (m *mongoDBPostRepository) New(ctx context.Context, post models.Post) error {
	panic("implement me")
}

func (m *mongoDBPostRepository) Update(ctx context.Context, post models.Post) error {
	panic("implement me")
}

func (m *mongoDBPostRepository) Delete(ctx context.Context, post models.Post) error {
	panic("implement me")
}

func (m *mongoDBPostRepository) find(ctx context.Context, filter bson.D) ([]*models.Post, error) {
	var posts []*models.Post
	result, err := m.Collection.Find(ctx, filter)
	if err != nil {
		log.Fatalf("error: Unable to perform user_posts.Find: %v\n", filter)
		return nil, err
	}

	err = result.All(ctx, &posts)
	return posts, err
}

func NewMongoDBPostRepository(db *mongo.Database) post.Repository {
	return &mongoDBPostRepository{db.Collection("user_posts")}
}
