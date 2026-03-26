// =================== УРОК 1: ===================
// МАРШРУТИЗАЦИЯ И ОТБРАБОТКА РАЗЛИЧНЫХ ПУТЕЙ
// ===============================================

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)

	// Start Server
	fmt.Println("http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error; Server don't starting.")
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login page")
}
func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Register page")
}
