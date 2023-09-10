package concurency

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var UserProfileCommand = &cobra.Command{
	Use:   "userprofile",
	Short: "Runs the user profile program",
	RunE:  runUserProfile,
}

type UserProfile struct {
	Id       int
	Comments []string
	Likes    int
	Friends  []int
}

type Response struct {
	data any
	err  error
}

func handleUserProfile(userId int) (*UserProfile, error) {
	var (
		respCh      = make(chan Response, 3)
		wg          = &sync.WaitGroup{}
		now         = time.Now()
		userProfile = &UserProfile{
			Id: userId,
		}
	)

	go getComments(userId, respCh, wg)
	go getLikes(userId, respCh, wg)
	go getFriends(userId, respCh, wg)

	wg.Add(3)
	wg.Wait()
	close(respCh)

	fmt.Println("Time 4 Now:", now)
	for resp := range respCh {
		if resp.err != nil {
			return nil, resp.err
		}
		fmt.Println(resp)

		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg
		}
	}

	return userProfile, nil
}

func getComments(userId int, respCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)
	comments := []string{
		"Hello World",
		"I love this. Its amazing",
		"Honestly speaking this should not be the way to go",
		"Where is this country heading?",
	}
	respCh <- Response{
		data: comments,
		err:  nil,
	}
	wg.Done()
}

func getLikes(userId int, respCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)
	respCh <- Response{
		data: 54,
		err:  nil,
	}
	wg.Done()
}

func getFriends(userId int, respCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	respCh <- Response{
		data: []int{102, 896, 563, 89, 256},
		err:  nil,
	}
	wg.Done()
}

func runUserProfile(_ *cobra.Command, _ []string) error {
	now := time.Now()
	fmt.Println("Time Now:", now)
	userProfile, err := handleUserProfile(10)

	if err != nil {
		return fmt.Errorf("An error occurred: '%v'", err)
	}

	fmt.Println(userProfile)

	fmt.Printf("It took %v to fetch user profile", time.Since(now))
	return nil
}
