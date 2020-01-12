package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type gRes struct {
	Code   int    `json:code`
	ImgUrl string `json:imgurl`
}

/**
获取头像url
*/
func GetAvaterUrls(size int) []string {
	var result []string
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(size)

	for i := 0; i < size; i++ {
		go func() {
			res, err := http.Get("https://api.uomg.com/api/rand.avatar?format=json")
			defer res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}

			var jsonResult gRes
			err = json.NewDecoder(res.Body).Decode(&jsonResult)
			if err != nil {
				log.Fatal(err)
			}
			mutex.Lock()
			result = append(result, jsonResult.ImgUrl)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return result
}
