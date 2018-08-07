package apiGin

import (
	b64 "encoding/base64"
	"github.com/jinzhu/gorm"
	"net/http"
	"github.com/gin-gonic/gin"
	"../dto"
)

func StartRouter() *gin.Engine{
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
	return router
}

func Getting(c *gin.Context) {

	//Simples
	c.String(200, "Recebido")

	//Caso venha algo no Path
	user := c.Param("User")
	if user == "" {
		user = c.Query("User")
	}
	c.String(http.StatusOK, "\nOlá %s", user)
}

func Posting(c *gin.Context) {
	var json dto.Login
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

func Register(c *gin.Context) {
	var json dto.Login
	if err := c.ShouldBindJSON(&json); err == nil {
		if json.User != "" && json.Password != "" {
			db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gotuto sslmode=disable password=vili")
			if err != nil {
				panic("Falha ao conectar ao banco de dados")
			}
			defer db.Close()

			var user = dto.User{Username: json.User, Userpassword: json.Password, Userpasswordb64: b64.StdEncoding.EncodeToString([]byte(json.Password))}

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
