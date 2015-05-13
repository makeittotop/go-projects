package main

import (
    "fmt"
    "net/http"
    //"log"
    "net/url"
    //"io/ioutil"
    "bytes"
    "encoding/json"
)

type Host_Data struct {
	url string
	path string
}

type Resp_Data struct {
	Code int64 `json:"code"`
	Message string `json:"message"`
	More_Info string `json:"more_info"`
	Status int `json:"status"`
}

/*
type myjar struct {
    jar map[string] []*http.Cookie
}

func (p* myjar) SetCookies(u *url.URL, cookies []*http.Cookie) {
    fmt.Printf("The URL is : %s\n", u.String())
    fmt.Printf("The cookie being set is : %s\n", cookies)
    p.jar [u.Host] = cookies
}

func (p *myjar) Cookies(u *url.URL) []*http.Cookie {
    fmt.Printf("The URL is : %s\n", u.String())
    fmt.Printf("Cookie being returned is : %s\n", p.jar[u.Host])
    return p.jar[u.Host]
}
*/

func main() {
	/*
		package main
		
		import (
		    "fmt"
		    "net/url"
		)
		
		func main() {
		
		    var Url *url.URL
		    Url, err := url.Parse("http://www.example.com")
		    if err != nil {
		        panic("boom")
		    }
		
		    Url.Path += "/some/path/or/other_with_funny_characters?_or_not/"
		    parameters := url.Values{}
		    parameters.Add("hello", "42")
		    parameters.Add("hello", "54")
		    parameters.Add("vegetable", "potato")
		    Url.RawQuery = parameters.Encode()
		
		    fmt.Printf("Encoded URL is %q\n", Url.String())
		}
    */
	
	var post_data = url.Values{}
	post_data.Add("To", "+971528703911")
	post_data.Add("From", "+") //"+14805265812")
	post_data.Add("Body", "test")
	
	//fmt.Println(post_data)
	//fmt.Println(post_data.Encode())
	
	var post_data_2 = url.Values{}
	post_data_2.Add("StatusCallback", "http://106.187.50.144:9999/twilio")
	
	data := post_data.Encode() + string("&StatusCallback=http://106.187.50.144:9999/twilio")
	//fmt.Println(data)

	buf := bytes.NewBufferString(data)
	fmt.Println(buf)

	
	//host_data := Host_Data{ "https://api.twilio.com", "/2010-04-01/Accounts/AC748702fc738d6f16455dd418b93a67d1/Messages/SM48645fe79370463797ea76ae70ae4897.json" }
	
	client := &http.Client{}
	
	req, err := http.NewRequest("POST", "https://api.twilio.com/2010-04-01/Accounts/AC748702fc738d6f16455dd418b93a67d1/Messages.json", buf)
	req.Host="api.twilio.com"
	//req.Header.Add("auth", "8C748702fc738d6f16455dd418b93a67d1:5b6293592a8c4a8e6ba2a36ea6ce1ecA")
	req.SetBasicAuth("AC748702fc738d6f16455dd418b93a67d1", "5b6293592a8c4a8e6ba2a36ea6ce1ec8")
	//req.Header.Add("User-Agent", "curl/7.19.7 (x86_64-redhat-linux-gnu) libcurl/7.19.7 NSS/3.14.0.0 zlib/1.2.3 libidn/1.18 libssh2/1.4.2")
	//req.Header.Add("Accept", "*/*")
	//req.Header.Add("Content-Length", "86")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//req.ContentLength = int64(buf.Len())
	
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	
	// io reader to string
	resp_buf := new(bytes.Buffer)
	resp_buf.ReadFrom(resp.Body)
	
	if resp.StatusCode == 201 {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failure!")
	}
	
	// Convert string into json	
	fmt.Println(resp_buf)
	
	var resp_data Resp_Data
	err = json.Unmarshal([]byte(resp_buf.String()), &resp_data)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%+v\n", resp_data)
	fmt.Printf("%d\n", resp_data.Code)
	fmt.Printf("%s\n", resp_data.Message)
	fmt.Printf("%s\n", resp_data.More_Info)
	fmt.Printf("%d\n", resp_data.Status)
	
	/*
	resp, err := http.Post(host_data.url, host_data.path, buf)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(resp)
	}
	*/
	
}

