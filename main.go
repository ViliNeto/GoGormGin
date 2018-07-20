package main

//Database gotuto

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"encoding/json"
)

type Login struct {
	User string
	Password string
}

//Função usando net/http do Go
func rootPage(w http.ResponseWriter, r *http.Request) {
	//GET simples
	//fmt.Fprintf(w, "Welcome to this page!")

	//GET com URL Path
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	//POST
	decoder := json.NewDecoder(r.Body)

	var data Login
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Hi there, I love %s!", data)

}

func getting(c *gin.Context){

	//Simples
	c.String(200, "Recebido")

	//Caso venha algo no Path
	user := c.Param("User")
	if user == "" {
		user = c.Query("User")
	}
	c.String(http.StatusOK, "\nHello %s", user)
}

func posting(c *gin.Context){
	var json Login
	if err := c.ShouldBindJSON(&json); err == nil {
		if json.User == "vili" && json.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func main() {
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	//router := gin.Default()
	//gin.SetMode("release")

	gin.SetMode("release")

	router := gin.New()

	// Install the default logger, not required
	router.Use(gin.Logger())

	//Se usando gin.New() ou seja sem o default middleware chamar o Recovery para garantir que o serviço não caia
	router.Use(gin.Recovery())

	////GET ou POST usando net/http
	//http.HandleFunc("/", rootPage)
	//log.Fatal(http.ListenAndServe(":8081", nil))

	//GET e POST usando GIN

	//GET Simples ou Query
	router.GET("/someGet", getting)

	//GET Path
	router.GET("/someGet/:User", getting)

	//POST
	router.POST("/somePost", posting)

	router.Run(":8080")



}