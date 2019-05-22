package main

// mengimport package "net/http"
import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// tanda "/" tanda main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// tampilkan text
		fmt.Fprintf(w, "Hello, nama Edy Purnomo")
	})

	//  "/about" artinya berada dihalaman http://localhost/About
	http.HandleFunc("/About", func(w http.ResponseWriter, r *http.Request) {
		// tampilkan text
		fmt.Fprintf(w, "Belajar routing")
	})

	http.HandleFunc("/Waktu", func(w http.ResponseWriter, r *http.Request) {
		var time1 = time.Now()
		// tampilkan Waktu
		fmt.Printf("Waktu saat ini %v\n", time1)
	})

	// port yg digunakan
	http.ListenAndServe(":8081", nil)
}
