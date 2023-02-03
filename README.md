# Order Management
## Introduction
This is simple order management CRUD project that uses mysql database.

| Methods  | Endpoint | Description
| ------------- | ------------- | ------------- |
| "POST"  | "/add/order"  | add order details |
| "POST"  | "/add/item?orderId={orderId}"  | add items to order | 
| "GET" | "/generate/invoice/{orderId}" | creates invoice id
| "GET  | "/get/order/{orderId}" | get order with items in sorted order | 
| "GET" | "/get/orders?status={status}" | get all orders based on status |
| "DELETE" | "/order/{orderId}/remove" | delete order |

## Prerequisites
- Go version 1.13 or higher
- Mysql

## Getting Started
1. Clone the repository ( and making sure GOPATH points to the directory).
```sh
git clone https://github.com/PGaur1398/OrderManagement.git
cd OrderManagement
```
2. Install Dependencies
```ssh
go mod download
```
3. Create Database and configure the connection in environment.env
```sh
CREATE DATABASE order_management;
USE order_management;
CREATE TABLE `orders` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    order_id VARCHAR(255) NOT NULL unique,
    order_status VARCHAR(255) NOT NULL,
    invoice_id VARCHAR(255),
    total_amount DECIMAL(10,2),
    currency_unit VARCHAR(255) NOT NULL
);
CREATE TABLE `order_items` (
    id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT NOT NULL,
    item_description VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    quantity INT NOT NULL,
    CONSTRAINT UC_order_items UNIQUE (order_id,item_description),
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    INDEX (order_id)
);
```
4. Start the server
```ssh
go run main.go
```

## TODOS
- containerize the project.
