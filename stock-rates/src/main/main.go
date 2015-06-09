package main

import (
  "fmt"
  "time"
  "net/url"
  "log"
  "net/http"
  "io/ioutil"
  "encoding/xml"
  // runtime
  "runtime"
)

type XMLStockQuote struct {
	XMLName  xml.Name  `xml:"StockQuote"`
	Symbol string `xml:"Symbol"`
	Name string `xml:"Name"`
	LastPrice string `xml:"LastPrice"`
	Timestamp string `xml:"Timestamp"`
}

func main() {
	// give the app more cores
	runtime.GOMAXPROCS(4)
	//start the timer
	start := time.Now()
	// url to query
	url_string := "http://dev.markitondemand.com/Api/v2/Quote"
	u, err := url.Parse(url_string)
	if err != nil {
		log.Fatal(err)
	}
	// list the symbols to stock quote
	symbols := []string { "goog", "aapl", "ibm", "ms", "msft", "fb", }
	//
	count := 0
	for _, symbol := range symbols {
		go func(symbol string) {
			start_time := time.Now()
			// add the query string
			q := u.Query()
			q.Set("symbol", symbol)
			u.RawQuery = q.Encode()
			fmt.Println("Query: ", symbol)
			// get the response 
			resp, err := http.Get(u.String())
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println(resp.Body)
			// save the response in a byte slice
			fmt.Println("Success: ", symbol)
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println(body)
			// declare the XMLString interface and unmarshal body into it
			v := XMLStockQuote{}
			err = xml.Unmarshal(body, &v)
			if err != nil {
				log.Fatal(err)
			}	
			
			fmt.Println(v.Name, v.LastPrice, "@", v.Timestamp)
			count += 1
			fmt.Println("Done: ", symbol)
			fmt.Printf("Time elapsed for %s : %s\n", symbol, time.Since(start_time).String())
		}(symbol)
	}
	
	i := 0
	for {
		i = count
		fmt.Println("i, count: ", i, count)
		if i >= len(symbols) {
			break
		} else {
			dur, _ := time.ParseDuration("10ms")
			time.Sleep(dur)
		}
	}
	
	fmt.Printf("Total time taken: %s\n", time.Since(start).String())
}
