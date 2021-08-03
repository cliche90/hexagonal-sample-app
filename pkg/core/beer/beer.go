package beer

type Beer struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Brewery string  `json:"brewery"`
	Abv     float32 `json:"abv"`
	Desc    string  `json:"desc"`
}

type Beers struct {
	Beers []Beer `json:"beers"`
}
