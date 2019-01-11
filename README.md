## How to run search_server and elastic ?

Run in the project root folder

```
 make build-linux && make build-docker && docker-compose up -d
```

## How to run test ?

Run in the project root folder

```
 make test
```

## How to test search_server API ?

You can test API in your browser:

```
http://localhost:8080/v1/products
http://localhost:8080/v1/products?q=tshirt
http://localhost:8080/v1/products?q=sneakers
http://localhost:8080/v1/products?q=sneakers&sort=price:desc&page=1&per_page=2
http://localhost:8080/v1/products?q=sneakers&page=1&per_page=5&sort=price:asc&filter=brand:Nike
http://localhost:8080/v1/products?q=tshirt&sort=price:desc&page=1&per_page=3
http://localhost:8080/v1/products?q=tshirt&sort=price:desc&page=1&per_page=3&filter=brand:Asics
```

