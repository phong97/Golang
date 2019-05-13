package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/pat"
	"github.com/user/Golang/week1/nang_cao/cache"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("week1")

type account struct {
	username string `json:username`
	password string `json:password`
}

type Claims struct {
	username string `json:"username"`
	jwt.StandardClaims
}

func generatePassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePassword(hashedPwd string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func generateToken(username string, expireTime time.Time) (string, error) {
	claims := &Claims{
		username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func Register(w http.ResponseWriter, r *http.Request) {
	//get user
	var user account
	r.ParseForm()
	user.username = r.Form.Get("username")
	user.password = r.Form.Get("password")
	log.Println(user.username)

	//set user to redis
	client := cache.GetInstance().Client
	cache.Set(user.username, generatePassword(user.password), 0, client)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(user.username))
}

func Login(w http.ResponseWriter, r *http.Request) {
	//get user
	var user account
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.username = r.Form.Get("username")
	user.password = r.Form.Get("password")
	//check password
	client := cache.GetInstance().Client
	hashedPwd := cache.Get(user.username, client)
	if !comparePassword(hashedPwd, user.password) {
		w.WriteHeader(http.StatusUnauthorized)
		message := "Login failed"
		w.Write([]byte(message))
		return
	}

	//generate token
	expireTime := time.Now().Add(5 * time.Minute)
	token, err := generateToken(user.username, expireTime)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//create cookie
	cookie := &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expireTime,
	}
	http.SetCookie(w, cookie)
	message := "Login success"
	w.Write([]byte(message))
}

func GetDisplayName(w http.ResponseWriter, r *http.Request) {
	//get token from client
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get the JWT string from the cookie
	tokenString := cookie.Value

	//new claims
	claims := &Claims{}

	//parsing tokenString and save the results in claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(fmt.Sprintf("The access token correct")))
}

func main() {
	r := pat.New()
	r.Post("/register", Register)
	r.Post("/login", Login)
	r.Get("/display", GetDisplayName)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
	log.Println("Server run port 8080")
}
