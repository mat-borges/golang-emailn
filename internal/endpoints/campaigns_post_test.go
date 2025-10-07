package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct{
	mock.Mock
}

func (r *serviceMock) Create(newCampaign contract.NewCampaignDto) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}

var body = contract.NewCampaignDto{
	Name:     "New Campaign",
	Content:  "This is a new campaign",
	Emails: []string{"test@example.com"},
}

func Test_CampaignPosts_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaignDto) bool {
		if request.Name == body.Name && request.Content == body.Content && len(request.Emails) == len(body.Emails) {
			return true
		}
		return false
	})).Return("123x", nil)
	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/campaigns", &buf)
	rr := httptest.NewRecorder()

	json, status, err := handler.CampaignPost(rr, req)

	assert.Equal(http.StatusCreated, status)
	assert.NotNil(json)
	assert.Nil(err)
}

func Test_CampaignPosts_should_inform_err_when_exists(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/campaigns", &buf)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)
}