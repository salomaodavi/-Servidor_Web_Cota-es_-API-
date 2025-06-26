# 📈 Servidor Web de Cotações Simplificado com Go 🚀

---

## Conecte-se comigo! 🤝

Ficarei feliz em conectar e trocar ideias! Você pode me encontrar no LinkedIn:

[![Meu LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/salomao-davi)

---

Seja bem-vindo(a) ao meu projeto de Servidor Web de Cotações Simplificado! 🎉 Este é um exemplo prático de como construir uma API RESTful leve e eficiente usando a linguagem Go (Golang). O objetivo é demonstrar a facilidade de criação de serviços de backend com Go, fornecendo cotações de moedas em formato JSON.

Este projeto é ideal para quem deseja explorar o ecossistema web do Go, entender como funcionam as requisições HTTP e ver uma API básica em ação. É um excelente ponto de partida para desenvolver serviços mais complexos! 🌐

## 🌟 O que este projeto faz?

Este servidor é um "fornecedor" de dados de cotação. Ele:
* Inicia um servidor HTTP: Ele "escuta" em uma porta específica (padrão: 8080) por requisições de outros aplicativos ou do seu navegador.

* Oferece um endpoint ``/cotacao``: Ao acessar essa rota, ele responde com um conjunto de dados simulados de cotações de moedas (Dólar, Euro, Real).

* Retorna dados em JSON: A resposta é formatada em JSON, um padrão universal para troca de dados entre sistemas.

# ⚙️ Como rodar o projeto
Rodar este servidor Go é super simples e rápido!

### 1. Pré-requisitos:
* Certifique-se de ter o Go (Golang) 1.16 ou superior instalado na sua máquina. Você pode baixar em https://go.dev/dl/

### 2. Clone o repositório (ou copie o código para um arquivo ``.go``):
````
git clone [LINK_DO_SEU_REPOSITORIO]
cd [NOME_DA_PASTA_DO_PROJETO]
````

### 3. Execute o script:
````
go run main.go
````
(Assumindo que você salvou o código como ``main.go``)

### 4. Acesse a API!
* Abra seu navegador ou use uma ferramenta como curl para acessar a rota:
````
http://localhost:8080/cotacao
````
(Assumindo que você salvou o código como ``main.go``)

* Você verá uma resposta JSON similar a esta:
# Json
````
[
    {
        "moeda": "USD",
        "valor": 5.25,
        "data": "26-06-2025 10:30:00"
    },
    {
        "moeda": "EUR",
        "valor": 5.7,
        "data": "26-06-2025 10:30:00"
    },
    {
        "moeda": "BRL",
        "valor": 1.00,
        "data": "26-06-2025 10:30:00"
    }
]
````
---
# Imagens:

![Captura de tela 2025-06-26 181606](https://github.com/user-attachments/assets/2b888efc-d1e5-4cc1-95fe-fecae1e68218)

![Captura de tela 2025-06-26 181620](https://github.com/user-attachments/assets/6605e25d-5a73-4f72-94fb-0ec6c29ec824)

![Captura de tela 2025-06-26 181713](https://github.com/user-attachments/assets/71d45769-5ee6-421b-9bc9-46c33787831e)

---

# 💻 O Código
Aqui está o código completo do servidor web de cotações:
````
package main

import (
	"encoding/json" // Pacote para codificar e decodificar JSON
	"fmt"           // Pacote para formatação de I/O
	"log"           // Pacote para logging
	"net/http"      // Pacote para funcionalidade de cliente e servidor HTTP
	"time"          // Pacote para funcionalidades de tempo
)

// Cotacao representa a estrutura de dados para uma cotação de moeda.
// As tags `json:"moeda"` etc. são usadas para mapear os campos da struct para chaves JSON.
type Cotacao struct {
	Moeda string  `json:"moeda"` // Nome da moeda (ex: USD, EUR)
	Valor float64 `json:"valor"` // Valor da cotação
	Data  string  `json:"data"`  // Data e hora da cotação
}

func main() {
	// http.HandleFunc associa uma rota (/cotacao) a uma função manipuladora (handleCotacao).
	http.HandleFunc("/cotacao", handleCotacao)

	fmt.Println("Servidor Go de Cotações iniciado na porta :8080 🚀")
	// http.ListenAndServe inicia o servidor HTTP.
	// Ele espera um endereço (neste caso, ":8080" significa todas as interfaces na porta 8080)
	// e um handler (nil significa usar o DefaultServeMux, onde as rotas foram registradas com HandleFunc).
	log.Fatal(http.ListenAndServe(":8080", nil)) // log.Fatal encerra o programa se houver um erro
}

// handleCotacao é a função que lida com as requisições HTTP para a rota /cotacao.
// w (http.ResponseWriter) é onde escrevemos a resposta HTTP.
// r (*http.Request) contém os detalhes da requisição HTTP de entrada.
func handleCotacao(w http.ResponseWriter, r *http.Request) {
	// Cria um slice de Cotacao com dados simulados.
	// time.Now().Format(...) formata a data e hora atual.
	cotacoes := []Cotacao{
		{Moeda: "USD", Valor: 5.25, Data: time.Now().Format("02-01-2006 15:04:05")},
		{Moeda: "EUR", Valor: 5.70, Data: time.Now().Format("02-01-2006 15:04:05")},
		{Moeda: "BRL", Valor: 1.00, Data: time.Now().Format("02-01-2006 15:04:05")},
	}

	// Define o cabeçalho 'Content-Type' da resposta como 'application/json'.
	// Isso informa ao cliente que o corpo da resposta é um JSON.
	w.Header().Set("Content-Type", "application/json")

	// json.NewEncoder(w).Encode(cotacoes) escreve a representação JSON do slice 'cotacoes'
	// diretamente no http.ResponseWriter (w).
	err := json.NewEncoder(w).Encode(cotacoes)
	if err != nil {
		// Se houver um erro ao codificar o JSON, loga o erro e define um status HTTP 500 (Internal Server Error).
		log.Printf("Erro ao codificar JSON: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
	}
}
````

# 🚀 Próximos passos e possíveis melhorias
Este projeto é uma base sólida! Aqui estão algumas ideias para torná-lo ainda mais robusto e interessante:

* Consumo de API Real: Em vez de dados simulados, integrar com uma API de cotações de moedas real (ex: Banco Central do Brasil, ExchangeRate-API, Open Exchange Rates). 🏦

* Tratamento de Erros: Implementar tratamento de erros mais sofisticado e respostas HTTP apropriadas (ex: 404 Not Found, 400 Bad Request).

* Parâmetros de Requisição: Permitir que o cliente especifique a moeda desejada via parâmetros de URL (ex: ``/cotacao?moeda=USD``).

* Dockerização: Empacotar a aplicação em um container Docker para facilitar a implantação. 🐳

* Testes Unitários: Escrever testes para garantir que o endpoint da API funcione conforme o esperado. ✅

* Logging: Implementar um sistema de logging mais detalhado para monitorar as requisições e a saúde do servidor.

## 📄 Licença

MIT License

Copyright (c) [2025] [Salomão Davi Cardoso Barros]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
---

[![Status do Projeto](https://img.shields.io/badge/Status-Em%20Desenvolvimento-yellow.svg)](https://github.com/salomao-davi/[nome-do-seu-repositorio])
[![Licença MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

---









