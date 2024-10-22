package models

type Personalidades struct {
	Id       int    `json:id`
	Nome     string `json:"nome"`
	Historia string `json:historia`
}
