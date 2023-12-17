# Bank Management Service

This Go project provides a simple API for managing bank accounts. The project includes functionality to create a bank account, retrieve a list of all accounts, update account details and amounts, and remove a bank account.

## Project Functionalities

### Create a Bank Account

**Endpoint:** `POST /account/create`  
**Functionality:** Allows the creation of a new bank account by providing details such as name, government ID, address, and initial amount.

### List All Bank Accounts

**Endpoint:** `GET /accounts/list`  
**Functionality:** Retrieves a list of all bank accounts stored in the MongoDB database.

### Update Bank Account Details and Amount

**Endpoint:** `PUT /account/update/{id}`  
**Functionality:** Updates the details and amount of a specific bank account identified by its MongoDB ObjectID. Requires providing the updated parameters in the request body.

### Remove a Bank Account

**Endpoint:** `DELETE /account/remove/{id}`  
**Functionality:** Deletes a specific bank account identified by its MongoDB ObjectID.

## Running the Project

### Install Dependencies

1. Make sure you have Go installed on your machine.
2. Make sure you have MongoDB installed and running on your machine.
3. Install the necessary dependencies using:

   ```bash
   go get gofr.dev/pkg/gofr
   go get go.mongodb.org/mongo-driver/mongo
   ```
4. To run the Project execute the following command in the terminal:

   ```bash
   go run main.go


## API Testing with Postman

To facilitate the testing of the Bank Account Management API, you can use Postman, a popular API testing tool. Follow the steps below to test the various functionalities provided by the API.

### Prerequisites

1. **Postman Installed:**
   Ensure that you have Postman installed on your machine. If not, you can download it from [Postman's official website](https://www.postman.com/).

2. **Go Server Running:**
   Make sure that your Go server is running. If not, start the server using the following command in the terminal:

   ```bash
   go run main.go


### Testing Endpoints
Running Tests
1. Open Postman and create a new request collection.

2. Add requests for each endpoints with the specified details.

3. Click on the "Send" button to execute the requests and observe the responses.

4. Ensure that the responses match the expected results based on the API functionalities.
---
*Example for the POST, GET, PUT AND DELETE*

**Endpoint:** `POST /account/create`  
**Request:**
- **Method:** POST
- **URL:** `http://localhost:8080/account/create`
- **Headers:** Content-Type: application/json
- **Body:**
  ```json
  {
    "name": "Tarun Gupta",
    "govid": "123456789",
    "address": "Main Street, Mumbai, Maharashtra",
    "amount": 1000
  }

---

**Endpoint:** `GET /accounts/list`  
**Request:**
- **Method:** GET
- **URL:** `http://localhost:8080/accounts/list`

---

**Endpoint:** `PUT /account/update/{id}`  
**Request:**
- **Method:** PUT
- **URL:** Replace `{id}` with the actual ObjectID of the bank account you want to update.
  Example: `http://localhost:8080/account/update/5f5b4f1c7d6ff7e5884ea57c`
- **Headers:** Content-Type: application/json
- **Body:**
  ```json
  {
    "amount": 1500
  }

---

**Endpoint:** `DELETE /account/remove/{id}`  
**Request:**
- **Method:** DELETE
- **URL:** Replace `{id}` with the actual ObjectID of the bank account you want to remove.
  Example: `http://localhost:8080/account/remove/5f5b4f1c7d6ff7e5884ea57c`
