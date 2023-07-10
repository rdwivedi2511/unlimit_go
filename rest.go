package main

import (
    "fmt"
    "log"
    "net/http"
	"api"
	"encoding/json"
	"db"
	//"context"
	 "golang.org/x/time/rate"
)

func rateLimiterMiddleware(next func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
    limiter := rate.NewLimiter(3, 3) 
    return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
        if !limiter.Allow() {
            writer.Write([]byte("rate limit exceeded "))
            return
        } else {
            returnAllActivities(writer, request)
        }
    })
}


func returnAllActivities(w http.ResponseWriter, r *http.Request){
 
    fmt.Println("Endpoint Hit: returnAllActivities")
	
	var activities []api.Activity
	ch := make(chan api.Activity, 3)
	 go api.Call(ch)
	
	 for v := range ch {
	    db.Save(v)
		activities = append(activities,v )
		fmt.Println("added value in Array", activities)
        fmt.Println("read value ", v,"from ch")
		
    } 

	fmt.Println("final value in Array", activities)
	 json.NewEncoder(w).Encode(activities)
}


func handleRequests() {
    fmt.Println("hit url from browser -- url http://localhost:8080/activities")
	http.HandleFunc("/activities", rateLimiterMiddleware(returnAllActivities))
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}