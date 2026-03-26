
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// @type   - ключ для создания типа данных
// @User   - название типа
// @struct - составной тип данных (несколько полей)
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)

	fmt.Println("http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server Closed")
	}
}

// @w - Объект в который мы записываем ответ клиенту
// @r - Указатель на структуру запроса, которая содержит информацию которую мы получаем от клиента
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home page")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка типа запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// @var  - Объявляем переменную
	// @user - Название переменной
	// @User - Тип данных данной переменной
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	fmt.Printf("Login attempt: email=%s, password=%s\n", user.Email, user.Password)

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "login success",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to decode json", http.StatusBadRequest)
		return
	}

	fmt.Printf("Register attempt: email=%s, password=%s\n", user.Email, user.Password)

	// @w.Header() — возвращает заголовки ответа (map с заголовками)
	// @Set — устанавливает заголовок
	// @"Content-Type" — заголовок, который говорит клиенту, как интерпретировать ответ
	// @"application/json" — означает, что ответ — это JSON
	w.Header().Set("Content-Type", "application/json")

	// Успешность создания объекта. Статус 201
	w.WriteHeader(http.StatusCreated)

	response := map[string]string{
		"message": "login success",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode json", http.StatusInternalServerError)
		return
	}
}
