package utils

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

var keymaker = "keymaker"

type ErrorMessage struct {
	Err string `json:"error"`
}

func WriteError(msg string) []byte {
	data, err := json.Marshal(ErrorMessage{msg})
	if err != nil {
		panic(err)
	}
	return data
}

type AuthRequest struct {
	User string `json:"user"`
	Pwd  string `json:"pwd"`
}

type AuthResponse struct {
	User        string `json:"user"`
	Pwd         string `json:"pwd"`
	Roles       []int  `json:"roles"`
	AccessToken string `json:"accessToken"`
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//name := r.FormValue("programName")
	//password := r.FormValue("programPassword")
	name := req.User
	password := req.Pwd

	if len(name) == 0 || len(password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please provide name and password to obtain the token"))
		return
	}
	if (name == "neo" && password == "keanu") || (name == "morpheus" && password == "lawrence") {
		token, err := getToken(name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error generating JWT token: " + err.Error()))
		} else {
			w.Header().Set("Authorization", "Bearer "+token)
			w.WriteHeader(http.StatusOK)
			res := AuthResponse{name, password, []int{2001}, token}
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			//w.Write([]byte("Token: " + token))
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Name and password do not match"))
		return
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)

			w.Write(WriteError("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		name := claims.(jwt.MapClaims)["name"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)

		r.Header.Set("name", name)
		r.Header.Set("role", role)

		next.ServeHTTP(w, r)
	})
}

func getToken(name string) (string, error) {
	signingKey := []byte(keymaker)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"role": "redpill",
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	claims, err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error verifying JWT token: " + err.Error()))
		return
	}
	name := claims.(jwt.MapClaims)["name"].(string)
	//role := claims.(jwt.MapClaims)["role"].(string)

	res := AuthResponse{name, "", []int{2001}, tokenString}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(keymaker)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
