package database

import (
	"emailn/internal/domain/campaign"
	"errors"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)
	return nil
}

func (c *CampaignRepository) GetAll() ([]campaign.Campaign, error) {
	if len(c.campaigns) == 0 {
		return nil, errors.New("no campaigns found")
	}

	return c.campaigns, nil
}