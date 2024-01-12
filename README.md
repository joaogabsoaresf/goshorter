# Encurtador de URL com Go, Gin e MongoDB

### Obs: No momento o dominio está sem o certificado de segurança, isso será resolvido nos próximos dias.

A aplicação está online e pode ser usada perfeitamente. O link está anexado ao projeto, porém a página principal ainda se encontra vazia.

Para receber o seu link encurtado basta enviar um POST seguindo o curl abaixo:

```bash
curl --location 'http://makeshorter.biz/api/v1/url' \
--header 'Content-Type: application/json' \
--data '{
    "original_path":"https://seu-site.com/"
}'
```

Para visualizar as informações sobre o seu link encurtado, pode enviar um GET contendo o ID como parêmetro, seguindo o curl abaixo:

```bash
curl --location 'http://makeshorter.biz/api/v1/url?shorter_id=VyQMs'
```

## Pré-requisitos

Certifique-se de ter o Go instalado em sua máquina. Caso ainda não tenha, você pode baixá-lo [aqui](https://golang.org/dl/).

É preciso baixar o MongoDB também, a documentação está [aqui](https://www.mongodb.com/docs/v3.0/tutorial/install-mongodb-on-windows/).

Além disso, é necessário instalar as dependências do projeto. Para isso, utilize o seguinte comando:

```bash
go mod download
```

# Executando o Projeto

Você pode executar o projeto de duas maneiras diferentes:

1. Usando o comando go
Inicie o projeto diretamento pelo comando do GO:

```bash
go run main.go
```

O servidor será iniciado e estará acessível em [localhost:8080](http://localhost:8080)

# Atualizações em Andamento
Ainda não tive tempo de subir todas as atualizações, devo fazer em breve as seguintes adições:

- Documentação com Swagger;
- Adição de testes;
- Fluxo com Docker;

## Contribuindo
Se você quiser contribuir para este projeto, sinta-se à vontade para abrir uma issue ou enviar uma solicitação pull. Espero que este projeto evolua com a ajuda da comunidade!
