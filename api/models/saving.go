package models

type Saving struct {
	ID       int                    `json:"id"`
	UserId   int                    `json:"userid"`
	Desc     string                 `json:"desc"`
	Target   float64                `json:"target"`
	Achieved float64                `json:"achieved"`
	Details  []AddSavingDetailModel `json:"details"`
}

type AddSavingModel struct {
	UserId int     `json:"userid"`
	Desc   string  `json:"desc"`
	Target float64 `json:"target"`
}

func NewSaving() Saving {
	return Saving{}
}

type AddSavingDetailModel struct {
	UserID int     `json:"userId"`
	Desc   string  `json:"desc"`
	Amount float64 `json:"amount"`
}
