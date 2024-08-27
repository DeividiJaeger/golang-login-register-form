package main

import (
    "fmt"
    "html/template"
    "net/http"
)

func main() {
    // Handler para exibir a página de login
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        t, err := template.ParseFiles("templates/login.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        t.Execute(w, nil)
    })

    // handler para processar o login via HTMX
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            username := r.FormValue("username")
            password := r.FormValue("password")

            // validação de login
            if username == "admin" && password == "1234" {
                fmt.Fprintf(w, "Login successful! Welcome, %s.", username)
            } else {
                w.WriteHeader(http.StatusUnauthorized)
                fmt.Fprintf(w, "Invalid credentials, please try again.")
            }
        }
    })
    
    // arquivos estáticos como CSS
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Iniciar o servidor na porta 3000
    fmt.Println("Server running at http://localhost:3000")
    http.ListenAndServe(":3000", nil)
}
