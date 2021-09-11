package models

type Saving struct {
	Total float64
}

func New() Saving {
	return Saving{}
}

type AddSavingModel struct {
	UserID int     `json:"userId"`
	Amount float64 `json:"amount"`
}
