package main 


import (
   
    "fmt"
    "encoding/json"
    "net/http"
    "github.com/julienschmidt/httprouter"
       

)

type JsonInput struct {
    Name string  `json:"name"`
   
}

type JsonOutput struct {
    Greeting string `json:"greeting"`
}


func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))

}   

func hello2(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

    ip := JsonInput{}
    

    json.NewDecoder(req.Body).Decode(&ip)

   

    op := JsonOutput{}
    op.Greeting = "Hello," + ip.Name + "!"

    mj , _ := json.Marshal(op)
    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(201)
    fmt.Fprintf(rw, "%s", mj)


}


func main() {
    
    router := httprouter.New()  
    router.GET("/hello/:name", hello)
    router.POST("/hello", hello2)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: router,
    }
    server.ListenAndServe()
}
