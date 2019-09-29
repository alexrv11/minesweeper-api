package models

type InputGame struct {
	Dimension int `json:"dimension" binding:"required"`
	NumberOfBomb int `json:"number_of_bomb" binding:"required"`
}
