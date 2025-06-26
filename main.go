package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Cotacao representa a estrutura de dados para uma cotação de moeda.
type Cotacao struct {
	Moeda string  `json:"moeda"`
	Valor float64 `json:"valor"`
	Data  string  `json:"data"`
}

func main() {
	http.HandleFunc("/cotacao", handleCotacao) // Define o manipulador para a rota /cotacao

	fmt.Println("Servidor Go de Cotações iniciado na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Inicia o servidor na porta 8080
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	// Cria dados de cotação simulados
	cotacoes := []Cotacao{
		{Moeda: "USD", Valor: 5.25, Data: time.Now().Format("02-01-2006 15:04:05")},
		{Moeda: "EUR", Valor: 5.70, Data: time.Now().Format("02-01-2006 15:04:05")},
		{Moeda: "BRL", Valor: 1.00, Data: time.Now().Format("02-01-2006 15:04:05")},
	}

	// Define o cabeçalho Content-Type para JSON
	w.Header().Set("Content-Type", "application/json")

	// Codifica os dados para JSON e envia a resposta
	json.NewEncoder(w).Encode(cotacoes)
}
