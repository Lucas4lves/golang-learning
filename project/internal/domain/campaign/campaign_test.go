package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign Name"
	content  = "Body"
	contacts = []string{
		"golang@gmail.com",
		"adieu@go.com",
	}
)

func TestNewCampaign(t *testing.T) {

	now := time.Now().Add(-time.Minute)

	//Act
	c, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(t, name, c.CampaignName, "Should be equal")
	assert.NotNil(t, c.GetXid())
	assert.NotNil(t, c.CreatedAt)
	assert.Equal(t, len(c.Contacts), len(contacts), "Should be equal")
	assert.Greater(t, c.CreatedAt, now)

}

// Arrange, Act, Assert

/*
	CampaignId   string
	CampaignName string
	CreatedAt    time.Time
	Content      string
	Contacts     []Contact
*/

/*	if c.CampaignId != "1" {
		t.Errorf("Expected 1, got %s", c.CampaignId)
	} else if c.CampaignName != name {
		t.Errorf("Expected %s", c.CampaignName)
	} else if c.Content != content {
		t.Errorf("Expected %s, got %s", content, c.Content)
	} else if len(c.Contacts) != len(contacts) {
		t.Errorf("Email listing error. Expected length: %d, got: %d", len(c.Contacts), len(contacts))
	}*/
