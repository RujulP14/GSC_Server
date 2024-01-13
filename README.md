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

### Installation

1. **Clone the repository:**

```bash
   git clone https://github.com/yourusername/server-api.git
```

2. **Navigate to the project directory:**

```bash
   cd server-api
```

3. **Install dependencies:**

```bash
   go mod download
```

4. **Run the application:**

```bash
    go run server.go
```

### API Endpoints

###### Base URL:

```
http://localhost:8080
```

#### User

1. **Create User:**

```
POST /api/user/
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "secretpassword"
}
```

2. **Update User:**

```
PUT /api/user/:id
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "secretpassword"
}
```

3. **Delete User:**

```
DELETE /api/user/:id
```

4. **Get All user:**

```
GET /api/user/
```

5. **Get Particular User:**

```
GET /api/user/:id
```

6. **Login User:**

```
POST /api/user/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "secretpassword"
}
```

#### NGO

1. **Create NGO:**

```
POST /api/ngo/
Content-Type: application/json

{
  "email": "ngo@example.com",
  "password": "secretpassword",
  "profile: {}
}
```

2. **Update NGO:**

```
PUT /api/ngo/:id
Content-Type: application/json

{
    "profile": {
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

3. **Delete NGO:**

```
DELETE /api/ngo/:id
```

4. **Get All NGOs:**

```
GET /api/ngo/
```

5. **Get Particular NGO:**

```
GET /api/ngo/:id
```

6. **Login NGO:**

```
POST /api/ngo/login
Content-Type: application/json

{
  "email": "ngo@example.com",
  "password": "secretpassword"
}
```
