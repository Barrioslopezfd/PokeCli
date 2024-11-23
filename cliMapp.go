package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)


func mapp(conf *config) error {
    if conf.previous == "" {
        err := errors.New("Already on the first page")
        log.Fatal(err)
        return err
    }
    req, err := http.NewRequest("GET", conf.previous, nil)
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
