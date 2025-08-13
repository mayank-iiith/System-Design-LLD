package main

// Iterator is a behavioral design pattern that allows sequential traversal through a complex data structure without exposing its internal details.
// With Iterator, clients can go over elements of different collections in a similar fashion using a single iterator interface.

import (
	"lld/behavioral_design_pattens/iterator/profile"
	socialnetworks "lld/behavioral_design_pattens/iterator/social_networks"
	"lld/behavioral_design_pattens/iterator/spammer"
	"log"
)

func main() {
	//network := socialnetworks.NewFacebook(createTestProfiles())
	network := socialnetworks.NewLinkedin(createTestProfiles())

	socialSpammer := spammer.NewSocialSpammer(network)
	err := socialSpammer.SendSpamToFriends("anna.smith@bing.com", "Hey! This is Anna's friend Josh. Can you do me a favor and like this post [link]?")
	if err != nil {
		log.Fatalf("ERROR: %+v\n", err)
	}

	err = socialSpammer.SendSpamToCoworkers("anna.smith@bing.com", "Hey! This is Anna's boss Jason. Anna told me you would be interested in [link].")
	if err != nil {
		log.Fatalf("ERROR: %+v\n", err)
	}
}

func createTestProfiles() []*profile.Profile {
	profiles := make([]*profile.Profile, 0)

	add := func(p *profile.Profile, err error) {
		if err == nil {
			profiles = append(profiles, p)
		}
	}

	add(profile.NewProfile("Anna Smith", "anna.smith@bing.com", "friends:mad_max@ya.com", "friends:catwoman@yahoo.com", "coworkers:sam@amazon.com"))
	add(profile.NewProfile("Maximilian", "mad_max@ya.com", "friends:anna.smith@bing.com", "coworkers:sam@amazon.com"))
	add(profile.NewProfile("Billie", "bill@microsoft.eu", "coworkers:avanger@ukr.net"))
	add(profile.NewProfile("John Day", "avanger@ukr.net", "coworkers:bill@microsoft.eu"))
	add(profile.NewProfile("Sam Kitting", "sam@amazon.com", "coworkers:anna.smith@bing.com", "coworkers:mad_max@ya.com", "friends:catwoman@yahoo.com"))
	add(profile.NewProfile("Liza", "catwoman@yahoo.com", "friends:anna.smith@bing.com", "friends:sam@amazon.com"))

	return profiles
}
