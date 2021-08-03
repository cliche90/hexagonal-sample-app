package beer

type Beer struct {
	Name    string  `json:"name"`
	Brewery string  `json:"brewery"`
	Abv     float32 `json:"abv"`
	Desc    string  `json:"desc"`
}
