package domains

import (
	"errors"
	"fmt"
	"homesite/models"
)

type SavingDomain struct {
	Savings []models.Saving `json:"savings"`
}

func NewSavingService() SavingDomain {
	s := SavingDomain{}
	s.Savings = []models.Saving{}
	return s
}

func (s *SavingDomain) FindById(id int) (*models.Saving, error) {
	for i, v := range s.Savings {
		if v.ID == id {
			fmt.Println(&v == &s.Savings[i])
			return &s.Savings[i], nil
		}
	}

	return nil, errors.New("OBJECT NOT FOUND")
}
