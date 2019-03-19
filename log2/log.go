package log2

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

//const requestIDKey  = 42
type key int
const requestIDKey  = key(42)
func Print(ctx context.Context, msg string)  {
	id,ok := ctx.Value(requestIDKey).(int64)
	if !ok{
		log.Panicln("could not find request ID in context")
		return
	}
	log.Printf("[%d] %s",id,msg)
}

func Decorate(f http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		ctx := r.Context()
		id := rand.Int63()
		ctx =  context.WithValue(ctx, requestIDKey, id)
		r = r.WithContext(ctx)
		f(w, r)
	}
}