package models

type Movie struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Director *Director `json:"director"`
}
