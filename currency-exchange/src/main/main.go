package main

import (
  "net/http"
  "fmt"
  "time"
  "io/ioutil"
  "encoding/xml"
  "regexp"
  "strings"
)

// Structure - top
type XMLRss struct {
	XMLName  xml.Name  `xml:"rss"`
	Channel XMLChannel `xml:"channel"`
	/*
	Title string `xml:"title"`
	Link string `xml:"link"`
	PubDate string `xml:"pubDate"`
	Description string `xml:"description"`
	Category string `xml:"category"`
	*/
}

// Structure - One level down
type XMLChannel struct {
	XMLName  xml.Name  `xml:"channel"`
	Items    []XMLItem `xml:"item"`
}

// Structure - Two level down
type XMLItem struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	PubDate string `xml:"pubDate"`
	Description string `xml:"description"`
	Category string `xml:"category"`
}

func main() {
	start := time.Now()
	
	// Compile the required regex
	r, err := regexp.Compile(`(\d+).*\s+(\d+\.\d+).*`)
	if err != nil {
		panic(err)
	}
	
	url := "http://themoneyconverter.com/rss-feed/USD/rss.xml"
	// Get the response stream off the web
	resp, err := http.Get(url)
	if err != nil {
		panic(err);
	}
	
	// Parse the response stream into into a byte buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	// Unmarshal the byte buffer into an XML stream
	v := XMLRss{}	
	err = xml.Unmarshal(body, &v)
	if err != nil {
		panic(err)
	}
	
	// 
	//fmt.Printf("XMLName: %#v\n", v.XMLName.Local)
	items := v.Channel.Items
	fmt.Printf("Total items: %d\n", len(items))
	
	
	for _, value := range items {
		//fmt.Printf("Item # %d: \n", count + 1)
		//fmt.Println(value.Title, value.Description)
		description := strings.Replace(value.Description, ",", "", -1)
		curr_items := strings.Split(value.Title, "/")
		matches := r.FindStringSubmatch(description)
		fmt.Printf("%s/%s : %s/%s\n",curr_items[1], curr_items[0],  matches[1], matches[2])
	}
	
	fmt.Printf("Total time taken: %s\n", time.Since(start).String())
}

