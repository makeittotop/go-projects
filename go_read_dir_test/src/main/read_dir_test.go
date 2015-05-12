package main

import (
    "fmt";
    "testing";
    "io/ioutil";
    "log";
)

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

func BenchmarkReadDir(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		list_dir("/home/abhishek/dev");
	}
}

func BenchmarkFooBar(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		fmt.Sprintf("Hello");
	}
}