package socialnetworks

import (
	"fmt"
	"lld/behavioral_design_pattens/iterator/helper"
	"lld/behavioral_design_pattens/iterator/iterators"
	"lld/behavioral_design_pattens/iterator/profile"
)

type Linkedin struct {
	profiles []*profile.Profile
}

func NewLinkedin(profiles []*profile.Profile) *Linkedin {
	return &Linkedin{
		profiles: profiles,
	}
}

func (l *Linkedin) RequestProfile(profileEmail string) (*profile.Profile, error) {
	helper.SimulateNetworkLatency()
	fmt.Printf("LinkedIn: Loading profile %q over the network...\n", profileEmail)
	return l.findProfile(profileEmail)
}

func (l *Linkedin) RequestProfileFriends(profileEmail, contactType string) ([]string, error) {
	helper.SimulateNetworkLatency()
	fmt.Printf("LinkedIn: Loading %q list of profile %q over the network...\n", contactType, profileEmail)

	p, err := l.findProfile(profileEmail)
	if err != nil {
		return nil, err
	}

	return p.GetContacts(contactType), nil
}

func (l *Linkedin) findProfile(profileEmail string) (*profile.Profile, error) {
	for _, p := range l.profiles {
		if p.GetEmail() == profileEmail {
			return p, nil
		}
	}
	return nil, fmt.Errorf("profile %q doesn't exists", profileEmail)
}

func (l *Linkedin) CreateFriendsIterator(profileEmail string) (profile.ProfileIterator, error) {
	return iterators.NewLinkedinIterator(l, "friends", profileEmail)
}

func (l *Linkedin) CreateCoworkersIterator(profileEmail string) (profile.ProfileIterator, error) {
	return iterators.NewLinkedinIterator(l, "coworkers", profileEmail)
}
