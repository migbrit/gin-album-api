Gin API Album
O projeto gin-api-album é uma API simples desenvolvida em Go (Golang) utilizando o framework web Gin. Este projeto é utilizado na documentação oficial da linguagem Go como parte do material de aprendizado.

Objetivo
O objetivo desta API é demonstrar como criar uma aplicação web básica em Go, abordando conceitos como roteamento, manipulação de solicitações HTTP, e interação com dados em formato JSON. A API gerencia uma lista de álbuns de música, permitindo operações como obtenção da lista de álbuns e adição de novos álbuns.

Estrutura do Projeto
O projeto está estruturado da seguinte maneira:

album.go: Contém a definição da estrutura album e as funções relacionadas a manipulação dos álbuns.

main.go: Inicia a aplicação, configura as rotas da API e define os manipuladores para cada rota.

Dependências
Gin: O framework web utilizado para simplificar o desenvolvimento da API. Para instalar, utilize o comando:
bash
Copy code
go get -u github.com/gin-gonic/gin
Como Executar
Clone o repositório:

bash
Copy code
git clone https://github.com/seu-usuario/gin-api-album.git
cd gin-api-album
Execute a aplicação:

bash
Copy code
go run main.go
A API estará acessível em http://localhost:8080.

Rotas da API
1. Obter Lista de Álbuns
Endpoint: GET /albums

Retorna a lista de todos os álbuns em formato JSON.

2. Adicionar Novo Álbum
Endpoint: POST /albums

Adiciona um novo álbum com base nos dados JSON fornecidos no corpo da requisição. Exemplo de corpo de requisição:

json
Copy code
{
  "id": "4",
  "title": "Kind of Blue",
  "artist": "Miles Davis",
  "price": 29.99
}
A resposta conterá os detalhes do álbum recém-adicionado.

Contribuições
Contribuições são bem-vindas! Sinta-se à vontade para abrir problemas (issues) ou enviar pull requests para melhorar este projeto.

Licença
Este projeto está licenciado sob a Licença MIT - consulte o arquivo LICENSE para obter detalhes.
