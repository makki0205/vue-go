# gojwt

## Usage

### setting

The default for salt is random 
The default for exp is `3600` Seconds 
```go
jwt.SetSalt("D79998A7-3F2B-4505-BE2B-6E68500AAE37")
jwt.SetExp(60 * 60 * 24)

```

### Generate

```go
claims := map[string]string{
  "user": "nick",
  "email": "nick@github.com",
}
token := jwt.Generate(claims)

```

### Decode

```go
payload, err := jwt.Decode(token)
if err != nil{
	panic(err)
}
fmt.Println(payload["user"])

```