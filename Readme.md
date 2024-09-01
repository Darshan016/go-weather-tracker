# Weather App

This is a simple weather application built using Go. It fetches weather data from the weather API and displays it to the user. The application is also dockerized for easy deployment.

## Features

- Fetches current weather data from a weather API
- Displays weather information in a user-friendly format
- Dockerized for easy deployment

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have [Go](https://golang.org/doc/install) installed on your machine.
- You have [Docker](https://docs.docker.com/get-docker/) installed on your machine.
- You have an API key from a weather API provider (e.g., weatherapi).

## Installation

### Clone the Repository

```bash
git clone https://github.com/Darshan016/go-weather-tracker.git
cd go-weather-tracker
```

## Build the image

```bash
docker build -t weather-tracker .
```

## Run the image

```bash
docker run -p 8000:8000 -it weather-tracker
```

## Testing

Open your browser and hit this url:
```
localhost:8000/weather/city_name
```
Solution for: https://roadmap.sh/projects/weather-api-wrapper-service
Feel free to contribute.
