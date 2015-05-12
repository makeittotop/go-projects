package main

import (
    "net/http";
    "log";
    /*
    "fmt";
    //"html";
    //"io/ioutil";
    "os";
    "strings";
    "bufio";
    */
)

func main() {
	/*
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	    //fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path));
	    w.Write([]byte("Hello,  " + html.EscapeString(r.URL.Path)));
	});
	
	
	http.Handle("/", new(my_handler));
	*/
	
	log.Fatal(http.ListenAndServe("localhost:7778", http.FileServer(http.Dir("public"))));
}

/*
type my_handler struct {
	http.Handler;
}

func (mh *my_handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := "public" + r.URL.Path;
	
	//wd, err := os.Getwd();
	//fmt.Printf("Current Working Dir - %s\n", wd);
	
	fmt.Println(path);
	
	//data, err := ioutil.ReadFile(string(path));
	f, err := os.Open(string(path));
	
	if err == nil {
		buffered_reader := bufio.NewReader(f);
		
		var content_type string;
		
		if strings.HasSuffix(path, ".png") {
			content_type = "image/png";
		} else if strings.HasSuffix(path, ".html") {
			content_type = "text/html";
		} else if strings.HasSuffix(path, ".css") {
			content_type = "text/css";
		} else if strings.HasSuffix(path, ".js") {
			content_type = "application/javascript";
		} else if strings.HasSuffix(path, ".mov") {
			content_type = "video/mov";
		} else if strings.HasSuffix(path, ".mp4") {
			content_type = "video/mp4";
		} else {
			content_type = "text/plain";
		}
		
		w.Header().Add("Content Type", content_type);
		
		//w.Write(data);
		buffered_reader.WriteTo(w);
	} else {
		w.WriteHeader(404);
		w.Write([]byte("404 - " + http.StatusText(404)));
	}
}
*/