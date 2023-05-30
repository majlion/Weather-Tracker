import React, { useState } from 'react';
import axios from 'axios';

function App() {
    const [city, setCity] = useState('');
    const [weatherData, setWeatherData] = useState(null);

    const fetchWeatherData = () => {
        axios.get(`/api/weather?city=${city}`)
            .then(response => {
                setWeatherData(response.data);
            })
            .catch(error => {
                console.log('Error fetching weather data:', error);
            });
    };

    return (
        <div>
            <h1>Weather Tracking App</h1>
            <input
                type="text"
                value={city}
                onChange={event => setCity(event.target.value)}
                placeholder="Enter city name"
            />
            <button onClick={fetchWeatherData}>Get Weather Data</button>

            {weatherData && (
                <div>
                    <h2>Weather for {city}</h2>
                    <p>Temperature: {weatherData.temperature}°C</p>
                    <p>Humidity: {weatherData.humidity}%</p>
                    <p>Description: {weatherData.description}</p>
                </div>
            )}
        </div>
    );
}

export default App;
