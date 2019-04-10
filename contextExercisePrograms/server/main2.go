package main

import (
	"fmt"
	log22 "howtogo/log2"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/context"
)

func main() {
	//http.HandleFunc("/",handler)
	http.HandleFunc("/", log22.Decorate(handler))
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, int(42), int64(100))
	/*Output: ctx = context.WithValue(ctx, int(42), int64(100))
	  Here because of key 42, the random id value is overwritten with 100. To aviod this collision
	  make this value 42 of a type key that we do not export from log like:
	  type key int
	  const requestIDKey  = key(42)
	  Note: Now requestIDKey is type key only log package can use it.
	  2019/02/04 19:16:17 [100] handler started
	  2019/02/04 19:16:22 [100] handler ended
	*/
	log22.Print(ctx, "handler started")
	defer log22.Print(ctx, "handler ended")
	select {
	case <-time.After(5 * time.Second):
		//case <-time.After(4*time.Second):
		// After sleep of duration "hello" msg will be send to client
		fmt.Fprint(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		log22.Print(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
Output:
E:\workspace\TDD\gorilla_mux\src\GCP\server>go run main2.go
2019/02/04 19:02:43 [5577006791947779410] handler started
2019/02/04 19:02:49 [5577006791947779410] handler ended

*/
