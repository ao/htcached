package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type httpObj struct {
	url string
	// headers   []string
	body      string
	retrieved time.Time
}

var cache = make(map[string]httpObj)

func httpServer(port int, res string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		fmt.Println(r.URL)
		_, exists := cache[url]
		if !exists {
			//no cached item found in memory, get a new one
			resp, err := http.Get(res + url)
			if err != nil {
				fmt.Println(err)
				fmt.Fprintf(w, "HTCached: Some error getting backend URL")
			} else {
				body, _ := ioutil.ReadAll(resp.Body)
				newHttpObj := httpObj{url: url, body: string(body), retrieved: time.Now()}
				cache[url] = newHttpObj
			}
		}
		fmt.Fprintf(w, string(cache[url].body))
	})

	sPort := strconv.Itoa(port)

	fmt.Println("HTCached started at 0.0.0.0:" + sPort)
	http.ListenAndServe(":"+sPort, nil)

}

func main() {
	var frontPort int = 80
	var cachedResource string = "http://localhost:8080"

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 2 {
		frontPort, _ = strconv.Atoi(argsWithoutProg[0])
		cachedResource = argsWithoutProg[1]
	}

	httpServer(frontPort, cachedResource)
}
