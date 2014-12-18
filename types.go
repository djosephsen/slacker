package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	MsgId int
	Ws    *websocket.Conn
	Token string
}

type AuthResponse struct {
	Ok       bool      `json:"ok,omitempty"`
	Url      string    `json:"url,omitempty"`
	Self     Self      `json:"self,omitempty"`
	Channels []Channel `json:"channels,omitempty"`
	Team     Team
	Users    []User  `json:"users,omitempty"`
	Groups   []Group `json:"groups,omitempty"`
	IMs      []IM    `json:"ims,omitempty"`
	Bots     []Bot  `json:"bots,omitempty"`
}

type Bot struct{
      Icons map[string]interface{}
      Deleted bool
      Name string
      Id string
}

type Self struct {
	Created        int32                  `json:"created,omitempty"`
	ManualPresence string                 `json:"manual_presence,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Id             string                 `json:"id,omitempty"`
	Prefs          map[string]interface{} `json:"prefs,omitempty"`
}

type Team struct {
	Id                   string
	Name                 string
	Email_domain         string `json:"email_domain"`
	Domain               string
	Msg_edit_window_mins int32
	Over_storage_limit   bool
	Prefs                []Pref `json:"prefs,omitempty"`
}

type Pref map[string]interface{}
	 

type Channel struct {
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Created     int32      `json:"created,omitempty"`
	Creator     string   `json:"creator,omitempty"`
	IsArchived  bool     `json:"is_archived"`
	IsGeneral   bool     `json:"is_general"`
	IsMember    bool     `json:"is_member"`
	LastRead    string   `json:"last_read,omitempty"`
	LastSet     int32   `json:"last_set,omitempty"`
	Latest      Event    `json:"latest,omitempty"`
	Members     []string `json:"members,omitempty"`
	Purpose     Purpose
	Topic       Topic
	UnreadCount int `json:"unread_count,omitempty"`
}

type User struct {
	Id                  string
	Name                string
	Deleted             bool
	Color               string
	Profile             Profile
	Is_admin            bool
	Is_owner            bool
	Is_primary_owner    bool
	Is_restricted       bool
	Is_ultra_restricted bool
	Has_files           bool
}

type Profile struct {
	First_name string
	Last_name  string
	Real_name  string
	Email      string
	Skype      string
	Phone      string
	Image_24   string
	Image_32   string
	Image_48   string
	Image_72   string
	Image_192  string
}

type Group struct {
	id          string
	Name        string
	Is_group    bool
	Created     string
	Creator     string
	Is_archived bool
	Members     []string
	Topic       Topic
	Purpose     Purpose
}

type Topic struct {
	Value    string
	Creator  string
	Last_set string
}

type Purpose struct {
	Value    string
	Creator  string
	Last_set string
}

type IM struct {
	ID				string
	IsOpen          bool
	UnreadCount    int
	Latest         map[string]interface{}
	LastRead		string
	created		int32
	IsIm			bool
}
type Event struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
	User    string `json:"user"`
	Ts      string `json:"ts,omitempty"`
}
