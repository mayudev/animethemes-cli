package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/spf13/cobra"
)

// GetSeason returns all anime released in a season (e.g. Spring) of a year (e.g. 2007)
func GetSeason(season string, year int, page int) *AnimeSearch {
	v := url.Values{}

	// Set query params
	v.Set("filter[year]", strconv.Itoa(year))
	v.Set("filter[season]", season)
	v.Set("page[number]", strconv.Itoa(page))

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
