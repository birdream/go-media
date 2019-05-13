package defs

// requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// VideoInfo
type VideoInfo struct {
	ID           string
	AuthorID     int
	Name         string
	DisplayCtime string
}

type Comment struct {
	ID      string
	VideoID string
	Author  string
	Content string
}
