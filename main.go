package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sriosdev/zipper"
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
	// Get flags parameters
	path := flag.String("f", "", "File path")
	port := flag.Uint("p", 3000, "Server port")
	flag.Parse()

	if len(*path) == 0 {
		fmt.Println("The file path is required")
		flag.PrintDefaults()
		os.Exit(2)
	}

	file, err := os.Open(*path)
	if err != nil {
		fmt.Println("Can't open the file or does not exist")
		os.Exit(1)
	}

	fi, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	// Directories are compressed into a ZIP file
	if fi.IsDir() {
		fmt.Println("\nCompressing directory...")

		file, err = zipper.ZipFolder(file)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}

		fmt.Println("Directory compressed successfuly!")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileName := filepath.Base(file.Name())
		w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
		http.ServeFile(w, r, file.Name())

		fmt.Println("Finishing...")
		os.Exit(0)
	})

	addr := buildAddr(*port)

	fmt.Println("\n  Your file is waiting at:")
	fmt.Print("  - Download link: ")
	fmt.Printf("\033[1;36m%s\033[0m", addr)
	fmt.Println("\n\n  Remember you can only access link once.\n")
	log.Fatalln(http.ListenAndServe(addr, nil))
}
