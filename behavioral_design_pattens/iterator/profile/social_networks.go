package profile

type SocialNetwork interface {
	CreateFriendsIterator(profileEmail string) (ProfileIterator, error)
	CreateCoworkersIterator(profileEmail string) (ProfileIterator, error)
	RequestProfile(email string) (*Profile, error)
	RequestProfileFriends(email string, contactType string) ([]string, error)
}
