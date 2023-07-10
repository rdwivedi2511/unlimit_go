
package api

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "time"
	"encoding/json"
	"context"
    "errors"
)

type Activity struct {
    Key string `json:"key"`
    Activity string `json:"activity"`
}


func Call(ch chan Activity) {
  for i := 0; i < 3; i++ {
 client := &http.Client{}
        req, err := http.NewRequest(http.MethodGet, "http://www.boredapi.com/api/activity", nil)
        if err != nil {
           fmt.Print(err.Error())
           os.Exit(1)
        }
        ctx, cancel := context.WithTimeout(req.Context(), 2*time.Second)
        defer cancel()
        req = req.WithContext(ctx)
        response, err := client.Do(req)
        // call server
        if errors.Is(err, context.DeadlineExceeded) {
            log.Println("Activity-API not available")
        }
        if os.IsTimeout(err) {
            log.Println("IsTimeoutError: true")
        }
        if err != nil {
            log.Fatal(err)
        }
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
 
 x := map[string]string{}

	json.Unmarshal([]byte(string(responseData)), &x)
	fmt.Println(x)
	
	
		var p Activity
	for k, v := range x {
		switch k {
		case "key":
			p.Key = v
		case "activity":
			p.Activity = v
		}
	}
	fmt.Printf("%+v\n", p)
	  ch <- p
	//return p
	fmt.Println("1 capacity is", cap(ch))
    fmt.Println("1 length is", len(ch))
	}
	 close(ch)
}

