# Go Opportunities

[![Build Status](https://github.com/valdir-alves3000/go-opportunities/workflows/Go%20Opportunities%20Actions/badge.svg)](https://github.com/valdir-alves3000/go-opportunities/actions)


## Descrição
Go Opportunities é uma aplicação desenvolvida em GoLang para gerenciamento de vagas de emprego. O projeto permite criar, buscar, atualizar e deletar vagas de trabalho, utilizando o GORM como ORM para interação com o banco de dados.

## Tecnologias Utilizadas
- GoLang
- GORM
- SQLite
- Testes com Testify

## Instalação
1. Clone o repositório:
   ```sh
   git clone https://github.com/valdir-alves3000/go-opportunities.git
   cd go-opportunities
   ```
2. Instale as dependências:
   ```sh
   go mod tidy
   ```

## Testes
Os testes unitários estão implementados na pasta `test/unit`. Para rodar os testes, utilize:
```sh
go test ./unit/...
```