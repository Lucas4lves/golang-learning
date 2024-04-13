package campaign

import (
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	CampaignId   string
	CampaignName string
	CreatedAt    time.Time
	Content      string
	Contacts     []Contact
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))

	for i, value := range emails {
		contacts[i].Email = value
	}

	return &Campaign{
		CampaignId:   xid.New().String(),
		CampaignName: name,
		CreatedAt:    time.Now(),
		Content:      content,
		Contacts:     contacts,
	}, nil
}

func (c *Campaign) GetXid() string {
	return c.CampaignId
}
