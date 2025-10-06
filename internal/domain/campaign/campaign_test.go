package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name    = "New Campaign"
	content = "This is a new campaign"
	contacts = []string{"contact1@example.com", "contact2@example.com"}
	fake  = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

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

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign("", content, contacts)

	assert.Nil(campaign)
	assert.NotNil(err)
	assert.EqualError(err, "name must be at least 5 characters long")
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Nil(campaign)
	assert.NotNil(err)
	assert.EqualError(err, "name must be no more than 24 characters long")
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(name, "", contacts)

	assert.Nil(campaign)
	assert.NotNil(err)
	assert.EqualError(err, "content must be at least 5 characters long")
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(name, fake.Lorem().Text(1050), contacts)

	assert.Nil(campaign)
	assert.NotNil(err)
	assert.EqualError(err, "content must be no more than 1024 characters long")
}

func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(name, content, nil)

	assert.Nil(campaign)
	assert.NotNil(err)
	assert.EqualError(err, "contacts must be at least 1 characters long")
}

func Test_NewCampaign_MustValidateContactsEmail(t *testing.T) {
	assert := assert.New(t)
	contacts := []string{"invalid-email"}

	campaign, err := NewCampaign(name, content, contacts)

	assert.Nil(campaign)
	assert.NotNil(err)
	assert.EqualError(err, "email must be a valid email")
}

