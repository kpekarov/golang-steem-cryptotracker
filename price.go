package main

import (
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type Result map[string]float64

func main() {
	doEvery(time.Second * 10, getPrice)
}

func getPrice(t time.Time) {
	From := "STEEM"
	To := "USD"
	url := "https://min-api.cryptocompare.com/data/price?fsym=" + From + "&tsyms=" + To;

	w := http.Client{
		Timeout: time.Second * 2,
	}

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	res, getErr := w.Do(req)
	if getErr != nil {
		log.Println(getErr)
	} else {

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Println(readErr)
		} else {

			r := Result{}
			jsonErr := json.Unmarshal(body, &r)
			if jsonErr != nil {
				log.Println(jsonErr)
			} else {

				for _, p := range r {
					fmt.Printf("%v: 1 "+From+" = "+strconv.FormatFloat(p, 'f', 6, 64)+" "+To+"\n", t)
				}
			}
		}
	}
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}
