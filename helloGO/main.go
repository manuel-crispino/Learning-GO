package main

import (
	"helloGo/handlers"
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Errore nel caricamento del template ", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	// Servire file statici (CSS, JS)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Servire la pagina HTML
	http.HandleFunc("/", handler)

	// Route per gestire l'input
	http.HandleFunc("/submit", handlers.SubmitHandler)

	// Avviare il server
	port := ":8080"
	log.Println("🚀 Server avviato su http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
