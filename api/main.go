package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// WeatherData represents the structure of weather data
type WeatherData struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Description string  `json:"description"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")

	// Initialize router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/weather", getWeatherData).Methods("GET")

	// Start server
	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Function to fetch weather data from the API
func fetchWeatherData(city string, apiKey string) (*WeatherData, error) {
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey + "&units=metric"

	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}

	weatherData := WeatherData{}
	if err := json.Unmarshal(resp.Body(), &weatherData); err != nil {
		return nil, err
	}

	return &weatherData, nil
}

// Handler function to get weather data
func getWeatherData(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	weatherData, err := fetchWeatherData(city, apiKey)
	if err != nil {
		log.Println("Error fetching weather data:", err)
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}
