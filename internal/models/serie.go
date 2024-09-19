package models

import "encoding/json"

type Serie struct {
	ID             int      `json:"series_id"`
	Num            int      `json:"num"`
	Name           string   `json:"name"`
	Title          string   `json:"title"`
	Year           string   `json:"year"`
	StreamType     string   `json:"stream_type"`
	Cover          string   `json:"cover"`
	Plot           string   `json:"plot"`
	Cast           string   `json:"cast"`
	Director       string   `json:"director"`
	Genre          string   `json:"genre"`
	ReleaseDate    string   `json:"release_date"`
	LastModified   string   `json:"last_modified"`
	Rating         string   `json:"rating"`
	Rating5based   float32  `json:"rating_5based"`
	BackdropPath   []string `json:"backdrop_path"`
	YoutubeTrailer string   `json:"youtube_trailer"`
	EpisodeRunTime string   `json:"episode_run_time"`
	CategoryId     string   `json:"category_id"`
	CategoryIds    []int    `json:"category_ids"`
}

type Season struct {
	ID           int     `json:"id"`
	AirDate      string  `json:"air_date"`
	EpisodeCount int     `json:"episode_count"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	SeasonNumber int     `json:"season_number"`
	VoteAverage  float64 `json:"vote_average"`
	Cover        string  `json:"cover"`
	CoverBig     string  `json:"cover_big"`
}

type SerieInfo struct {
	Name           string   `json:"name"`
	Title          string   `json:"title"`
	Year           string   `json:"year"`
	Cover          string   `json:"cover"`
	Plot           string   `json:"plot"`
	Cast           string   `json:"cast"`
	Director       string   `json:"director"`
	Genre          string   `json:"genre"`
	ReleaseDate    string   `json:"release_date"`
	LastModified   string   `json:"last_modified"`
	Rating         string   `json:"rating"`
	Rating5based   float64  `json:"rating_5based"`
	BackdropPath   []string `json:"backdrop_path"`
	YoutubeTrailer string   `json:"youtube_trailer"`
	EpisodeRunTime string   `json:"episode_run_time"`
	CategoryId     string   `json:"category_id"`
	CategoryIds    []int    `json:"category_ids"`
}

type EpisodeInfo struct {
	Plot         string      `json:"plot"`
	Rating       json.Number `json:"rating"`
	Season       string      `json:"season"`
	Bitrate      int         `json:"bitrate"`
	TmdbId       json.Number `json:"tmdb_id"`
	Duration     string      `json:"duration"`
	MovieImage   string      `json:"movie_image"`
	ReleaseDate  string      `json:"release_date"`
	DurationSecs int         `json:"duration_secs"`
	CoverBig     string      `json:"cover_big"`
}

type Episode struct {
	ID                 string      `json:"id"`
	EpisodeNum         string      `json:"episode_num"`
	Title              string      `json:"title"`
	ContainerExtension string      `json:"container_extension"`
	Info               EpisodeInfo `json:"info"`
	Subtitles          []string    `json:"subtitles"`
	CustomSid          string      `json:"custom_sid"`
	Added              string      `json:"added"`
	Season             int         `json:"season"`
	DirectSource       string      `json:"direct_source"`
}

type SerieDetails struct {
	Info     SerieInfo            `json:"info"`
	Seasons  []Season             `json:"seasons"`
	Episodes map[string][]Episode `json:"episodes"`
}
