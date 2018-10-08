package main

import (
	"net/http"

	"./apigin"
	"./apigo"
)

func main() {

	router := apigin.StartRouter()

	////GET ou POST usando net/http
	http.HandleFunc("/", apigo.RootPage)

	//GET Simples ou Query
	router.GET("/someGet", apigin.Getting)

	//GET Path
	router.GET("/someGet/:User", apigin.Getting)

	//POST
	router.POST("/somePost", apigin.Posting)

	//REGISTER POST
	router.POST("/register", apigin.Register)

	//Router Go
	//log.Fatal(http.ListenAndServe(":8081", nil))
	//Router Gin
	router.Run(":8080")

}
