package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var salt string
var exp int = 3600

func init() {
	rand.Seed(time.Now().UnixNano())
	// Create a random string of random length for our salt
	randombytes := randomBytes(rand.Intn(64))
	if randombytes == nil {
		panic(errors.New("Error creating random salt"))
	}
	salt = string(randombytes)
}
func SetSalt(saltString string) {
	salt = saltString
}
func SetExp(IntExp int) {
	exp = IntExp
}

// Generate compiles and signs a JWT from a claim and an expiration time in seconds from current time.
func Generate(claim map[string]string) string {
	ex := time.Now().Add(time.Second * time.Duration(exp))
	expiration := ex.Format("2006-01-02 15:04:05")
	// Build the jwt header by hand since alg and typ aren't going to change (for now)
	header := base64.StdEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT","exp":"` + expiration + `"}`))
	// Build json payload and base64 encode it
	pl2, err := json.Marshal(claim)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	payload := base64.StdEncoding.EncodeToString([]byte(pl2))
	// Create a new secret from our salt and the paylod json string.
	secret := sha256wrapper(salt + string(pl2))
	// Build signature with the new secret and base64 encode it.
	hash := hmac256(header+"."+payload, secret)
	signature := base64.StdEncoding.EncodeToString([]byte(hash))
	jwt := header + "." + payload + "." + signature
	return jwt
}

// Decode decodes a JWT and returns the payload as a map[string]string.
func Decode(jwt string) (map[string]string, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return nil, errors.New("Invalid JWT Structure")
	}
	header, _ := base64.StdEncoding.DecodeString(parts[0])
	payload, _ := base64.StdEncoding.DecodeString(parts[1])
	signature, _ := base64.StdEncoding.DecodeString(parts[2])
	// JSON decode payload
	var pldat map[string]string
	if err := json.Unmarshal(payload, &pldat); err != nil {
		return nil, errors.New("Invalid Token")
	}
	// JSON decode header
	var headdat map[string]interface{}
	if err := json.Unmarshal(header, &headdat); err != nil {
		return nil, errors.New("Invalid Token")
	}
	// Extract and parse expiration date from header
	layout := "2006-01-02 15:04:05"
	exp := headdat["exp"].(string)
	expParsed, err := time.ParseInLocation(layout, exp, time.Now().Location())
	if err != nil {
		return nil, errors.New("Invalid Token")
	}
	// Check how old the JWT is.  Return an error if it is expired
	now := time.Now()
	if now.After(expParsed) {
		return nil, errors.New("Expired JWT")
	}
	// This probably should be one of the first checks, preceeding the date check.  If the signature of the JWT doesn't match there is likely fuckery afoot
	ha := hmac256(string(parts[0])+"."+string(parts[1]), sha256wrapper(salt+string(payload)))
	if ha != string(signature) {
		return nil, errors.New("Invalid JWT signature")
	}

	return pldat, nil
}

func randomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil
	}
	return b
}

func hmac256(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func sha256wrapper(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
