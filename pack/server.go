package pack

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func SetMyCookie(response http.ResponseWriter) {
	cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue"}
	http.SetCookie(response, &cookie)
}

func GenericHandler(response http.ResponseWriter, request *http.Request) {

	SetMyCookie(response)
	response.Header().Set("Content-type", "text/plain")

	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}

	fmt.Fprint(response, "FooWebHandler says ... \n")
	fmt.Fprintf(response, " request.Method     '%v'\n", request.Method)
	fmt.Fprintf(response, " request.RequestURI '%v'\n", request.RequestURI)
	fmt.Fprintf(response, " request.URL.Path   '%v'\n", request.URL.Path)
	fmt.Fprintf(response, " request.Form       '%v'\n", request.Form)
	fmt.Fprintf(response, " request.Cookies()  '%v'\n", request.Cookies())
}

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	webpage, err := ioutil.ReadFile("./template/home.html")
	if err != nil {
		http.Error(response, fmt.Sprintf("home.html file error %v", err), 500)
	}
	fmt.Fprint(response, string(webpage))
	//--------------------------------------
	request.ParseForm()
	fmt.Println("email:", request.Form["email"])
	fmt.Println("subject:", request.Form["subject"])
	fmt.Println("message:", request.Form["message"])
}

func Server_start() {
	port := 8080
	portstring := strconv.Itoa(port)

	mux := http.NewServeMux()
	mux.Handle("/home", http.HandlerFunc(HomeHandler))
	mux.Handle("/generic/", http.HandlerFunc(GenericHandler))

	log.Print("Listening on port " + portstring + " ... ")
	err := http.ListenAndServe(":"+portstring, mux)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
