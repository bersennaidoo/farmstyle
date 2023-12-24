package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/bersennaidoo/farmstyle/application/rest/problems"
	"github.com/bersennaidoo/farmstyle/domain/models"
	"github.com/bersennaidoo/farmstyle/infrastructure/repositories/mongo"
	"github.com/kataras/golog"
)

type ReviewsHandler struct {
	reviewsRepository *mongo.ReviewsRepository
	log               *golog.Logger
}

func NewReviewsHandler(reviewsRepository *mongo.ReviewsRepository, log *golog.Logger) *ReviewsHandler {
	return &ReviewsHandler{
		reviewsRepository: reviewsRepository,
		log:               log,
	}
}

func (rh *ReviewsHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	maxRating := query.Get("maxRating")
	log.Printf("\nmaxRating: %s\n", maxRating)
	var reviewList *[]models.Review
	if maxRating != "" {
		i, _ := strconv.Atoi(maxRating)
		filters := models.ReviewFilters{
			MaxRating: i,
		}
		reviewList = rh.reviewsRepository.GetReviewsFiltered(filters)
	} else {
		reviewList = rh.reviewsRepository.GetReviews()
	}
	writeJson(200, reviewList)(w, r)
}

func (rh *ReviewsHandler) AddReview(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var review models.Review
	err := decoder.Decode(&review)
	if err != nil {
		ErrorResponse(problems.FailedToParseJson(problems.ProblemJson{
			Detail: err.Error(),
		}))(w, r)
		return
	}

	res, _ := rh.reviewsRepository.AddReview(review)
	writeJson(201, res)(w, r)
}
