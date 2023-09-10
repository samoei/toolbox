package concurency

import (
	"fmt"

	"github.com/spf13/cobra"
)

var UserProfileCommand = &cobra.Command{
	Use:   "userprofile",
	Short: "Runs the user profile program",
	Run:   runUserProfile,
}

type UserProfile struct {
	Id       int
	Comments []string
	Likes    int
	Friends  []int
}

func handleUserProfile(userId int) (*UserProfile, error) {
	comments, err := getComments(userId)

	if err != nil {
		return nil, err
	}
	likes, err := getLikes(userId)

	if err != nil {
		return nil, err
	}
	friends, err := getFriends(userId)

	if err != nil {
		return nil, err
	}

	userProfile := &UserProfile{
		Id:       userId,
		Comments: comments,
		Likes:    likes,
		Friends:  friends,
	}
	return userProfile, nil
}

func getComments(userId int) ([]string, error) {
	comments := []string{
		"Hello World",
		"I love this. Its amazing",
		"Honestly speaking this should not be the way to go",
		"Where is this country heading?",
	}

	return comments, nil
}

func getLikes(userId int) (int, error) {
	return 54, nil
}

func getFriends(userId int) ([]int, error) {
	return []int{102, 896, 563, 89, 256}, nil
}

func runUserProfile(_ *cobra.Command, _ []string) {
	userProfile, err := handleUserProfile(10)

	if err != nil {
		fmt.Print("Error:", err)
	}

	fmt.Println(userProfile)

}
