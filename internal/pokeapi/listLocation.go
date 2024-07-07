package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespLocations{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, error
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocations, err
	}

	locationResp := RespLocations{}
	err := json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespLocations{}, err
	}

	c.cache.Add(url, data)
	return locationResp, nil

}
