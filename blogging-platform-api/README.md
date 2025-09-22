# Blogging Platform API
Build a RESTful API for a personal blogging platform

Project URL: [https://roadmap.sh/projects/blogging-platform-api](https://roadmap.sh/projects/blogging-platform-api)

<image src="https://store.ardanlabs.com/cdn/shop/products/gopher-patch01-02_c7b3aa27-3c46-4c4f-823f-580bfea20dde.jpg?v=1661445057" width=1000 center>


## 1. Overwiew
You are required to create a simple RESTful API with basic CRUD operations for a personal blogging platform. CRUD stands for Create, Read, Update, and Delete. The goals of this project are to help you:

- Understand what the RESTful APIs are including best practices and conventions

- Learn how to create a RESTful API

- Learn about common HTTP methods like GET, POST, PUT, PATCH, DELETE

- Learn about status codes and error handling in APIs

- Learn how to perform CRUD operations using an API

- Learn how to work with databases

<br/>

---
---
## 2. Requirements

You should create a RESTful API for a personal blogging platform. The API should allow users to perform the following operations:

- Create a new blog post

- Update an existing blog post

- Delete an existing blog post

- Get a single blog post

- Get all blog posts

- Filter blog posts by a search term

Given below are the details for each API operation.

#### Create Blog Post
Create a new blog post using the `POST` method
```js
POST /posts
{
  "title": "My First Blog Post",
  "content": "This is the content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"]
}

```

Each blog post should have the following fields:

```js
{
  "title": "My First Blog Post",
  "content": "This is the content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"]
}

```

The endpoint should validate the request body and return a `201 Created` status code with the newly created blog post i.e.

```js
{
  "id": 1,
  "title": "My First Blog Post",
  "content": "This is the content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"],
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

or a `400 Bad Request` status code with error messages in case of validation errors.

#### Update Blog Post
Update an existing blog post using the PUT method

```js
PUT /posts/1
{
  "title": "My Updated Blog Post",
  "content": "This is the updated content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"]
}
```

The endpoint should validate the request body and return a `200 OK` status code with the updated blog post i.e.

```js
{
  "id": 1,
  "title": "My Updated Blog Post",
  "content": "This is the updated content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"],
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:30:00Z"
}
```

or a `400 Bad Request` status code with error messages in case of validation errors. It should return a `404 Not Found` status code if the blog post was not found.

#### Delete Blog Post
Delete an existing blog post using the DELETE method

```js
DELETE /posts/1
```

The endpoint should return a `204 No Content` status code if the blog post was successfully deleted or a `404 Not Found` status code if the blog post was not found.

#### Get Blog Post
Get a single blog post using the GET method

```js
GET /posts/1
```

The endpoint should return a `200 OK` status code with the blog post i.e.

```js
{
  "id": 1,
  "title": "My First Blog Post",
  "content": "This is the content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"],
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

or a `404 Not Found` status code if the blog post was not found.

#### Get All Blog Posts
Get all blog posts using the GET method

```js
GET /posts
```

The endpoint should return a `200 OK` status code with an array of blog posts i.e.

```js
[
  {
    "id": 1,
    "title": "My First Blog Post",
    "content": "This is the content of my first blog post.",
    "category": "Technology",
    "tags": ["Tech", "Programming"],
    "createdAt": "2021-09-01T12:00:00Z",
    "updatedAt": "2021-09-01T12:00:00Z"
  },
  {
    "id": 2,
    "title": "My Second Blog Post",
    "content": "This is the content of my second blog post.",
    "category": "Technology",
    "tags": ["Tech", "Programming"],
    "createdAt": "2021-09-01T12:30:00Z",
    "updatedAt": "2021-09-01T12:30:00Z"
  }
]
```

You don't have to implement pagination, authentication or authorization for this project. You can focus on the core functionality of the API.

While retrieving posts, user can also filter posts by a search term. You should do a wildcard search on the title, content or category fields of the blog posts. For example:

```js
GET /posts?term=tech
```

This should return all blog posts that have the term "tech" in their title, content or category. You can use a simple SQL query if you are using a SQL database or a similar query for a NoSQL database.

For databases, you can use:

- MySQL if you are using SQL

- MongoDB if you are using NoSQL

<br/>

---
---
## 3. Execution
To execute this project, run the following command in your terminal in order ro install all external dependencies, being in the `./number-guessing-game/` folder of the directory :
```sh
$ make init
```
Then run the following command:
```sh
$ go run .
```

<br/>

---
---
## 4. Troubleshooting

- If you encounter difficulties in code execution, try to run this command in the root of the directory in order to initialize the folder as a Go module :
```sh
$ go mod init booking-app
```
- To install Go compiler, follow the instructions on the link below : <br/>
[https://go.dev/doc/install](https://go.dev/doc/install)
- If you encounter a `missing go.sum` error, please run the following command:
```sh
$ go get number-guessing-game/utilities
```

