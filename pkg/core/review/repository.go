package review

type Repository interface {
	AddReview(Review) error
}
