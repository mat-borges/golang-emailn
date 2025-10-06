package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contacts struct {
	Email string
}

type Campaign struct {
	ID          string
	Name        string
	Content     string
	Contacts    []Contacts
	CreatedOn   time.Time
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(emails) == 0 {
		return nil, errors.New("at least one contact is required")
	}

	contacts := make([]Contacts, len(emails))
	for i, email := range emails {
		contacts[i].Email = email
	}

	return &Campaign{
		ID:          xid.New().String(),
		Name:        name,
		Content:     content,
		Contacts:    contacts,
		CreatedOn:   time.Now(),
	}, nil
}