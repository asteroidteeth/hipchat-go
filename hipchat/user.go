package hipchat

import (
	"fmt"
	"net/http"
)

// UserService gives access to the user related methods of the API.
type UserService struct {
	client *Client
}

type Users struct {
	Items      []User     `json:"items"`
	StartIndex int        `json:"startIndex"`
	MaxResults int        `json:"maxResults"`
	Links      UsersLinks `json:"links"`
}

// RoomsLinks represents the HipChat room list link.
type UsersLinks struct {
	Self string `json:"self"`
	Prev string `json:"prev"`
	Next string `json:"next"`
}

type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	MentionName string    `json:"mention_name"`
	Links       UserLinks `json:"links"`
}

type UserLinks struct {
	Self string `json:"self"`
}

type UserMessageRequest struct {
	Message       string `json:"message"`
	Notify        bool   `json:"notify,omitempty"`
	MessageFormat string `json:"message_format,omitempty"`
}

// List returns all the users authorized.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/get_all_users
func (s *UserService) List() (*Users, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "user", nil)
	if err != nil {
		return nil, nil, err
	}

	users := new(Users)
	resp, err := s.client.Do(req, users)
	if err != nil {
		return nil, resp, err
	}
	return users, resp, nil
}

func (s *UserService) Get(id int) (*User, *http.Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("user/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, err
	}
	return user, resp, nil
}

func (s *UserService) Message(id int, msgReq *UserMessageRequest) (*http.Response, error) {
	reqUrl := fmt.Sprintf("user/%d/message", id)
	req, err := s.client.NewRequest("POST", reqUrl, msgReq)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
