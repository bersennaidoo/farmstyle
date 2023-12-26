package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/bersennaidoo/farmstyle/domain/models"
	"github.com/bersennaidoo/farmstyle/foundation/emsg"
	"github.com/bersennaidoo/farmstyle/foundation/hash"
	"github.com/bersennaidoo/farmstyle/foundation/token"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client     *mongo.Client
	tokenmaker *token.PasetoMaker
}

func NewUserRepository(client *mongo.Client, tokenmaker *token.PasetoMaker) *UserRepository {
	return &UserRepository{
		client:     client,
		tokenmaker: tokenmaker,
	}
}

func (u *UserRepository) AddUser(user models.NewUser) (*models.NewUser, error) {
	collection := u.client.Database("agentco").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := models.User{}
	_ = collection.FindOne(ctx, bson.D{{"username", user.Username}}).Decode(&result)
	if result.Username == user.Username {
		return nil, emsg.CreateAlreadyExists(emsg.ProblemJson{
			Instance: "v1" + "/" + user.Username,
			Detail:   fmt.Sprintf("User with username, %s already exists.", user.Username),
		})
	}

	err := hash.HashPassword(&user.Password)
	if err != nil {
		return nil, err
	}
	user.Uuid = uuid.New().String()
	collection = u.client.Database("agentco").Collection("users")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, _ = collection.InsertOne(ctx, user)
	user.Password = ""
	user.FullName = ""

	return &user, nil
}

func (u *UserRepository) CreateToken(user models.UserLogin) (string, error) {
	// create token string and return to handler
	collection := u.client.Database("agentco").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := models.UserLogin{}
	_ = collection.FindOne(ctx, bson.D{{"username", user.Username}}).Decode(&result)
	if ok := hash.CheckPassword(result.Password, user.Password); !ok {
		return "", emsg.InvalidCreds(emsg.ProblemJson{
			Detail: "Username or password is invalid",
		})
	}
	token, _ := u.tokenmaker.CreateToken(result.Username, "123")
	return token, nil
}
