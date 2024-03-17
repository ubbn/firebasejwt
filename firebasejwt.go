package firebasejwt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const FirebaseKeysEndpoint = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"

// Read public certificate
func readPEM(kid string) string {
	resp, err := http.Get(FirebaseKeysEndpoint)
	if err != nil {
		return ""
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	publicCertificates := make(map[string]string)
	err = json.Unmarshal(body, &publicCertificates)
	if err != nil {
		return ""
	}

	return publicCertificates[kid]
}

// Do basic format checks on token string
func verifyBasics(tokenString string) error {
	if tokenString == "" {
		return fmt.Errorf("token string is empty")
	}
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return fmt.Errorf("wrong number of parts")
	}
	return nil
}

// Parse base64 encoded JWT issued by Firebase authenticator service
// and return token's claim in a map
func ParseFirebaseJWT(tokenString string) (map[string]interface{}, error) {
	err := verifyBasics(tokenString)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is as expected
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		kid, ok := token.Header["kid"]
		if !ok {
			return nil, fmt.Errorf("key id is missing in JWT header")
		}

		pem := readPEM(kid.(string))
		pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, err
		}

		return pubKey, nil
	})

	if err != nil {
		if token.Claims != nil {
			claims := token.Claims.(jwt.MapClaims)
			return claims, err
		}
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return claims, err
	}

	return claims, nil
}
