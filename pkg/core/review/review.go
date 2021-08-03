package review

type Review struct {
	BeerID       string `json:"beer_id"`
	ReviewerName string `json:"reviewer_name"`
	Score        int    `json:"score"`
	Text         string `json:"text"`
}

type Reviews struct {
	Reviews []Review `json:"reviews"`
}
