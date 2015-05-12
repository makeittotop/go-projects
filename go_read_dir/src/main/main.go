package main

import (
    "fmt";
    "io/ioutil";
    "log";
)

func main() {
	fmt.Println("Reading dir ~abhishek/dev...");
	
	list_dir("/home/abhishek/dev")
}

func list_dir(dir_name string) {
	fmt.Printf("Dir -- %s\n", dir_name);
	
	file_list, err := ioutil.ReadDir(dir_name);
	
	if err == nil {
			for _, file_handle := range file_list {
				if file_handle.IsDir() {
					//fmt.Printf("Dir -- %s\n", file_handle.Name());
					dir_path := dir_name + "/" + file_handle.Name();
					list_dir(dir_path);
				} else {
					fmt.Printf("File in %s -- %s\n", dir_name, file_handle.Name());
				}
			}
		} else {
			log.Fatal(err);
		}	
}