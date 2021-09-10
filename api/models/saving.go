package models

type Saving struct {
	Total float64
}

func New() Saving {
	return Saving{}
}

type AddSavingModel struct {
	Amount float64 `json:"amount"`
}
