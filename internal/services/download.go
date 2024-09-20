package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/aurennunes/iptv/internal/models"
)

type Download struct {
	Xtream  *Xtream
	BaseDir string
}

func NewDownload(xtream *Xtream, baseDir string) *Download {

	return &Download{
		Xtream:  xtream,
		BaseDir: baseDir,
	}
}

func (d *Download) createDirIfNotExists(dir string) error {
	destinPath := path.Join(d.BaseDir, dir)

	if _, err := os.Stat(destinPath); os.IsNotExist(err) {
		return os.MkdirAll(destinPath, 0755)
	}
	return nil
}

func (d *Download) downloadEpisode(episode models.Episode, seasonDir string) error {
	resp, err := http.Get(d.Xtream.BuildUrl("series", episode.ID, episode.ContainerExtension))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fileName := strings.ReplaceAll(episode.Title, " ", "_") + ".mp4"
	filePath := path.Join(d.BaseDir, seasonDir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	totalBytes, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	var bytesDownloaded int64
	buf := make([]byte, 4096)
	startTime := time.Now()

	printProgress := func(currentBytes int64) {
		percent := float64(currentBytes) / float64(totalBytes) * 100
		bars := int(percent / 2)
		fmt.Printf("\r[%-50s] %.2f%%", strings.Repeat("=", bars), percent)
	}

	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			file.Write(buf[:n])
			bytesDownloaded += int64(n)
			printProgress(bytesDownloaded)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	// _, err = io.Copy(file, resp.Body)
	fmt.Printf("\nDownload completo em %v\n", time.Since(startTime))
	return nil
}

func (d *Download) Serie(ID string) {
	serieDetails, err := d.Xtream.GetSeriesDetails(ID)

	if err != nil {
		fmt.Printf("Erro ao obter informações da série: %v\n", err)
		return
	}

	// Cria uma pasta para a série
	seriesDir := serieDetails.Info.Name
	err = d.createDirIfNotExists(seriesDir)
	if err != nil {
		fmt.Printf("Erro ao criar pasta da série: %v\n", err)
		return
	}

	fmt.Printf("Baixando série: %s\n", serieDetails.Info.Name)

	for season, episodes := range serieDetails.Episodes {
		seasonDir := path.Join(seriesDir, "Temporada_"+season)

		err := d.createDirIfNotExists(seasonDir)
		if err != nil {
			fmt.Printf("Erro ao criar pasta da temporada: %v\n", err)
			continue
		}

		for _, episode := range episodes {
			fmt.Printf("Baixando episódio: %s (Temporada %s)\n", episode.Title, season)
			if err := d.downloadEpisode(episode, seasonDir); err != nil {
				fmt.Printf("Erro ao baixar episódio %s: %v\n", episode.Title, err)
			} else {
				fmt.Printf("Episódio %s baixado com sucesso!\n", episode.Title)
			}
		}
	}
}

func (d *Download) SerieSeason(ID, seasonNum string) {
	serieDetails, err := d.Xtream.GetSeriesDetails(ID)

	if err != nil {
		fmt.Printf("Erro ao obter informações da série: %v\n", err)
		return
	}

	// Cria uma pasta para a série
	seriesDir := serieDetails.Info.Name
	err = d.createDirIfNotExists(seriesDir)
	if err != nil {
		fmt.Printf("Erro ao criar pasta da série: %v\n", err)
		return
	}

	fmt.Printf("Baixando série: %s\n", serieDetails.Info.Name)

	if episodes := serieDetails.Episodes[seasonNum]; len(episodes) > 0 {
		seasonDir := path.Join(seriesDir, "Temporada_"+seasonNum)

		err := d.createDirIfNotExists(seasonDir)
		if err != nil {
			fmt.Printf("Erro ao criar pasta da temporada: %v\n", err)
			return
		}

		for _, episode := range episodes {
			fmt.Printf("Baixando episódio: %s (Temporada %s)\n", episode.Title, seasonNum)
			if err := d.downloadEpisode(episode, seasonDir); err != nil {
				fmt.Printf("Erro ao baixar episódio %s: %v\n", episode.Title, err)
			} else {
				fmt.Printf("Episódio %s baixado com sucesso!\n", episode.Title)
			}
		}
	}
}

func (d *Download) SerieSeasonEpisode(ID, seasonNum, episodeNum string) {
	serieDetails, err := d.Xtream.GetSeriesDetails(ID)

	if err != nil {
		fmt.Printf("Erro ao obter informações da série: %v\n", err)
		return
	}

	// Cria uma pasta para a série
	seriesDir := serieDetails.Info.Name
	err = d.createDirIfNotExists(seriesDir)
	if err != nil {
		fmt.Printf("Erro ao criar pasta da série: %v\n", err)
		return
	}

	fmt.Printf("Baixando série: %s\n", serieDetails.Info.Name)

	if episodes := serieDetails.Episodes[seasonNum]; len(episodes) > 0 {
		seasonDir := path.Join(seriesDir, "Temporada_"+seasonNum)

		err := d.createDirIfNotExists(seasonDir)
		if err != nil {
			fmt.Printf("Erro ao criar pasta da temporada: %v\n", err)
			return
		}

		for _, episode := range episodes {
			if episode.EpisodeNum == episodeNum {
				fmt.Printf("Baixando episódio: %s (Temporada %s)\n", episode.Title, seasonNum)
				if err := d.downloadEpisode(episode, seasonDir); err != nil {
					fmt.Printf("Erro ao baixar episódio %s: %v\n", episode.Title, err)
				} else {
					fmt.Printf("Episódio %s baixado com sucesso!\n", episode.Title)
				}
			}
		}
	}
}
