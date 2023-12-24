package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			var payload map[string]interface{}
			err := json.NewDecoder(r.Body).Decode(&payload)
			if err != nil {
				http.Error(w, "Ошибка при разборе JSON-запроса", http.StatusBadRequest)
				log.Printf("Ошибка при разборе JSON-запроса: %v", err)
				return
			}

			if !isValidJSONFormat(payload) {
				http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
				log.Println("Ошибка: Неверный формат JSON")
				return
			}

			responseData := map[string]interface{}{
				"message": "Регистрация прошла успешно",
				"data":    payload,
			}
			jsonResponse, err := json.Marshal(responseData)
			if err != nil {
				http.Error(w, "Ошибка при формировании JSON-ответа", http.StatusInternalServerError)
				log.Printf("Ошибка при формировании JSON-ответа: %v", err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			w.Write(jsonResponse)

			log.Printf("Успешная регистрация: %v", payload)
		} else {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			log.Printf("Метод не поддерживается")
		}
	})

	log.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func isValidJSONFormat(data map[string]interface{}) bool {

	requiredFields := map[string]string{
		"username": "string",
		"email":    "string",
		"password": "string",
	}

	for field, expectedType := range requiredFields {
		value, ok := data[field]
		if !ok {
			return false
		}

		switch expectedType {
		case "string":
			if _, ok := value.(string); !ok {
				return false
			}
		default:
			return false
		}
	}

	return true
}
