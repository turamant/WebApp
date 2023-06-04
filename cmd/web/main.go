package main


import (
	"net/http"
	"log"
	"fmt"
	"os"

)

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello web page")
}


const LOCALHOST = "127.0.0.1"
const PORT = "4000"
func main(){
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("Starting server on %s", "4000")

	http.HandleFunc("/", home)
	
	err := http.ListenAndServe(LOCALHOST + ":" + PORT, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}