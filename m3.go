package main

import (
	"fmt"
	"log"
	"net/http"

	mx "github.com/gorilla/mux"
)

func main() {
	r := mx.NewRouter()

	// Only matches if domain is "www.example.com".
	//r.HandleFunc("/products/{key}", ProductsHandler).Host("www.example.com")
	// Matches a dynamic subdomain.
	//r.Host("{subdomain:[a-z]+}.example.com")

	// Only matches if domain is "localhost".
	/*r.Host("localhost")
	r.Methods("GET")
	r.HandleFunc("/products/{key}", ProductsHandler)*/
	//OR
	r.HandleFunc("/products/{key}", ProductsHandler).Host("localhost").Methods("GET")

	/*r.HandleFunc("/products", ProductsHandler).
	Host("localhost").
	Methods("GET").
	Schemes("http")*/
	// Host > Methods > Schemes > HandleFunc(HandleFunc invked only when pattern matched)

	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/products/{key}", ProductsHandler).Methods("GET")
	//http.HandleFunc("/products/csdc", ProductsHandler)
	log.Fatalln(http.ListenAndServe("localhost:8080", r))
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("Request:%+v\n", r)
	fmt.Println("Inside ProductsHandler")
	vars := mx.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Key: %v\n", vars["key"])
	w.Write([]byte("Inside ProductsHandler"))
}

/*
fmt.Printf("Request:%+v\n", r)
Request:
&{
 Method:GET
 URL:/products/csdc
 Proto:HTTP/1.1
 ProtoMajor:1
 ProtoMinor:1
 Header:map[Accept-Language:[en-US,en;q=0.9]
			Connection:[keep-alive]
			User-Agent:[Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36]
			Accept:[* /*]
			Accept-Encoding:[gzip, deflate, br]]
 Body:{}
 GetBody:<nil>
 ContentLength:0
 TransferEncoding:[]
 Close:false
 Host:localhost:8080
 Form:map[]
 PostForm:map[]
 MultipartForm:<nil>
 Trailer:map[]
 RemoteAddr:127.0.0.1:60305
 RequestURI:/products/csdc
 TLS:<nil>
 Cancel:<nil>
 Response:<nil>
 ctx:0xc000048340
}
*/
