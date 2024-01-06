package main

type RegistrationRequest struct {
	Person Person `json:"person"`
	Status string `json:"status"`
}
