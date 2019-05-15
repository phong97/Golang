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
	Username string `json:"username"`
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
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// get username, password from request
// then, generate password and save in redis
func Register(w http.ResponseWriter, r *http.Request) {
	//get user
	var user account
	r.ParseForm()
	user.username = r.Form.Get("username")
	user.password = r.Form.Get("password")
	log.Println(user.username)

	//set user to redis
	client := cache.GetInstance().Client
	status := cache.Set(user.username, generatePassword(user.password), 0, client)
	if !status {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(user.username))
}

// get username, password from request
// compare password and if ok, generate token and set cookie
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

// get token from cookie and check token
// if ok, return username
func GetDisplayName(w http.ResponseWriter, r *http.Request) {
	//get token from client
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprint("No Cookie")))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get the JWT string from the cookie
	tokenString := cookie.Value
	if tokenString == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprint("empty token")))
		return
	}

	//parsing tokenString and save the results in token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprint("token is invalid")))
			return
		}
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprint("token is invalid")))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprint("token is invalid")))
		return
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		w.Write([]byte(fmt.Sprintf("username : %v", claims.Username)))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprint("token is invalid")))
	}
}

func main() {
	r := pat.New()
	r.Post("/register", Register)
	r.Post("/login", Login)
	r.Get("/display", GetDisplayName)

	http.Handle("/", r)
	log.Println("Server run port 8080")
	http.ListenAndServe(":8080", nil)
}
