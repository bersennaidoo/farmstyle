package mongo

import (
	"github.com/bersennaidoo/farmstyle/domain/models"
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

func (u *UserRepository) AddUser(user models.NewUser) (models.NewUser, error) {
	/*collection := u.client.Database("agentco").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, _ = collection.InsertOne(ctx, user)*/

	return user, nil
}
