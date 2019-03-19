package main

import (
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main(){
	// sending cancelation signal through request
	ctxBackground := context.Background()
	ctx,cancel := context.WithTimeout(ctxBackground, 5*time.Second)
	defer cancel()

	req,err:=http.NewRequest(http.MethodGet, "http://localhost:8080",nil)
	if err!= nil{
		log.Fatal("http.NewRequest:",err)
	}
	req = req.WithContext(ctx)
	res,err :=http.DefaultClient.Do(req)
	// If server .Do(req) take more than 4 second to receive HTTP response from client then Timeout occur and it returns error
	// saying http.DefaultClient.Do:Get http://localhost:8080: context deadline exceeded
	if err!= nil{
		log.Fatal("http.DefaultClient.Do:",err)
		//2019/02/04 16:55:27 http.DefaultClient.Do:Get http://localhost:8080: context deadline exceeded
	}
	/* // sending manual cancelation signal
	// press ctrl + c to terminate the client before server send response, indeed it will send a cancelation signal to server
	res,err := http.Get("http://localhost:8080")
	if err != nil{
		log.Fatal(err)
	}*/
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK{
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body)
}