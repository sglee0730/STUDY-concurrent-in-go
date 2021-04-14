package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)


func main() {
	req, err := http.NewRequest("GET","https://andcloud.io", nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1000*time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("ERROR:",err)
		return
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}