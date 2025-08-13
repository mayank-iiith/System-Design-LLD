package iterators

import (
	"fmt"
	"lld/behavioral_design_pattens/iterator/profile"
)

type FacebookIterator struct {
	facebook        profile.SocialNetwork
	contactType     string
	profileEmail    string
	emails          []string
	currentPosition int
}

func NewFacebookIterator(facebook profile.SocialNetwork, contactType, profileEmail string) (profile.ProfileIterator, error) {
	emails, err := facebook.RequestProfileFriends(profileEmail, contactType)
	if err != nil {
		return nil, err
	}

	return &FacebookIterator{
		facebook:        facebook,
		contactType:     contactType,
		profileEmail:    profileEmail,
		emails:          emails,
		currentPosition: 0,
	}, nil
}

func (i *FacebookIterator) HasNext() bool {
	return i.currentPosition < len(i.emails)
}

func (i *FacebookIterator) GetNext() (*profile.Profile, error) {
	if !i.HasNext() {
		return nil, fmt.Errorf("no more data")
	}
	nextEmail := i.emails[i.currentPosition]
	nextProfile, err := i.facebook.RequestProfile(nextEmail)
	if err != nil {
		return nil, err
	}

	i.currentPosition++
	return nextProfile, nil
}

func (i *FacebookIterator) Reset() {
	i.currentPosition = 0
}
