package main

import (
    "encoding/json"
    "fmt"
    "github.com/Barrioslopezfd/PokeCli/internal/pokefetch"
)


func maps(conf *config, direction string) error {
    pokeRes := pokeResponse{}
    var url string
    if direction == "previous" {
        if conf.previous == "" {
            fmt.Println("Already on the first page")
            return nil
        }
        url = conf.previous
    } else {
        url = conf.next
    }
    res,err := pokefetch.PokeFetch(url)
    if err != nil {
       return err
    }
    err = json.Unmarshal(res, &pokeRes)
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
