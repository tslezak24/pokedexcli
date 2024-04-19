package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreArea(area string) Encounters {
	exploreURL := url + "/location-area/" + area

	if val, ok := c.cache.GetFromCache(exploreURL); ok {
		response := Encounters{}
		err := json.Unmarshal(val, &response)
		if err != nil {
			fmt.Println("unable to unmarshal JSON")
		}

		return response
	}

	res, err := http.Get(exploreURL)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Println(res.StatusCode)
	}
	if err != nil {
		fmt.Println(err)
	}

	var response Encounters

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println(err)
	}

	return response
}
