package main

import "net/http"
import "fmt"

type String string

type Struct struct{
	Greeting string
	Punct string
	Who string
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, s.Greeting + s.Punct + s.Who)
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, s)
}

func main(){
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}
