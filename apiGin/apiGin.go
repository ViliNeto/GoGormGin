package apigin

import (
	b64 "encoding/base64"
	"net/http"
	"time"

	"../connection"
	"../dto"
	"github.com/gin-gonic/gin"
)

//StartRouter Função usada para se Startar o router do Gin
func StartRouter() *gin.Engine {
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

//Getting Função usada para se receber uma chamada get do Router do Gin
func Getting(c *gin.Context) {

	//Simples
	c.String(http.StatusOK, "Recebido")

	//Caso venha algo no Path
	user := c.Param("User")
	if user == "" {
		user = c.Query("User")
	}
	if user != "" {
		c.String(http.StatusOK, "\nOlá %s", user)
	} else {
		c.String(http.StatusOK, "\nOlá usuário desconhecido")

	}

}

//Posting Função usada para se receber uma chamada post do Router do Gin
func Posting(c *gin.Context) {
	var json dto.Login
	if err := c.ShouldBindJSON(&json); err == nil {
		if json.User == "vili" && json.Pass == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "Você está logado"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Acesso não autorizado"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//Registrar Função usada para se receber uma chamada post do Router do Gin e gravar dentro de um banco de dados a requisição
func Registrar(c *gin.Context) {
	var json dto.Login
	if err := c.ShouldBindJSON(&json); err == nil {
		if json.User != "" && json.Pass != "" {

			//Abre conexão com o Banco
			db, _ := connection.AbrirConexao()
			defer db.Close()

			//Mapeia resposta da requisição para Struct User
			var user = dto.User{Username: json.User, Userpass: json.Pass, Userpassb64: b64.StdEncoding.EncodeToString([]byte(json.Pass))}

			//Inicializa estrutura User que voltará de um select no Database
			var returnUser dto.User

			//Verifica se já existe
			db.Where("username = ?", json.User).First(&returnUser)

			//Caso valor já exista não se cria o registro
			if returnUser.Username != "" {
				c.JSON(http.StatusConflict, gin.H{"status": "Registro já existente"})
			} else {
				// db.NewRecord(user)
				db.Create(&user)
				c.JSON(http.StatusOK, gin.H{"status": "Registrado com sucesso"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//AtualizaSenha é um metodo que é chamado para alterar a senha do usuário
func AtualizaSenha(c *gin.Context) {
	var resp dto.Login
	if err := c.ShouldBindJSON(&resp); err == nil {
		if resp.User != "" && resp.Pass != "" {

			//Abre conexão
			db, _ := connection.AbrirConexao()
			defer db.Close()

			//Select do banco de dados será gravado nessa variável
			var returnUser dto.User

			//Verifica se já existe
			db.Where("username = ?", resp.User).First(&returnUser)

			//Caso encontre o usuário
			if returnUser.Username != "" {
				//checa se a senha dele confere
				if b64.StdEncoding.EncodeToString([]byte(resp.Pass)) == returnUser.Userpassb64 {
					//Caso nova senha seja vazia
					if resp.NewPass == "" {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Nova senha inválida"})
					} else {
						go atualizaSenhaDatabase(resp)
						c.JSON(http.StatusAccepted, gin.H{"status": "Atualização em andamento"})
					}
				} else {
					//Caso a senha não seja igual a original
					c.JSON(http.StatusForbidden, gin.H{"status": "Senha incorreta"})
				}
			} else {
				//Caso o usuário não seja encontrado
				c.JSON(http.StatusForbidden, gin.H{"status": "Usuário não encontrado"})
			}
		} else {
			//Caso usuário ou senha sejam vazios
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	} else {
		//Caso não consiga fazer o unmarshal do json
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func atualizaSenhaDatabase(login dto.Login) {
	//Simulando uma execução em fila (não processada no momento)
	time.Sleep(10 * time.Second)
	db, _ := connection.AbrirConexao()
	defer db.Close()
	//Update a senha do usuário
	err := db.Exec("UPDATE users SET userpasswordb64 = ?, userpassword = ?  WHERE username = ?", b64.StdEncoding.EncodeToString([]byte(login.NewPass)), login.NewPass, login.User)
	println(err)
}
