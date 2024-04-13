package campaign

import "time"

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

func NewCampaign(name string, content string, emails []string) *Campaign {

	contacts := make([]Contact, len(emails))

	for i, value := range emails {
		contacts[i].Email = value
	}

	return &Campaign{
		CampaignId:   "1",
		CampaignName: name,
		CreatedAt:    time.Now(),
		Content:      content,
	}
}
