package api

	
type Fact struct {
	ID int `json:"id"`
	Category string `json:"category"`
	Title string `json:"title"`
	Fact string `json:"fact"`
	Verified bool `json:"verified"`
	Source string `json:"source"`
	YearDiscovered int `json:"year_discovered"`
	InterestingRating float64 `json:"interesting_rating"`
}