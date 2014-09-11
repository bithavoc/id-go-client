# Bithavoc's Identity Client for Golang

 Golang client for Bithavoc.io's Identity hub.
 
 See [id.bithavoc.io](http://id.bithavoc.io)

## Setup

`go get github.com/bithavoc/id-go-client`

## Initialization

```go

import bithavocid "github.com/bithavoc/id-go-client"

func main() {
    client := bithavocid.NewClient("<app-id>")
}

```

## SignUp

SignUp for a new account.

Example:

```go
err := client.SignUp(bithavocid.SignUp{
    Email:    "bill@gates.com",
    Password: "msdos",
    Fullname: "Bill G",
})

if err != nil {
    fmt.Printf("Error signing up: %s\n", err.Error())
}

```

### Confirm

Confirm your email address.

Example:

```go
authCode, err := client.Confirm("<emailedCode>")

if err != nil {
    fmt.Printf("Error confirming account: %s\n", err.Error())
} else {
    fmt.Printf("Authorization Code: %s\n", authCode.Code)
}

```

### Login

Log-in into your account.

Example:

```go
authCode, err := client.LogIn(bithavocid.Credentials{
    Email: "bill@example.com",
    Password: "msdos",
})

if err != nil {
    fmt.Printf("Error: %s\n", err.Error())
} else {
    fmt.Printf("Authorization Code: %s\n", authCode.Code)
}

```

## Tests

`go test`

## License

MIT (See LICENSE)


