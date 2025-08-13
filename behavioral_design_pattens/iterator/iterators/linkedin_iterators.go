package iterators

import (
	"fmt"
	"lld/behavioral_design_pattens/iterator/profile"
)

type LinkedinIterator struct {
	linkedin        profile.SocialNetwork
	contactType     string
	profileEmail    string
	emails          []string
	currentPosition int
}

func NewLinkedinIterator(linkedin profile.SocialNetwork, contactType, profileEmail string) (profile.ProfileIterator, error) {
	emails, err := linkedin.RequestProfileFriends(profileEmail, contactType)
	if err != nil {
		return nil, err
	}

	return &LinkedinIterator{
		linkedin:        linkedin,
		contactType:     contactType,
		profileEmail:    profileEmail,
		emails:          emails,
		currentPosition: 0,
	}, nil
}

func (i *LinkedinIterator) HasNext() bool {
	return i.currentPosition < len(i.emails)
}

func (i *LinkedinIterator) GetNext() (*profile.Profile, error) {
	if !i.HasNext() {
		return nil, fmt.Errorf("no more data")
	}
	nextEmail := i.emails[i.currentPosition]
	nextProfile, err := i.linkedin.RequestProfile(nextEmail)
	if err != nil {
		return nil, err
	}

	i.currentPosition++
	return nextProfile, nil
}

func (i *LinkedinIterator) Reset() {
	i.currentPosition = 0
}
