package spammer

import (
	"fmt"
	"lld/behavioral_design_pattens/iterator/profile"
)

type SocialSpammer struct {
	network profile.SocialNetwork
}

func NewSocialSpammer(network profile.SocialNetwork) *SocialSpammer {
	return &SocialSpammer{
		network: network,
	}
}

func (s *SocialSpammer) SendSpamToFriends(profileEmail, message string) error {
	fmt.Println("\nIterating over friends...")
	fmt.Println()

	itr, err := s.network.CreateFriendsIterator(profileEmail)
	if err != nil {
		return fmt.Errorf("sending spam to friends failed, error: %w", err)
	}

	for itr.HasNext() {
		nextProfile, err := itr.GetNext()
		if err != nil {
			return fmt.Errorf("sending spam to friends failed, error: %w", err)
		}
		s.sendMessage(nextProfile.GetEmail(), message)
	}
	return nil
}

func (s *SocialSpammer) SendSpamToCoworkers(profileEmail, message string) error {
	fmt.Println("\nIterating over coworkers...")
	fmt.Println()

	itr, err := s.network.CreateCoworkersIterator(profileEmail)
	if err != nil {
		return fmt.Errorf("sending spam to coworkers failed, error: %w", err)
	}

	for itr.HasNext() {
		nextProfile, err := itr.GetNext()
		if err != nil {
			return fmt.Errorf("sending spam to coworkers failed, error: %w", err)
		}
		s.sendMessage(nextProfile.GetEmail(), message)
	}
	return nil
}

func (s *SocialSpammer) sendMessage(email, message string) {
	fmt.Printf("Sent message to: %q. Message body: %q\n", email, message)
}
