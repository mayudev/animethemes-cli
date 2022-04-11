package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

func SearchAnime(query string) *AnimeSearch {
	v := url.Values{}

	// Set query params
	v.Set("q", query)

	resp, err := http.Get(API_ENDPOINT + "/api/anime?" + v.Encode())
	cobra.CheckErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	cobra.CheckErr(err)

	result := &AnimeSearch{}

	err = json.Unmarshal(body, &result)
	cobra.CheckErr(err)

	return result
}
