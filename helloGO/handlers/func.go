package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parsing del form
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Errore nel parsing del form", http.StatusBadRequest)
			return
		}

		// Converte tutti i valori in una mappa più semplice (string -> string)
		formData := make(map[string]string)
		for key, values := range r.Form {
			formData[key] = values[0] // Prende solo il primo valore per ogni chiave
		}

		// Converte la mappa in JSON
		jsonData, _ := json.Marshal(formData)

		// Stampa il JSON con tutti i valori del form nella risposta HTTP
		fmt.Fprintf(w, "Dati ricevuti: %s", string(jsonData))

		// Stampa il JSON nel terminale per debug
		fmt.Println("Dati ricevuti:", string(jsonData))

	} else {
		http.Error(w, "Metodo non consentito", http.StatusMethodNotAllowed)
	}
}
