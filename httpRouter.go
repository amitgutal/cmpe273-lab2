package main

import (
	"log"
    "net/http"
	"encoding/json"
	"io"
	"fmt"
	"io/ioutil"
	"github.com/gorilla/mux"
)
type Requests struct {

    Name      string    `json:"name"`
}
type Response struct {

    Greeting      string    `json:"greeting"`   
}
type Route struct {
 
    Interpreter http.HandlerFunc
    Entity string
    Function string
    Design  string    
}

type Routes []Route

var connections = Routes{
    
    Route{
	        GetRequest,
           "GetRequest",
           "GET",
           "/hello/{userInput}",
         
         },
	Route{
		    PostRequest,
		   "PostRequest",
		   "POST",
		   "/hello",
	     },
}

func Configurations() *mux.Router {

    links := mux.NewRouter().StrictSlash(true)
		 
    for _, loop := range connections {
        links.
            Methods(loop.Function).
            Path(loop.Design).
            Name(loop.Entity).
            Handler(loop.Interpreter)
    }
   
      return links
}
func GetRequest(responseWriter http.ResponseWriter, responseRequest *http.Request) {
   
    parameters := mux.Vars(responseRequest)
    request_String := parameters["userInput"]
    fmt.Fprintln(responseWriter, "Hello,", request_String)
}

func PostRequest(resp_writer http.ResponseWriter, resp_Req *http.Request) {
	
	var response Response
	var request Requests
	
	body, post_Error := ioutil.ReadAll(io.LimitReader(resp_Req.Body, 11111111))
	
	if post_Error != nil {
		
		  fmt.Println("Could not parse the Post Request")
		  panic(post_Error)
		
	}

	Body_error := resp_Req.Body.Close()
		
	if Body_error != nil {
		
		   fmt.Println("Could not parse the Body Request")
		   panic(Body_error)
	}

	Body_err := json.Unmarshal(body, &request)
		
	if Body_err != nil {
		
		   fmt.Println("Could not parse the Body Request")
		   panic(Body_error)
	}
	
	response.Greeting = "Hello, " + request.Name + "!"
	
	resp_writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	resp_writer.WriteHeader(http.StatusCreated)
	
	Writer_err := json.NewEncoder(resp_writer).Encode(response)
	
	if  Writer_err != nil {
		   
		   fmt.Println("Could not parse the Writer Request")
		   panic(Writer_err)
	}
}

func main() {
	
    links := Configurations()
	fmt.Println(" Server Running ")
    log.Fatal(http.ListenAndServe(":8080", links))
	
}