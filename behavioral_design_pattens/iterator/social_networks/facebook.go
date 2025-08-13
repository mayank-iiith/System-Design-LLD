package socialnetworks

import (
	"fmt"
	"lld/behavioral_design_pattens/iterator/helper"
	"lld/behavioral_design_pattens/iterator/iterators"
	"lld/behavioral_design_pattens/iterator/profile"
)

type Facebook struct {
	profiles []*profile.Profile
}

func NewFacebook(profiles []*profile.Profile) *Facebook {
	return &Facebook{
		profiles: profiles,
	}
}

func (f *Facebook) RequestProfile(profileEmail string) (*profile.Profile, error) {
	helper.SimulateNetworkLatency()
	fmt.Printf("Facebook: Loading profile %q over the network...\n", profileEmail)
	return f.findProfile(profileEmail)
}

func (f *Facebook) RequestProfileFriends(profileEmail, contactType string) ([]string, error) {
	helper.SimulateNetworkLatency()
	fmt.Printf("Facebook: Loading %q list of profile %q over the network...\n", contactType, profileEmail)

	p, err := f.findProfile(profileEmail)
	if err != nil {
		return nil, err
	}

	return p.GetContacts(contactType), nil
}

func (f *Facebook) findProfile(profileEmail string) (*profile.Profile, error) {
	for _, p := range f.profiles {
		if p.GetEmail() == profileEmail {
			return p, nil
		}
	}
	return nil, fmt.Errorf("profile %q doesn't exists", profileEmail)
}

func (f *Facebook) CreateFriendsIterator(profileEmail string) (profile.ProfileIterator, error) {
	return iterators.NewFacebookIterator(f, "friends", profileEmail)
}

func (f *Facebook) CreateCoworkersIterator(profileEmail string) (profile.ProfileIterator, error) {
	return iterators.NewFacebookIterator(f, "coworkers", profileEmail)
}
