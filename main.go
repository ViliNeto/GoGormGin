package main

//Database gotuto

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Login struct {
	User     string
	Password string
}

type User struct {
	Usersid          int64
	Username        string
	Userpassword    string `gorm:"default:''"`
	Userpasswordb64 string `gorm:"default:''"`
}

func main() {

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router := gin.Default()

	gin.SetMode("release")

	//Usando gin.New()
	//router := gin.New()

	// Se usar o gin.Default() não precisa startar o logger
	//router.Use(gin.Logger())

	//Se usando gin.New() ou seja sem o default middleware chamar o Recovery para garantir que o serviço não caia
	//router.Use(gin.Recovery())

	////GET ou POST usando net/http
	http.HandleFunc("/", rootPage)
	//log.Fatal(http.ListenAndServe(":8081", nil))

	//GET e POST usando GIN

	//GET Simples ou Query
	router.GET("/someGet", getting)

	//GET Path
	router.GET("/someGet/:User", getting)

	//POST
	router.POST("/somePost", posting)

	//REGISTER POST
	router.POST("/register", register)

	router.Run(":8080")

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
	fmt.Fprintf(w, "Fala %s!", data)

}

func getting(c *gin.Context) {

	//Simples
	c.String(200, "Recebido")

	//Caso venha algo no Path
	user := c.Param("User")
	if user == "" {
		user = c.Query("User")
	}
	c.String(http.StatusOK, "\nOlá %s", user)
}

func posting(c *gin.Context) {
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

func register(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err == nil {
		if json.User != "" && json.Password != "" {
			db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gotuto sslmode=disable password=vili")
			if err != nil {
				panic("Falha ao conectar ao banco de dados")
			}
			defer db.Close()

			var user = User{Username: json.User, Userpassword: json.Password, Userpasswordb64: b64.StdEncoding.EncodeToString([]byte(json.Password))}

			//Verifica se já existe o valor caso sim = true
			var response = db.NewRecord(user)

			db.Create(&user)

			if response {
				c.JSON(http.StatusConflict, gin.H{"status": "Registro já existente"})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": "Registrado com sucesso"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
