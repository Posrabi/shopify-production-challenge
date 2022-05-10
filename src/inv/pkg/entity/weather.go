package entity

// Simple weather report, I am ignoring lots of temperatures/ids fields.
type WeatherReport struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Name string `json:"name"`
}
