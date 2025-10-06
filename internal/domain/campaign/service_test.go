package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


type repositoryMock struct{
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaignDto{
		Name:     "New Campaign",
		Content:  "This is a new campaign",
		Emails: []string{"contact1@example.com", "contact2@example.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(repositoryMock)
	service.Repo = repoMock
	repoMock.On("Save", mock.Anything).Return(nil)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	repoMock := new(repositoryMock)
	repoMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
		campaign.Content != newCampaign.Content ||
		len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
		})).Return(nil)

	service.Repo = repoMock

	service.Create(newCampaign)

	repoMock.AssertExpectations(t)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(repositoryMock)
	service.Repo = repoMock

	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repoMock := new(repositoryMock)
	repoMock.On("Save", mock.Anything).Return(errors.New("repository error"))

	service.Repo = repoMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}