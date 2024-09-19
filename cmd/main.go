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

	baseDIR := os.Getenv("BASE_DIR")
	baseURL := os.Getenv("XTREAM_URL")
	username := os.Getenv("XTREAM_USER")
	password := os.Getenv("XTREAM_PASS")

	xtream := services.NewXtream(baseURL, username, password)
	download := services.NewDownload(xtream, baseDIR)

	var seriesID int
	fmt.Print("Digite o ID da série: ")
	fmt.Scan(&seriesID)

	download.Serie(strconv.Itoa(seriesID))
}
