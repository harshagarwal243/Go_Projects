package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
	 _ "github.com/go-sql-driver/mysql"
	 "github.com/satori/go.uuid"
	 "github.com/julienschmidt/httprouter"
	 "net/http"
)

func main() {
	pass := "harsh"
	b ,_ := bcrypt.GenerateFromPassword([]byte(pass),bcrypt.MinCost)
	s := string(b)
	fmt.Println(s)
	u , _:= uuid.NewV4()
	fmt.Println(u.String())

	mux := httprouter.New()
	mux.GET("/",index)
	http.ListenAndServe(":3031",mux)
}

func index(w http.ResponseWriter,r *http.Request,_ httprouter.Params){
	fmt.Fprintf(w,"<h1>Hello how are you</h1>")
}