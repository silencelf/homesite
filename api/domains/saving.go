package domains

import (
	"errors"
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

func (s *SavingDomain) FindById(id int) (models.Saving, error) {
	for _, v := range s.Savings {
		if v.ID == id {
			return v, nil
		}
	}

	return models.NewSaving(), errors.New("not_found")
}
