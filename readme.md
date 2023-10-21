# Imersão Full Cycle | Desafio 01 - Criando app Golang com gRPC

## 1️⃣ 🔖 Descrição do desafio 01:

Dado o seguinte protofile gRPC:
https://gist.github.com/argentinaluiz/8a2e1a7a32da3d1b107d88ec9f1007b2

Crie um servidor gRPC em Golang que realize as duas operações:

Criar produtos
Consultar produtos

A aplicação deve usar o SQLite como banco de dados e o GORM como ORM. Use o modo AutoMigrate para gerar as tabelas automaticamente.

Você deve disponibilizar um o script main.go que a rodar go run main.go deve levantar o servidor e deixar a porta 50051 acessível via localhost.

## 2️⃣ 🔖 Descrição do desafio 02:
Neste desafio, você deve pegar a aplicação do desafio 1 e integra-la com Docker Compose, ou seja, ao rodar o comando "docker compose up" já deve subir automaticamente o servidor gRPC na porta 50051.

A segunda etapa do desafio é criar uma aplicação Nest.js que seja o gRPC client do Golang, portanto teremos uma API REST na porta 3000 que terá 2 endpoints:

```POST /products``` - que ao ser acessado deve fazer uma chamada rpc para o CreateProduct do Golang e devolver como resultado os dados do produto criado

```GET /products``` - Ao ser acessado deve fazer uma chamada rpc para o FindProducts e retornar a lista de produtos do Golang

Disponibilize a aplicação Nest.js no mesmo docker-compose.yaml.

Portanto, ao rodar o docker compose up deve subir a aplicação Golang e Nest.js automaticamente.

## 🖱️ Como executar a aplicação:

1. Clone o repositório
2. Na pasta `golang` crie o arquivo .env baseado no arquivo de exemplo
3. Tendo o Docker instalado em sua máquina e estando na pasta raiz do projeto, execute o comando no terminal: `docker compose up`
4. Aguarde as mensagens `gRPC server has been started on port 50051` e `Nest application successfully started` serem exibidas para indicar sucesso ao subir a aplicação
5. Na pasta nestjs, execute as chamadas para as rotas `GET /products` e `POST /products` utilizando o arquivo `api.http`

## 🧪 TDD: Executando os testes

A aplicação em golang foi desenvolvida com base na técnica TDD. Você pode executar os testes no terminal estando na pasta `golang` com o comando `go test ./...`
