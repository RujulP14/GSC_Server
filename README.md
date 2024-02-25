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

7. **Get User's Favourite Blog:**

```
GET /users/:id/favourite-blog
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
    "contents": [
        {
            "role": "user",
            "parts": [
                {
                    "text": "What are you?"
                }
            ]
        }
    ]
}
```

```
Response example
{
    "text": "I am a large language model, trained by Google."
}
```

#### Videos

1. **Upload videos**

```
POST /videos/
Content-Type: application/json

{
    "id": "6",
    "title": "Yoga for Menstrual Health",
    "videoUrl": "https://www.youtube.com/watch?v=zcvo9VLVHWc&pp=ygUXbWVuc3RydWFsIGhlYWx0aCB2aWRlb3M%3D",
    "uploader": "Yoga Instructors",
    "uploadDate": "2024-02-26T13:00:00Z",
    "description": "Practice gentle yoga poses and sequences specifically designed to support menstrual health and alleviate menstrual discomfort. Follow along with experienced yoga instructors to enhance your well-being during menstruation.",
    "tags": ["Yoga", "Wellness"],
    "category": "Wellness",
    "thumbnailUrl": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTrvodvTIzC8wKEoU3h2Z-X3qjHQEuNCYjCJQ&usqp=CAU",
    "comments": [],
    "likes": 0,
    "transcripts": "This is a transcript of the yoga for menstrual health video."
}
```

```
Response example
{
    "id": "OrnuOI0xFqCkonRXAMtZ",
    "message": "Video created successfully"
}
```

2. **Get Videos**

```
GET /videos/
```

3. **Get Particular Video**

```
GET /videos/:id/
```

4. **Update Video**

```
PUT /videos/:id/
Content-Type: application/json

{
    "title": "Introduction to Machine Learning",
    "uploader": "JohnDoe",
    "uploadDate": "2024-02-10T08:00:00Z",
    "description": "This video provides an overview of machine learning concepts.",
    "tags": [
        "machine learning",
        "artificial intelligence"
    ],
    "category": "Education",
    "thumbnailUrl": "https://example.com/thumbnail.jpg",
    "comments": [],
    "likes": 0,
    "transcripts": "Here are the transcripts of the video..."
}
```

```
Response Example
{
    "message": "Video updated successfully"
}
```

5. **Delete Video**

```
DELETE /videos/:id/
```

6. **Comment a Video**

```
POST /videos/:id/comments/
Content-Type: application/json
{
    "userID": "user1",
    "content": "This is the first comment"
}
```

```
Response Example
{
    "message": "Comment added successfully"
}
```

7. **Delete a Comment**

```
DELETE /videos/:video_id/comments?:comment_id/
```


#### Donate

1. **Donate to an NGO**

```
POST /transactions/donate/
Content-Type: application/json
{
    "senderID": "baE9bUIhfAKLSFEzDauO",
    "receiverID": "HnHCQyvhHRxdP0oHhdDa",
    "campaignID": "DZC5nZOkF5ggNXAIEriJ",
    "amount": 100
}
```

```
Response Example
{
    "message": "Transaction data stored successfully"
}
```

#### Events

1. **Create event**

```
POST /events/
Content-Type: application/json
{
    "title": "Women's Empowerment Conference",
    "date": "2024-11-08",
    "day": "Saturday",
    "time": "9:00 AM",
    "location": "City Auditorium",
    "description": "Join us for a transformative day dedicated to women's empowerment. This conference features inspiring keynote speakers, interactive workshops, and networking opportunities. Explore topics such as leadership, career advancement, financial independence, and personal development.",
    "banner": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTgpbXOZ6GIienKI9UUVsgCDI72wVwGMneZLQ",
    "organizerId": "vVUHa8ErJixf0ANNVSty",
    "participants": []
}
```

2. **Get All Events**

```
GET /events/
```

3. **Get Particular Event**

```
GET /events/:id/
```

4. **Update Event**

```
PUT /events/:id/
Content type: application/json
{
  "title": "Community Cleanup Drive",
  "date": "2024-02-16",
  "day": "Saturday",
  "time": "10:00 AM",
  "location": "Central Park",
  "description": "Join us for a community cleanup drive to help keep our city clean and green!",
  "banner": "https://example.com/cleanup_drive_banner.jpg"
}
```

5. **Register for event**

```
POST /events/:id/register?user=:id/
```

6. **Delete an event**

```
DELETE /events/:id/
```

#### Campaigns

1. **Create Campaign**

```
POST /campaigns/
Content type: application/json
{
    "ngoID": "vVUHa8ErJixf0ANNVSty",
    "title": "Menstrual Hygiene Awareness Drive",
    "description": "Join our initiative to raise awareness about menstrual hygiene and break the stigma surrounding menstruation. We conduct educational workshops, distribute sanitary products, and promote open discussions about menstrual health.",
    "imageUrl": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRc_rEHKqnCO6Ti4ljVRoWx3FaupvymV1zJyg",
    "raisedMoney": 1500,
    "totalGoal": 5000,
    "donors": []
}
```

```
Response Example
{
    "id": "rfOYhSGYYCeoMBPawzDP",
    "message": "Campaign created successfully"
}
```

2. **Get all campaigns**

```
GET /campaigns/
```

3. **Get a particular campaign**

```
GET /campaigns/:id/
```

4. **Update a campaign**

```
PUT /initiatives/:id/
Content type: application/json
{
    "title": "Save Forests"
}1
```

5. **Update Image**

```
PATCH /campaigns/:id/update-image/
Content type: application/json
{
    "ImageURL": "https://www.example.com/image.jpeg"
}
```

6. **Add/remove donor**

```
PATCH /campaigns/:id/add-donor
PATCH /campaigns/:id/remove-donor
```


#### Blogs

1. **Create a blog**

```
Content type: application/json
{
    "id": "6",
    "title": "Menstrual Health and Nutrition: What You Need to Know",
    "content": "Nutrition plays a key role in menstrual health. In this blog post, we explore the impact of diet and nutrition on menstrual cycles, offering practical tips, dietary recommendations, and nutrient-rich recipes to support women's health and well-being.",
    "author": "Nutritionists",
    "image": "https://www.shutterstock.com/image-vector/set-feminine-hygiene-products-tampons-260nw-2359875625.jpg",
    "category": "Nutrition",
    "date": "2024-02-26T13:00:00Z"
}
```

```
Response Example
{
    "id": "YwmCEwZBahQv6X4ehxiO",
    "message": "Blog created successfully"
}
```


2. **Get Blogs**

```
GET /blogs/
```


#### Articles

1. **Create Article**

```
Content type: application/json
{
    "id": "7",
    "title": "Menstrual Health Advocacy: Bridging the Gap",
    "authorName": "Advocacy Group",
    "uploadDate": "2024-02-27T13:00:00Z",
    "content": "This article advocates for better menstrual health policies and programs to address the needs of marginalized communities. It calls for increased awareness, access to menstrual products, and support for menstrual health initiatives.",
    "tags": ["Menstrual Health Advocacy", "Policy", "Marginalized Communities"],
    "category": "Social Justice",
    "thumbnailUrl": "https://www.shutterstock.com/shutterstock/photos/2120616122/display_1500/stock-vector-women-s-day-banner-for-product-demonstration-pink-pedestal-or-podium-with-number-and-hearts-on-2120616122.jpg",
    "comments": [],
    "likes": 0
  }
```

```
Response Example
{
    "id": "GzON4geEUT0KwUKfbwwq",
    "message": "Article created successfully"
}
```

2. **Update an Article**

```
Content type: application/json

{
    "id": "article1",
    "title": "Sample Article",
    "authorName": "John Doe",
    "uploadDate": "2024-02-06T12:00:00Z",
    "content": "This is a sample article content.",
    "tags": [
        "tag1",
        "tag2",
        "tag3"
    ],
    "category": "Sample Category",
    "thumbnailURL": "https://example.com/thumbnail.jpg",
    "comments": [],
    "likes": 0
}
```

```
Response Example

{
    "message": "Article updated successfully"
}
```
