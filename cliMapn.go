package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func mapn(conf *config) error {
	req, err := http.NewRequest("GET", conf.next, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	pokeRes := pokeResponse{}
	err = json.Unmarshal(body, &pokeRes)
	if err != nil {
		return err
	}
	conf.next = pokeRes.Next
	conf.previous = pokeRes.Previous

	for _, result := range pokeRes.Results {
		fmt.Println(result.Name)
	}

	return nil
}
