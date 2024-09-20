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

	var item int
	fmt.Println("Menu:")
	fmt.Println("")
	fmt.Println("1 - Biaxar serie completa")
	fmt.Println("2 - Uma temporada de uma serie")
	fmt.Println("3 - Baixar um ep de uma serie")
	fmt.Println("")

	fmt.Print("Escolha um item no menu: ")
	fmt.Scan(&item)

	switch item {
	case 1:
		var seriesID int
		fmt.Print("Digite o ID da série: ")
		fmt.Scan(&seriesID)

		download.Serie(strconv.Itoa(seriesID))
		break

	case 2:
		var seriesID int
		var seasonNum int

		fmt.Print("Digite o ID da série: ")
		fmt.Scan(&seriesID)

		fmt.Print("Digite o número da temporada: ")
		fmt.Scan(&seasonNum)

		download.SerieSeason(strconv.Itoa(seriesID), strconv.Itoa(seasonNum))
		break

	case 3:
		var seriesID int
		var seasonNum int
		var episodeNum int

		fmt.Print("Digite o ID da série: ")
		fmt.Scan(&seriesID)

		fmt.Print("Digite o número da temporada: ")
		fmt.Scan(&seasonNum)

		fmt.Print("Digite o número do episodio: ")
		fmt.Scan(&episodeNum)

		download.SerieSeasonEpisode(
			strconv.Itoa(seriesID),
			strconv.Itoa(seasonNum),
			strconv.Itoa(episodeNum),
		)
		break
	default:
		fmt.Println(":Opção desconhecida:")
	}

}
