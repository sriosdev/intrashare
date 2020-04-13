package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

func buildAddr(port uint) string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := localAddr.IP.String()

	return fmt.Sprintf("%s:%d", ip, port)
}

func main() {
	filePath := flag.String("f", "", "File path")
	port := flag.Uint("p", 3000, "Server port")
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

	addr := buildAddr(*port)

	fmt.Printf("Your file is waiting at http://%s\n", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}
