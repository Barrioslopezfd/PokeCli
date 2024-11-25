package pokefetch

import (
	"io"
	"net/http"
	"time"
	"github.com/Barrioslopezfd/PokeCli/internal/pokecache"
)

func PokeFetch(url string) ([]byte, error) {
    interval := time.Second * 5
    cache := pokecache.NewCache(interval)
    if cachedData, ok := cache.Get(url); ok {
        return cachedData, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }
    cache.Add(url, body)
    return body, nil
}
