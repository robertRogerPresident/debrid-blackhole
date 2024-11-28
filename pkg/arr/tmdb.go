package arr

import (
	"encoding/json"
	"net/http"
	url2 "net/url"
)

type TMDBResponse struct {
	Page    int `json:"page"`
	Results []struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		MediaType  string `json:"media_type"`
		PosterPath string `json:"poster_path"`
	} `json:"results"`
}

func SearchTMDB(term string) (*TMDBResponse, error) {
	resp, err := http.Get("https://api.themoviedb.org/3/search/multi?api_key=21564df127bf5aaec576d2a8975958a4&query=" + url2.QueryEscape(term))
	if err != nil {
		return nil, err
	}

	var data *TMDBResponse
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
