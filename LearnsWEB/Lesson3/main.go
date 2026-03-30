package main

import (
	// Пакет для работы с контекстом — передает данные между
	// функциями, управляет таймаутами и отменой операций
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Указатель на соединение с PostgreSql
var db *pgx.Conn

func main() {
	// Строка подключения - "протокол://пользователь:пароль@хост:порт/имя_бд"
	connectionString := "postgres://appuser:secret@localhost:5432/authdb"
	// Переменная для ошибок
	var err error

	// Подключаемся к базе. Создаём пустой контекст без лимитов на таймауты, ибо нам важно
	// именно подключиться к базе. Условно мы можем принебречь добавлением чего-то, но
	// явно не подключением к базе, ибо без неё нас ждёт смерть.
	db, err = pgx.Connect(context.Background(), connectionString)
	// Проверка сбоев в подключении к БД
	if err != nil {
		log.Fatal("Failed to connected to DataBase: ", err)
	}

	// Делаем отложенное закрытие базы перед выходом из main().
	// defer полезен чтобы в случае ошибки в коде гарантированно закрыть базу
	defer db.Close(context.Background())
	fmt.Println("Connected to DataBase")

	// Server actions
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start Server!")
	} else {
		fmt.Println("Server started!")
	}
}

// @w - Объект в который мы записывает ответ клиенту
// @r - Информация которую мы получили от клиента
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the HomePage!")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	// @db.Exec - выполняет SQL запрос.
	// @_ - Игнор возвращаемых значений.
	// @$1 @$2 плейсхолдеры для значений в SQL
	_, err := db.Exec(context.Background(),
		"INSERT INTO users (email, password_hash) VALUES ($1, $2)",
		user.Email, user.Password)
	if err != nil {
		log.Printf("DB ERROR: %v", err) // Добавь эту строку
		http.Error(w, "Email already exists or invalide data", http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Ставим "status 201" - что объект успешно создан
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user created"})
}

// w - Что отдать клиенту
// r - Что получаем от клиента
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"Message": "login success"}); err != nil {
		http.Error(w, "Failed to encoding Json", http.StatusInternalServerError)
		return
	}
}
