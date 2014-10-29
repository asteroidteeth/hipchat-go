package hipchat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUserGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user/1", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		fmt.Fprintf(w, `{"id": 1, "links": {"self": "s"}, "mention_name": "m", "name": "n"}`)
	})
	want := &User{ID: 1, Name: "n", MentionName: "m", Links: UserLinks{Self: "s"}}

	user, _, err := client.User.Get(1)
	if err != nil {
		t.Fatalf("User.Get returns an error %v", err)
	}
	if !reflect.DeepEqual(want, user) {
		t.Errorf("Room.Get returned %+v, want %+v", user, want)
	}
}

func TestUserList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method %s, want %s", r.Method, m)
		}
		fmt.Fprintf(w, `
		{
			"items": [{"id": 1, "links": {"self": "s"}, "mention_name": "m", "name": "n"}], 
			"startIndex":1,
			"maxResults":1,
			"links":{"Self":"s"}
		}`)
	})
	want := &Users{Items: []User{User{ID: 1, Name: "n", MentionName: "m", Links: UserLinks{Self: "s"}}}, StartIndex: 1, MaxResults: 1, Links: UsersLinks{Self: "s"}}

	users, _, err := client.User.List()
	if err != nil {
		t.Fatalf("Room.List returns an error %v", err)
	}
	if !reflect.DeepEqual(want, users) {
		t.Errorf("Room.List returned %+v, want %+v", users, want)
	}
}

func TestUserMessage(t *testing.T) {
	setup()
	defer teardown()

	args := &UserMessageRequest{Message: "m", MessageFormat: "text"}

	mux.HandleFunc("/user/1/message", func(w http.ResponseWriter, r *http.Request) {
		if m := "POST"; m != r.Method {
			t.Errorf("Request method %s, want %s", r.Method, m)
		}
		v := new(UserMessageRequest)
		json.NewDecoder(r.Body).Decode(v)

		if !reflect.DeepEqual(v, args) {
			t.Errorf("Request body %+v, want %+v", v, args)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.User.Message(1, args)
	if err != nil {
		t.Fatalf("User.Message returns an error %v", err)
	}
}
