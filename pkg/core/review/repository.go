package review

type Repository interface {
	AddReview(Review) error
	GetReviews() ([]Review, error)
}
