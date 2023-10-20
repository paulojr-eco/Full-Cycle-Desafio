# Imersão Full Cycle | Desafio 01 - Criando app Golang com gRPC

## 🔖 Descrição do desafio:

Dado o seguinte protofile gRPC:
https://gist.github.com/argentinaluiz/8a2e1a7a32da3d1b107d88ec9f1007b2

Crie um servidor gRPC em Golang que realize as duas operações:

Criar produtos
Consultar produtos

A aplicação deve usar o SQLite como banco de dados e o GORM como ORM. Use o modo AutoMigrate para gerar as tabelas automaticamente.

Você deve disponibilizar um o script main.go que a rodar go run main.go deve levantar o servidor e deixar a porta 50051 acessível via localhost.

## 🖱️ Como executar a aplicação:

1. Clone o repositório
2. Crie o arquivo .env baseado no arquivo de exemplo
3. Tendo o Docker instalado em sua máquina e estando na pasta raiz do projeto, execute o comando `docker-compose up -d`
4. Execute o comando `docker exec -it full-cycle-desafio-app-1 bash` para acessar o bash do app
5. Em sequência, rode a linha de comando para iniciar o servidor gRPC: `go run main.go`
6. A mensagem `gRPC server has been started on port 50051` deve ser exibida para indicar sucesso ao subir o servidor
7. Abra um novo terminal e acesse novamente o bash do app (`docker exec -it full-cycle-desafio-app-1 bash`)
8. Rode `evans -r repl` para acessar o client Evans
9. Primeiramente, acesse o package através do comando: `package github.com.paulojr.desafio`
10. Em seguida acesse o service: `service ProductService`
11. Finalmente realize a chamada para o serviço de criação de produto `(call CreateProduct)` e preencha todos os campos ou para o serviço de listagem de produtos `(call FindProducts)`

## 🧪 TDD: Executando os testes

Essa aplicação foi desenvolvida com base na técnica TDD. Você pode executar os testes no terminal do próprio clone do repositório com o comando `go test ./...`
