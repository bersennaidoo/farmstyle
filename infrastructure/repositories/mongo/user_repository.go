package mongo

import (
	"context"
	"time"

	"github.com/bersennaidoo/farmstyle/domain/models"
	"github.com/bersennaidoo/farmstyle/foundation/hash"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (u *UserRepository) AddUser(user models.NewUser) (*models.NewUser, error) {
	err := hash.HashPassword(&user.Password)
	if err != nil {
		return nil, err
	}
	user.Uuid = uuid.New().String()
	collection := u.client.Database("agentco").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, _ = collection.InsertOne(ctx, user)

	return &user, nil
}
