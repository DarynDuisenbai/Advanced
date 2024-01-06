package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Пример обработки JSON-запроса
			var registrationRequest RegistrationRequest
			err := json.NewDecoder(r.Body).Decode(&registrationRequest)
			if err != nil {
				http.Error(w, "Invalid JSON format", http.StatusBadRequest)
				log.Printf("Error decoding JSON request: %v", err)
				return
			}

			if !isValidJSONFormat(registrationRequest) {
				http.Error(w, "Invalid JSON structure", http.StatusBadRequest)
				log.Println("Error: Invalid JSON structure")
				return
			}

			responseData := map[string]interface{}{
				"status": "Registration successful",
				"data":   registrationRequest,
			}
			jsonResponse, err := json.Marshal(responseData)
			if err != nil {
				http.Error(w, "Error forming JSON response", http.StatusInternalServerError)
				log.Printf("Error forming JSON response: %v", err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			w.Write(jsonResponse)

			log.Printf("Registration successful: %+v", registrationRequest)
		} else if r.Method == http.MethodGet {

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("GET request successful"))
		} else {
			http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
			log.Printf("Method not supported")
		}
	})

	log.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Функция для валидации формата JSON
func isValidJSONFormat(request RegistrationRequest) bool {
	if request.Person.ID <= 0 ||
		request.Person.Name == "" ||
		request.Person.Age <= 0 ||
		request.Person.RegisteredAt == "" ||
		!isValidRFC3339Date(request.Person.RegisteredAt) {
		return false
	}
	return true
}

func isValidRFC3339Date(date string) bool {
	_, err := time.Parse(time.RFC3339, date)
	return err == nil
}
