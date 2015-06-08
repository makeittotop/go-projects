package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "flag"
    "os"
)

func main() {
	flag.Parse()
	root := flag.Arg(0)
	
    dir := root + string(filepath.Separator)
        
  	file_list, err := ioutil.ReadDir(dir);
	
	if err == nil {
			for _, file_handle := range file_list {
				if file_handle.IsDir() {
					//fmt.Printf("Dir -- %s\n", file_handle.Name());
					dir_path := dir + file_handle.Name();
					fi, err := os.Stat(dir_path)
					if err != nil {
						panic(err)
					}
					fmt.Println(dir_path, fi.Size())
				}
			}
	} else {
			panic(err);
	}	
		  
    /*
    entities, err := d.Readdir(-1)
    if err != nil {
    	panic(err)
    	
        fmt.Println(err)
        os.Exit(1)
        
    }
    
    for _, f := range entities {
        if f.Mode().IsDir() {
            fmt.Println(dir + f.Name(), f., " Bytes")
        }
    }
    */
}