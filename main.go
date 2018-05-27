package main

import (
	"net/http"
	"io"
)

func firstPage(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"I belive you,tcy")
}

func main(){
	http.HandleFunc("/",firstPage)
	http.ListenAndServe(":8000",nil)
}
