package model

type TopMovies struct {
	Items []struct {
		ID              string `json:"id"`
		Rank            string `json:"rank"`
		Title           string `json:"title"`
		FullTitle       string `json:"fullTitle"`
		Year            string `json:"year"`
		Image           string `json:"image"`
		Crew            string `json:"crew"`
		ImDbRating      string `json:"imDbRating"`
		ImDbRatingCount string `json:"imDbRatingCount"`
	} `json:"items"`
	ErrorMessage string `json:"errorMessage"`
}
