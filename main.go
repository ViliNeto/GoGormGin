package main

import (
	"./apiGin"
	"./apiGo"
	"net/http"
)

func main() {

	router := apiGin.StartRouter()

	////GET ou POST usando net/http
	http.HandleFunc("/", apiGo.RootPage)

	//GET Simples ou Query
	router.GET("/someGet", apiGin.Getting)

	//GET Path
	router.GET("/someGet/:User", apiGin.Getting)

	//POST
	router.POST("/somePost", apiGin.Posting)

	//REGISTER POST
	router.POST("/register", apiGin.Register)

	//Router Go
	//log.Fatal(http.ListenAndServe(":8081", nil))
	//Router Gin
	router.Run(":8080")

}
