package recommender

type RecommendProvider interface {
}
type recommendProvider struct {
}

func NewRecommendProvider() RecommendProvider {
	return &recommendProvider{}
}
