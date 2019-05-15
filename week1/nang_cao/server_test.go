package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// tạo ra một request có nhúng username, password vào form.
// kiểm tra status code == 200 và body == username
func TestRegister(t *testing.T) {
	req, err := http.NewRequest("POST", "/register", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)

	req.ParseForm()
	req.Form.Set("username", "phong")
	req.Form.Set("password", "12345")

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `phong`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// tạo ra một request có nhúng username, password vào form.
// kiểm tra status code và body có như mong đợi
// cho 2 trường hợp login thành công và thất bại
func TestLogin(t *testing.T) {
	type args struct {
		username, password, expected string
		statusCode                   int
	}
	var data = []args{
		{"phong", "12345", "Login success", 200},
		{"phong", "123456", "Login failed", 401},
	}

	// dữ liệu test
	tests := []struct {
		name string
		args args
	}{
		{"login success", data[0]},
		{"Login failed", data[1]},
	}

	// chạy test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/login", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Login)

			req.ParseForm()
			req.Form.Set("username", tt.args.username)
			req.Form.Set("password", tt.args.password)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.args.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.args.statusCode)
			}

			expected := tt.args.expected
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		})
	}
}

// tạo ra một request có nhúng cookie
// kiểm tra status code == 200 và body có như mong đợi
// cho trường hợp token đúng, token rỗng và token sai
func TestGetDisplayName(t *testing.T) {
	type args struct {
		token, expected string
		expireTime      time.Time
		statusCode      int
	}
	expireTime := time.Now().Add(5 * time.Minute)
	token, _ := generateToken("phong", expireTime)
	var data = []args{
		{token, "username : phong", expireTime, 200},
		{"", "empty token", expireTime, 400},
		{token + "phong", "token is invalid", expireTime, 401},
		{"phong", "token is invalid", expireTime, 400},
	}

	// dữ liệu test
	tests := []struct {
		name string
		args args
	}{
		{"The access token correct", data[0]},
		{"empty token", data[1]},
		{"token is invalid", data[2]},
	}

	// chạy test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/display", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetDisplayName)

			req.AddCookie(&http.Cookie{
				Name:    "token",
				Value:   tt.args.token,
				Expires: tt.args.expireTime,
			})

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.args.statusCode {
				t.Errorf("%v: handler returned wrong status code: got %v want %v",
					tt.name, status, tt.args.statusCode)
			}

			expected := tt.args.expected
			if rr.Body.String() != expected {
				t.Errorf("%v: handler returned unexpected body: got %v want %v",
					tt.name, rr.Body.String(), expected)
			}
		})
	}

	// trường hợp không có cookie
	req, err := http.NewRequest("GET", "/display", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetDisplayName)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("No Cookie: handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	expected := "No Cookie"
	if rr.Body.String() != expected {
		t.Errorf("No Cookie: handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
