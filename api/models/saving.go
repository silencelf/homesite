package models

type Saving struct {
	ID       int
	UserId   int
	Desc     string
	Target   float64
	Achieved float64
	Details  []AddSavingModel
}

func NewSaving() Saving {
	return Saving{}
}

type AddSavingModel struct {
	UserID int     `json:"userId"`
	Desc   string  `json:"desc"`
	Amount float64 `json:"amount"`
}
