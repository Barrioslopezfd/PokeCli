package main

import (
    "fmt"
	"encoding/json"
	"io"
	"net/http"
)


func mapp(conf *config) error {
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
    fmt.Println(conf.next, conf.previous)
    return nil
}   
