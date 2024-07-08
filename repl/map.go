package repl

import (
	"errors"
	"fmt"
)

func commandMapf(conf *Config, args ...string) error {
	locationsResp, err := conf.PokeapiClient.ListLocations(conf.nextLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = locationsResp.Next
	conf.prevLocationsURL = locationsResp.Previous

	for _, local := range locationsResp.Results {
		fmt.Println(local.Name)
	}
	return nil
}

func commandMapb(conf *Config, args ...string) error {
	if conf.prevLocationsURL == nil {
		return errors.New("You are on the first page")
	}

	locationsResp, err := conf.PokeapiClient.ListLocations(conf.prevLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = locationsResp.Next
	conf.prevLocationsURL = locationsResp.Previous

	for _, local := range locationsResp.Results {
		fmt.Println(local.Name)
	}
	return nil
}
