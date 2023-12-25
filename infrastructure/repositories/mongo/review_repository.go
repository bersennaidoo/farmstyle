package mongo

import (
	"context"
	"time"

	"github.com/bersennaidoo/farmstyle/domain/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReviewsRepository struct {
	mclient *mongo.Client
}

func NewReviewsRepository(mclient *mongo.Client) *ReviewsRepository {
	return &ReviewsRepository{
		mclient: mclient,
	}
}

func (r *ReviewsRepository) GetReviewsFiltered(filters models.ReviewFilters) []models.Review {
	collection := r.mclient.Database("farmstyle").Collection("reviews")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{"rating", filters.MaxRating}}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil
	}
	var reviews []models.Review
	if err = cur.All(context.Background(), &reviews); err != nil {
		return nil
	}

	return reviews
}

func (r *ReviewsRepository) GetReviews() []models.Review {
	collection := r.mclient.Database("farmstyle").Collection("reviews")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil
	}
	defer cur.Close(context.Background())

	var reviews []models.Review
	for cur.Next(context.Background()) {
		var review models.Review
		err = cur.Decode(&review)
		if err != nil {
			return nil
		}
		reviews = append(reviews, review)
	}

	if err := cur.Err(); err != nil {
		return nil
	}

	return reviews
}

func (r *ReviewsRepository) AddReview(review models.Review) (*models.Review, error) {
	collection := r.mclient.Database("farmstyle").Collection("reviews")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	review.Uuid = uuid.New().String()
	_, _ = collection.InsertOne(ctx, review)

	return &review, nil
}
