package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aurennunes/iptv/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	baseURL := os.Getenv("XTREAM_URL")
	username := os.Getenv("XTREAM_USER")
	password := os.Getenv("XTREAM_PASS")

	xtream := services.NewXtream(baseURL, username, password)
	download := services.NewDownload(xtream, "/home/auren/Downloads")

	var seriesID int
	fmt.Print("Digite o ID da série: ")
	fmt.Scan(&seriesID)

	download.Serie(strconv.Itoa(seriesID))
}
