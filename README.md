# Firebase JWT

This Go package provides a convienent method to parse auth JWT issued by [Firebase Authentication](https://firebase.google.com/docs/auth). 

## Auth token

An authentication token (auth token) is a computer generated code that helps to verify if user's identity is authentic or not. 
JWT or Json Web Token is one type of auth token which is used as Bearer token and self-signed. Read more at [JWT.io](https://jwt.io/) 

## Usage

### Installation guide

```bash
go get github.com/ubbn/firebasejwt@v0.1.0
```

### Import it

```go
import github.com/ubbn/firebasejwt

func main() {
    claims, err := firebasejwt.ParseFirebaseJWT("<token issued by firebase auth>")
    fmt.Println(claims["name"])
    fmt.Println(claims["email"])
}
```

### More

Package documentation can be found on [pkg.go.dev](https://pkg.go.dev/github.com/ubbn/firebasejwt)