package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	filePath := flag.String("f", "", "File path")
	flag.Parse()

	if len(*filePath) == 0 {
		fmt.Println("The file path is required")
		flag.PrintDefaults()
		os.Exit(2)
	}

	file, err := os.Open(*filePath)

	if err != nil {
		fmt.Println("Can't open the file or does not exist")
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileName := filepath.Base(file.Name())
		w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
		http.ServeFile(w, r, file.Name())
	})

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
