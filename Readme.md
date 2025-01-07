# Desafio Taghos

Este projeto foi desenvolvido como parte do desafio técnico da Taghos. Ele consiste em uma API para gerenciar livros, construída com **Golang**, utilizando **Fiber** como framework web, **MongoDB** como banco de dados, e **Swagger** para documentação das rotas.

## 🚀  Passo a Passo para Desenvolvimento

<details>

### 1️⃣ Configuração do Ambiente

1. Vá até o arquivo `.envExample`, duplique-o e renomeie para `.env`.
2. Configure as variáveis de ambiente de acordo com suas necessidades.
   - Certifique-se de definir corretamente o **MongoDB** que será utilizado (interno ou externo).
   - Caso utilize o **MongoDB interno**, abra as portas necessárias no arquivo `docker-compose.yml`.

---

### 2️⃣ Subir o Ambiente com Docker
Para subir o ambiente com Docker, siga os passos abaixo:

1. **Buildar os containers**:
   - Utilize o comando abaixo para buildar os containers:
   ```bash
   docker-compose build
   ```
2. **Subir os containers:**:
   - Após o build ser concluído, utilize o comando abaixo para subir os containers:
   ```bash
   docker-compose up
   ```
   Isso fará com que:
   - Um container com o MongoDB seja iniciado e configurado.
   - Um container com a aplicação dockerizada seja iniciado, permitindo que você use a aplicação sem precisar instalar ferramentas adicionais como air ou swagger na sua máquina.

---

### 3️⃣ Acessar o Container e Desenvolver
Para entrar no container da aplicação, use o comando:
```bash
docker-compose exec app bash
```
Dentro do container, você poderá:

- Modificar ou implementar novas funcionalidades na aplicação.
- Atualizar a documentação do Swagger.
- Rodar a aplicação com live reloading usando o `air`.

**Comandos úteis disponíveis no Makefile:**
- `make dev`: Roda a aplicação com live reloading (via `air`).
- `make swagger`: Atualiza a documentação do Swagger.
Outros comandos podem ser encontrados no arquivo `Makefile`.

---

### 4️⃣ Testando as Rotas
1. Inicie a aplicação com:
```bash
make dev
```
2. Acesse o endpoint `/indexing/book` para criar o banco de dados e os índices da coleção.
3. Navegue até o Swagger:
   - URL: `http://localhost:<porta-configurada>/swagger`
   - Aqui você encontrará todos os endpoints explicados e documentados, incluindo os exemplos de bodies das requisições.

---

### 5️⃣ Alternativa para Visualizar o Swagger
Se preferir, você pode usar o arquivo JSON do Swagger diretamente:
- Localize o arquivo `docs/swagger.json` no projeto.
- Acesse o site https://editor.swagger.io/.
- Importe o arquivo:
  - Menu: **File > Import File.**
  - Isso permitirá visualizar e testar as rotas diretamente no editor.

---

### 🔗 Recursos Adicionais
- **Live Reloading**: Use o comando `make dev` para desenvolvimento em tempo real.
- **Swagger**: Documentação gerada automaticamente para facilitar a integração e o teste das rotas.
- **MongoDB**: Banco de dados `NoSQL` para armazenar os dados do projeto.

</details>

---

### 📚 Tecnologias Utilizadas
- **Golang**: Linguagem principal para a aplicação.
- **Fiber**: Framework web utilizado.
- **MongoDB**: Banco de dados.
- **Swagger**: Para documentação de API.
- **Docker**: Para containerização.
- **Air**: Ferramenta para live reloading no desenvolvimento.

---

### 🗃️ Por que MongoDB e não PostgreSQL?

Optei pelo **MongoDB** devido à minha experiência com o banco, que inclui diversos projetos e um treinamento especializado que recebi na época em que precisei desenvolver um projeto importante na empresa em que trabalho. Esse treinamento focou em indexação e otimização, e o projeto resultante foi considerado um **case de sucesso** pela equipe do **MongoDB** na América Latina, sendo usado como exemplo para outros desenvolvedores da região.

Além disso, por ser um CRUD simples de criação de livros, o MongoDB se mostrou mais adequado, já que não há necessidade de relações complexas, tornando a implementação mais rápida e eficiente.

---

### 📋 Decisões Tomadas Durante o Desafio Técnico

Aproveitei a oportunidade deste desafio técnico para praticar a aplicação do **Swagger** e a **Arquitetura Hexagonal** no projeto. O maior trabalho foi justamente estruturar todos os arquivos utilizando essa arquitetura, o que me proporcionou uma experiência prática valiosa. Após a estruturação, a conexão com o **MongoDB** foi realizada de forma tranquila, considerando a familiaridade que tenho com o banco.

O primeiro endpoint implementado foi o **Create Book**, que serviu como base para os outros endpoints do CRUD. Depois de concluir o **Create**, os outros endpoints, como **Update**, **Get All** e **Delete**, foram mais fáceis e rápidos de desenvolver, devido à simplicidade da lógica e à estrutura já estabelecida.

A integração com o **Swagger** também foi uma ótima oportunidade para documentar a API de maneira eficiente, o que facilita a comunicação e o entendimento da estrutura do sistema.

No geral, o desafio foi uma excelente oportunidade de aplicar conceitos que já vinha estudando e, ao mesmo tempo, entregar uma solução prática e eficiente para o CRUD de livros.

**Estrutura do Projeto**

A estrutura do projeto foi organizada seguindo a Arquitetura Hexagonal, com os componentes distribuídos nas pastas principais. A organização dos arquivos ficou da seguinte forma:
```bash
├── bin
├── cmd
│   └── main.go
├── docs
├── infra
│   ├── config
│   └── database
│       └── mongodb
└── internal
    ├── adapter
    │   ├── handler
    │   │   └── http
    │   │       ├── middleware
    │   │       ├── responses
    │   │       └── routes
    │   └── repository
    │       └── postgres
    │           └── migrations
    ├── core
    │   ├── domain
    │   ├── port
    │   └── service
    └── framework
        ├── logs
        └── util
```
