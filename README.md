# 2vehiculo

## Project setup
```
sudo docker-compose up
```

## Routes

### Internal (Key Auth)
Use header AdminToken
```
GET http://localhost:8080/api/internal/orders
POST http://localhost:8080/api/internal/orders
GET http://localhost:8080/api/internal/orders/1
PUT http://localhost:8080/api/internal/orders/1
DEL http://localhost:8080/api/internal/orders/1
```

### USER (JWT Auth)
Use LOGIN endpoint 
```
POST http://localhost:8080/api/login
GET http://localhost:8080/api/orders
POST http://localhost:8080/api/orders
GET http://localhost:8080/api/orders/1
POST http://localhost:8080/api/orders/1/cancel
```