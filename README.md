# Desenho Arquitetural - Padrão MVC

## Descrição dos Componentes

### `/src/controllers/product_controller.go`:
Responsável por lidar com as requisições HTTP, processando a entrada (ex: dados de formulários) e preparando a resposta. O controlador comunica-se com o serviço para realizar a lógica de negócios e retornar a resposta ao cliente. No caso dos produtos, ele contém funções como `CreateProduct`, `GetProductByID`, `GetAllProducts`, etc.

### `/src/database/db.go`:
Responsável pela configuração da conexão com o banco de dados. Este arquivo contém as funções de configuração para o banco de dados SQLite, além da inicialização das tabelas e migrações de estrutura de dados. Em um ambiente de teste, ele também lida com a criação de um banco de dados temporário.

### `/db/migrations/0001_create_products_table.sql`:
Arquivo SQL usado para criar a tabela `products` no banco de dados. Ele define a estrutura da tabela com as colunas necessárias, como `ID`, `Name`, `Price`, etc. Migrações devem ser editadas manualmente.

### `/src/models/product.go`:
Define a estrutura do produto, ou seja, como os dados do produto são representados na aplicação. Este arquivo é onde as entidades são modeladas, e em Go, ele usa o ORM (GORM) para mapear os campos da struct para as colunas no banco de dados.

### `/src/repositories/product_repository.go`:
Contém funções que interagem diretamente com o banco de dados para recuperar, salvar, atualizar e excluir dados. O repositório faz a ponte entre a lógica de negócios (no serviço) e a persistência dos dados no banco de dados. Ele usa o GORM para realizar essas operações.

### `/src/routes/routes.go`:
Define as rotas HTTP e mapeia as requisições para os controladores apropriados. Aqui, as rotas são associadas aos métodos correspondentes no controlador de produto (por exemplo, `GET /products/{id}` chama o método `GetProductByID`).

### `/src/services/product_service.go`:
Contém a lógica de negócios do produto, como a criação de produtos, busca por produtos no banco, validações de dados e regras de negócio. O serviço usa o repositório para manipular os dados e retorna as informações ou erros adequados para o controlador.

### `Arquivos */*_test.go`:
Contém os testes unitários da aplicação. Controladores e services apenas estão no escopo deste projeto. Os testes são realizados para garantir que os contratos e as respostas estão de acordo com o esperado.

### `go.mod`:
Arquivo de configuração do Go que gerencia as dependências do projeto. Ele garante que as bibliotecas corretas sejam baixadas e que a versão correta do Go seja utilizada.

### `main.go`:
Arquivo principal da aplicação onde o servidor é iniciado, as dependências são configuradas e as rotas são registradas. Ele é responsável por iniciar o servidor HTTP e chamar as funções de configuração e inicialização do projeto.

## Explicação Geral do Fluxo da Aplicação

1. **Usuário/Cliente** faz uma requisição HTTP (por exemplo, `GET /products/1`) para a aplicação.
2. A rota configurada em `routes.go` mapeia a requisição para o controlador correspondente em `product_controller.go`.
3. O **controlador** processa a requisição, e, se necessário, invoca o **serviço** (`product_service.go`) para realizar a lógica de negócios.
4. O **serviço** acessa o **repositório** (`product_repository.go`) para interagir com o banco de dados (SQLite) e recuperar ou manipular os dados.
5. O resultado é enviado de volta ao **controlador**, que então retorna a resposta para o usuário (cliente).

## Conclusão
Essa arquitetura segue o padrão **MVC** (Model-View-Controller), onde:
- **Model** é representado pelas **entities** (como `models/product.go`),
- **View** é a interação com o usuário, que é gerenciada pela **API** (controlador),
- **Controller** cuida da comunicação entre a **View** (a interface da API) e o **Model** (dados persistidos no banco).

Esse padrão facilita a organização e escalabilidade do código, permitindo uma clara separação das responsabilidades, o que é essencial para a manutenção e a evolução de sistemas grandes e complexos.

---

## Estrutura de arquivos do Projeto

```plaintext
./bootCampArquitetura
├── /.vscode                                    # Pasta com as configurações de depuração do projeto.
├── /produtos-api                               # Produtos-api, representado em azul no diagrama C4
│   ├── /db/migrations                          # Cada nova migração deve ser armazenada aqui. O sistema as executa automaticamewnte
│   │   └── XXXXXXXX_create_products_table.sql  # Arquivo de migração para a criação da tabela 'products'
│   ├── /src
│   │   ├── /controllers
│   │   │   └── product_controller.go           # Controlador responsável por receber as requisições HTTP e delegar para os serviços apropriados
│   │   │   └── product_controller_test.go      # Testes unitários para a controladora de produtos
│   │   ├── /database
│   │   │   └── db.go                           # Configuração do banco de dados SQLite e funções auxiliares
│   │   ├── /models
│   │   │   └── product.go                      # Model que representa a estrutura de dados do Produto
│   │   ├── /repositories
│   │   │   └── product_repository.go           # Repositório responsável pela interação com o banco de dados
│   │   ├── /routes
│   │   │   └── routes.go                       # Definição das rotas HTTP e seu mapeamento para os controladores
│   │   ├── /services
│   │   │   └── product_service.go              # Serviço que contém a lógica de negócios para manipulação dos produtos
│   │   │   └── product_service_test.go         # Testes unitários para o serviço de produtos
├── go.mod                                  # Gerenciamento de dependências do Go
├── main.go                                 # Arquivo principal para iniciar o servidor e definir a configuração do projeto
├── product.sqlite                          # Base de dados da aplicação
└── c4_israelborba.drawio                    # Arquivo contendo o diagrama C4 da aplicação

```

## Comandos úteis

### Gerar migrations:
```sh
$ migrate create -ext sql -dir db/migrations NOME_DA_MIGRATION
# Obs: Os arquivos de migração são criados vazios e precisam ser preenchidos com o SQL desejado.
```
### Rodar migrations:
```sh 
migrate -path db/migrations -database "sqlite3://./products.sqlite" up
# Obs: no final da linha, observe a opção up. Ela é utilizada para subir as alterações. Caso seja necessário desfazer, basta substituir up por down.
```
#### Importante a aplicação, através da linha abaixo, executa as migrations automaticamente:
```golang
err = db.AutoMigrate(&models.Product{})
```

### Rodar a aplicação:
```sh
go run main.go
```

### Rodar testes unitários:
```sh
go test -v ./src/...
# Ou, rodar testes com o coverage detalhado em HTML:
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

### Swagger
Com o projeto em execução, você consegue acessar a documentação da API aqui: http://localhost:8080/swagger/index.html
