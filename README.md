# authware

Authentication and access control for HTTP API integrated with an Identity Provider (Okta)

<div align="center">
  <p>
    <img src="https://user-images.githubusercontent.com/4137581/199227135-962f0780-33d3-4812-82cd-b24bc0ae8866.png" height="130px"/>
  </p>
  
  <h1>Authware</h1>
  
  <!-- <img src="https://github.com/checkaayush/authware/workflows/build/badge.svg?branch=master"/> -->

  <!-- <a href="https://goreportcard.com/report/github.com/checkaayush/authware">
    <img src="https://goreportcard.com/badge/github.com/checkaayush/authware"/>
  </a> -->
</div>

## Introduction

Authware is an authentication and authorization service based on Role-Based Access Control using [Casbin](https://casbin.io/).

### Error Codes

| Code Range | Description                                                                                                                             |
| ---------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| 2xx        | This range of response code indicates that request was fulfilled successfully and no error was encountered.                             |
| 400        | This return code indicates that there was an error in fulfilling the request because the supplied parameters are invalid or inadequate. |
| 401        | This return code means that we are not able to authenticate your request. Please re-check your username and password.                   |
| 5xx        | This response code indicates that there was an internal server error while processing the request.                                      |

## Development

> Pre-requisites: Golang v1.17+

1. Clone the repository locally.
2. Add .env file in the repository root by modifying the .env.template file as needed.
3. From repository root, run:
```bash
make start
```
4. API will be up and running at http://localhost:5000.

#### Dependencies

* [echo](https://echo.labstack.com/) - Web framework
* [casbin](https://casbin.io/) - Authorization library

## TODO:

- Add unit tests for the core functionality.
- 
