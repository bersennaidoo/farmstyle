package mongo

import (
	"context"
	"time"

	"github.com/bersennaidoo/farmstyle/domain/models"
	"github.com/google/uuid"
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

func (r *ReviewsRepository) GetReviewsFiltered(filters models.ReviewFilters) *[]models.Review {
	return nil
}

func (r *ReviewsRepository) GetReviews() *[]models.Review {
	return nil
}

func (r *ReviewsRepository) AddReview(review models.Review) (*models.Review, error) {
	collection := r.mclient.Database("farmstyle").Collection("reviews")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	review.Uuid = uuid.New().String()
	_, _ = collection.InsertOne(ctx, review)

	return &review, nil
}
