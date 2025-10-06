package campaign

import (
	"emailn/internal/contract"
)

type Service struct {
	Repo Repository
}

func (s *Service) Create(newCampaign contract.NewCampaignDto) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repo.Save(campaign)
	if err != nil {
		return "", err
	}

	return campaign.ID, nil
}