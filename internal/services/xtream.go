package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aurennunes/iptv/internal/models"
)

type Xtream struct {
	URL      string      `json:"url"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Auth     models.Auth `json:"auth"`
}

func NewXtream(url, username, password string) *Xtream {
	return &Xtream{
		URL:      url,
		Username: username,
		Password: password,
	}
}

func (x *Xtream) BuildUrl(Type, Id, Ext string) string {
	url := fmt.Sprintf("%s/%s/%s/%s/%s.%s", x.URL, Type, x.Username, x.Password, Id, Ext)
	return url
}

func (x *Xtream) Authenticate() (*models.Auth, error) {
	url := fmt.Sprintf("%s/player_api.php?username=%s&password=%s&action=user", x.URL, x.Username, x.Password)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var auth models.Auth
	if err := json.NewDecoder(resp.Body).Decode(&auth); err != nil {
		return nil, err
	}
	return &auth, nil
}

func (x *Xtream) GetMovieCategories() ([]*models.Category, error) {
	url := fmt.Sprintf("%s/player_api.php?username=%s&password=%s&action=get_vod_categories", x.URL, x.Username, x.Password)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var categories []*models.Category
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func (x *Xtream) GetSerieCategories() ([]*models.Category, error) {
	url := fmt.Sprintf("%s/player_api.php?username=%s&password=%s&action=get_series_categories", x.URL, x.Username, x.Password)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var categories []*models.Category
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func (x *Xtream) GetSerieByCategory(ID string) ([]*models.Serie, error) {
	url := fmt.Sprintf("%s/player_api.php?username=%s&password=%s&action=get_series&category_id=%s", x.URL, x.Username, x.Password, ID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var series []*models.Serie
	if err := json.NewDecoder(resp.Body).Decode(&series); err != nil {
		return nil, err
	}
	return series, nil
}

func (x *Xtream) GetSeriesDetails(ID string) (*models.SerieDetails, error) {
	url := fmt.Sprintf("%s/player_api.php?username=%s&password=%s&action=get_series_info&series_id=%s", x.URL, x.Username, x.Password, ID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var serie *models.SerieDetails
	if err := json.NewDecoder(resp.Body).Decode(&serie); err != nil {
		return nil, err
	}
	return serie, nil
}

func (x *Xtream) GetEpisodesBySeason(SeriesId, SeasonNumber string) (*[]models.Episode, error) {
	serie, err := x.GetSeriesDetails(SeriesId)
	if err != nil {
		return nil, err
	}

	season := serie.Episodes[SeasonNumber]
	if len(season) == 0 {
		return nil, fmt.Errorf("Temporada %s não encontrada na série %s", SeasonNumber, SeriesId)
	}

	return &season, nil
}

func (x *Xtream) GetSerieStreamUrl(SeriesId, SeasonNumber, EpisodeNumber string) (string, error) {
	serie, err := x.GetEpisodesBySeason("50", "1")
	if err != nil {
		return "", err
	}

	for _, episode := range *serie {
		if episode.EpisodeNum == EpisodeNumber {
			return x.BuildUrl(
				"series",
				episode.ID,
				episode.ContainerExtension,
			), nil
		}
	}

	return "", fmt.Errorf("Episódio %s não encontrado na temporada %s da série %s", EpisodeNumber, SeasonNumber, SeriesId)
}
