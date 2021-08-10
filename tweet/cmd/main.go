package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pooriaPoorsarvi/event-driven-design-golang/tweet"
)

func main(){
	mux.NewRouter()
	fmt.Println(tweet.TweetService)
}
