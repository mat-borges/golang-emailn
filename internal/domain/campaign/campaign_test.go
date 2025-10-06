package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "New Campaign"
	content = "This is a new campaign"
	contacts = []string{"contact1@example.com", "contact2@example.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	println(campaign.ID)
	assert.NotEmpty(campaign.ID)
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Len(campaign.Contacts, len(contacts))
}

func Test_NewCampaign_IDisNotNill(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign("", content, contacts)

	assert.Nil(campaign)
	assert.NotNil(err)
	assert.EqualError(err, "name is required")
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(name, "", contacts)

	assert.Nil(campaign)
	assert.NotNil(err)
	assert.EqualError(err, "content is required")
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(name, content, []string{})

	assert.Nil(campaign)
	assert.NotNil(err)
	assert.EqualError(err, "at least one contact is required")
}