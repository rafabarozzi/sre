package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Define a rota padrão "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Envia a mensagem para o browser
		fmt.Fprint(w, "Hello From Go - k8s!!")
	})
	//Imprime a mensagem no console
	fmt.Printf("Server is running on port 3000...")
	//Inicia o service na porta 3000 e verifica se houve erro
	log.Fatal(http.ListenAndServe(":3000", nil))
}
