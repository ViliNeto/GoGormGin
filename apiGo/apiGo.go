package apigo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../dto"
)

//RootPage Função usando net/http do Go
func RootPage(w http.ResponseWriter, r *http.Request) {
	//GET simples
	//fmt.Fprintf(w, "Bem vindo!")

	//GET com URL Path
	//fmt.Fprintf(w, "Olá, eu amo %s!", r.URL.Path[1:])

	//POST
	decoder := json.NewDecoder(r.Body)

	var data dto.Login
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Fala %s!", data)

}
