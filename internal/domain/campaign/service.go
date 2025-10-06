package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
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
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}