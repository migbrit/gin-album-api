# Gin API Album

O projeto **gin-api-album** é uma API simples desenvolvida em Go (Golang) utilizando o framework web Gin. Este projeto é utilizado na documentação oficial da linguagem Go como parte do material de aprendizado.

## Objetivo

O objetivo desta API é demonstrar como criar uma aplicação web básica em Go, abordando conceitos como roteamento, manipulação de solicitações HTTP, e interação com dados em formato JSON. A API gerencia uma lista de álbuns de música, permitindo operações como obtenção da lista de álbuns e adição de novos álbuns.

## Estrutura do Projeto

O projeto está estruturado da seguinte maneira:

- **album.go:** Contém a definição da estrutura `album` e as funções relacionadas a manipulação dos álbuns.

- **main.go:** Inicia a aplicação, configura as rotas da API e define os manipuladores para cada rota.

## Dependências

- **Gin:** O framework web utilizado para simplificar o desenvolvimento da API. Para instalar, utilize o comando:
  ```bash
  go get -u github.com/gin-gonic/gin
