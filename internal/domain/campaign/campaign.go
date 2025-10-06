package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Contacts struct {
	Email string `validate:"email,required"`
}

type Campaign struct {
	ID          string `validate:"required"`
	Name        string `validate:"min=5,max=24,required"`
	Content     string `validate:"min=5,max=1024,required"`
	Contacts    []Contacts `validate:"min=1,dive,required"`
	CreatedOn   time.Time `validate:"required"`
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contacts, len(emails))
	for i, email := range emails {
		contacts[i].Email = email
	}

	campaign := &Campaign{
		ID:          xid.New().String(),
		Name:        name,
		Content:     content,
		Contacts:    contacts,
		CreatedOn:   time.Now(),
	}

	err := internalerrors.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}
	return nil, err

}