package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main()  {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080",nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	log.Print("handler started")
	defer log.Print("handler ended")
	select {
	case <-time.After(5*time.Second):
	//case <-time.After(4*time.Second):
		// After sleep of duration "hello" msg will be send to client
		fmt.Fprint(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Print("err:",err)
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
}