package main

import (
  "fmt"
  "time"
  "net/url"
  "log"
  "net/http"
  "io/ioutil"
  "encoding/xml"
)

type XMLStockQuote struct {
	XMLName  xml.Name  `xml:"StockQuote"`
	Symbol string `xml:"Symbol"`
	Name string `xml:"Name"`
	LastPrice string `xml:"LastPrice"`
	Timestamp string `xml:"Timestamp"`
}

func main() {
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
		func(symbol string) {
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
		}(symbol)
	}
	
	fmt.Printf("Total time taken: %s\n", time.Since(start).String())
}
