Desenho Arquitetural - Padrão MVC
Descrição do Fluxo:

Controller: Recebe as requisições HTTP, valida os dados e direciona a lógica para os serviços adequados.
Service: Contém a lógica de negócios (e.g., validações complexas, cálculos, etc.).
Model: Representa os dados da aplicação, mapeados para o banco de dados (ORM com gorm).
Banco de Dados: Armazena os registros de produtos usando SQLite.

--

Estrutura Inicial do Projeto

src/
├── controllers/
│   └── product_controller.go
├── models/
│   └── product.go
├── services/
│   └── product_service.go
├── database/
│   └── db.go
├── routes/
│   └── routes.go
├── utils/
│   └── response.go
└── main.go

--

## Comandos úteis:

- Gerar migrations:
``` sh
$ migrate create -ext sql -dir db/migrations NOME_DA_MIGRATION
```
Obs: Os arquivos de migração são criados vázios e precisam ser preenchidos com o SQL desejado

- Rodar migrations:
``` sh
migrate -path db/migrations -database "sqlite3://./products.sqlite" up
```
Obs: no final da linha, observe a opção 'up'. Ela é utilizada para subir as alterações.
Caso seja necessário desfazer, basta substituir 'up' por 'down'

### Vale a pena notar que a aplicação, através da linha abaixo, executa as migrations automaticamente
``` golang
err = db.AutoMigrate(&models.Product{})
```

- Rodar a aplicação:
``` sh
go run main.go
```
- Rodar testes unitários:
``` sh
go test -v ./src/...
```
- Ou, rodar testes com o coverage detalhado em HTML
``` sh
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

