package commands

import (
	"error"
	"fmt"
)

func CommandMapf(conf *repl.config, args ...string) error {
	locationsResp, err := conf.pokeapiClient.ListLocations(conf.nextLocationsURL)
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

func CommandMapb(conf *repl.config, args ...string) error {
	if conf.prevLocationsURL == nil {
		return errors.New("You are on the first page")
	}

	locationsResp, err := conf.pokeapiClient.ListLocations(conf.prevLocationsURL)
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
