package profile

import (
	"fmt"
	"strings"
)

type Profile struct {
	name     string
	email    string
	contacts map[string][]string // contact type -> list of contacts
}

func NewProfile(name, email string, contacts ...string) (*Profile, error) {
	profile := &Profile{
		name:     name,
		email:    email,
		contacts: make(map[string][]string),
	}

	for _, contact := range contacts {
		contactType := "friend"
		contactEmail := ""
		parts := strings.Split(contact, ":")
		if len(parts) == 1 {
			contactEmail = parts[0]
		} else if len(parts) == 2 {
			contactType = parts[0]
			contactEmail = parts[1]
		} else {
			return nil, fmt.Errorf("invalid contact format")
		}
		profile.contacts[contactType] = append(profile.contacts[contactType], contactEmail)
	}
	return profile, nil
}

func (p *Profile) GetEmail() string {
	return p.email
}

func (p *Profile) GetName() string {
	return p.name
}

func (p *Profile) GetContacts(contactType string) []string {
	return p.contacts[contactType]
}
