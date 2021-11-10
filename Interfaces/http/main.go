package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	/* The same thing, but uses Writer interface (with Write(p []byte) int64 error
	function inside)
	in os.Stdout and Reader interface (with Read(p []byte) int error function inside)
	Copy just make a copy some data from source to Writer

	io.Copy(os.Stdout, resp.Body) */

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote this many bytes: %v\n", len(bs))
	return len(bs), nil
}
