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
