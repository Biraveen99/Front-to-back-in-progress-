package package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}


func formHandler (w http.ResponseWriter, r *http.Request)
	//handles POST request when teh form is submitted
	if r.Method = http.MethodPost{
		
		//parsing the form data
		r.ParseForm()
		
		//now we need to get the "to-do" value fomr the form
		todo := r.FormValue("to-do")
		
		//bare printer ut verdien i server console
		fmt.println("recieved task:", todo)
		
		//respnderer tilbaek til nettleser(klienten)
		fmt.Fprintf(w, "task added: %s",todo)
	} else {
		//håndterer get requests
		http.ServerFile(w,r, "Front.html")
	}
