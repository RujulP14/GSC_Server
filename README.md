# NGO API

Welcome to the NGO API documentation. This API allows you to manage NGOs, including user registration, login, and NGO-related operations.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [API Endpoints](#api-endpoints)
  - [User](#user)
  - [NGO](#ngo)
  - [Chatbot](#chatbot)

## Introduction

This API provides functionality to manage NGOs and user registrations. It is built using Golang with the Gin framework and integrates with Firebase for authentication and Firestore for data storage.

## Features

- User creation, update, retrieval, deletion and login
- NGO creation, update, retrieval, deletion and login

## Getting Started

### Prerequisites

Before you begin, make sure you have the following installed:

- [Golang](https://golang.org/doc/install)
- [Firebase](https://firebase.google.com/docs/cli)
- [Gcloud](https://cloud.google.com/sdk/docs/install)

### Installation

1. **Clone the repository:**

```bash
   git clone https://github.com/yourusername/server-api.git
```

2. **Navigate to the project directory:**

```bash
   cd server-api
```

3. **Setup default credentials:**

```bash
   gcloud auth application-default login
```

4. **Install dependencies:**

```bash
   go mod download
```

5. **Run the server:**

```bash
    go run server.go
```

### API Endpoints

###### Base URL:

```
http://localhost:8000/api
```

#### User

1. **Signup User:**

```
POST /users/signup/
Content-Type: application/json

{
    "email": "vbhatnagar@gmail.com",
    "password": "1shaj",
    "profile": {
        "firstName": "Vanshaj",
        "lastName": "Bhatnagar"
    }
}
```

```
Response example
{
    "id": "qid9wiB26SeV5wjx3XIp",
    "message": "User created successfully"
}
```

2. **Login User:**

```
POST /users/login/
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "secretpassword"
}
```

```
Response example
{
    "message": "Login successful",
    "user": {
        "id": "ofbmb53DAlb7EMQAswDfXTJ1eno1",
        "email": "vbhatnagar@gmail.com",
        "password": "$2a$10$6J06PxI1tOVFuULQarN4..wA59NJLDhv0Mx4qfSh6gTu9OFZLpgwi",
        "profile": {
            "firstName": "Vanshaj",
            "lastName": "Bhatnagar",
            "dob": "",
            "profileImage": ""
        },
        "donations": null
    }
}
```

3. **Update User:**

```
PUT /users/:id/
Content-Type: application/json

{
    "firstName": "vanshaj"
}
```

```
Response example
{
    "message": "User updated successfully"
}
```

4. **Delete User:**

```
DELETE /users/:id/
```

```
Response example
{
    "message": "User deleted successfully"
}
```

5. **Get All user:**

```
GET /users/
```

6. **Get Particular User:**

```
GET /users/:id/
```

#### NGO

1. **Signup NGO:**

```
POST /ngos/signup/
Content-Type: application/json

{
    "email": "ngo@gmail.com",
    "password": "ngo123",
    "profile": {
        "ngoName": "NGO 123"
    }
}
```

```
Response example
{
    "id": "QKVyBXUgYgaClRnaKjnZ",
    "message": "NGO created successfully"
}
```

2. **Login NGO:**

```
POST /ngos/login/
Content-Type: application/json

{
    "email": "ngo@gmail.com",
    "password": "ngo123"
}
```

```
Response example
{
    "message": "Login successful",
    "ngo": {
        "id": "7NYWU0wKM6OFBCrxtpAoBkIL4dl1",
        "email": "ngo@gmail.com",
        "password": "$2a$10$OJS7M2MRmgeKgFaQe5AxF.T8WvtQZ8PPxNLP7JaEn.mQU07kqIFry",
        "profile": {
            "registrationNumber": "",
            "ngoName": "NGO 123",
            "worksFor": "",
            "address": "",
            "pincode": "",
            "city": "",
            "state": "",
            "country": "",
            "phoneNumber": "",
            "description": "",
            "logo": ""
        },
        "campaigns": null
    }
}
```

3. **Update NGO:**

```
PUT /ngos/:id/
Content-Type: application/json

{
    "profile": {
        "registrationNumber": "reg1234",
        "ngoName": "NGO 123",
        "worksFor": "Females",
        "address": "address 123, address 456",
        "pincode": "123456",
        "city": "Some City",
        "state": "Some State",
        "country": "Some Country",
        "phoneNumber": "9876543210",
        "description": "This NGO works for the welfare of women. Kindly support us with your contributions",
        "logo": "logo.png"
    }
}

```

4. **Delete NGO:**

```
DELETE /ngos/:id/
```

5. **Get All NGOs:**

```
GET /ngos/

```

6. **Get Particular NGO:**

```
GET /ngos/:id/
```

#### Chatbot

1. **Get Chatbot Response**

```
GET /chatbot
Content-type: application/json

{
    "inputText": "What is menstruation"
}
```

```
Response example
{
    "text": "Menstruation is the natural process of shedding the lining of the uterus (womb). It occurs when a woman is not pregnant. Menstruation is also known as a period.\n\nThe menstrual cycle is a complex process that is controlled by hormones. Each month, the ovaries release an egg (ovulation).
}
```
