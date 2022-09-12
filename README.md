# Simple e-Commerce App
This project repository demonstrates a customer perspective to experience simple purchase activity on a e-commerce system.

## Contents
* [Introduction](#introduction)
  * [Problem Statement](#problem-statement)
  * [Objective](#objective)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Run](#run)

## Introduction
### Problem Statement
| SKU   | Name  | Price | Inventory Qty |
| :---: | --- | ---   |          ---: |
| 120P90 | Google Home | $49.99 | 10 |
| 43N23P | MacBook Pro | $5,399.99 | 5 |
| A304SD | Alexa Speaker | $109.50 | 10 |
| 234234 | Raspberry Pi B | $30.00 | 2 |

**The system should have the following promotions:**
- Each sale of a MacBook Pro comes with a free Raspberry Pi B.
- Buy 3 Google Home for the price of 2.
- Buying more than 3 Alexa Speakers will have 10% discount on all Alexa Speakers.

### Objective
- [X] Develop a web server application to solve the problem statement using [Go](https://golang.org/doc/install) language.
- [X] Implement [GraphQL](https://graphql.org/) as API endpoint.
- [X] Implement CI to build and produce executable binary.
- [X] Implement CI to run test.

## Getting Started
### Prerequisites
1.  Please make sure you have [Go](https://golang.org/doc/install) installed on your system.
2.  Please make sure you have [Docker](https://docs.docker.com/engine/install/) installed on your system.

### Run
1. Executes following command in the root directory of project repository to run the program.
```bash
make run
```
2. Open [http://localhost:9090](http://localhost:9090) to open QraphQL browser or use [postman collection](docs/postman/postman_collection.json) to hit the API.