package model

// User represents a user structure for MongoDB.
type User struct {
    ID       string `json:"id,omitempty" bson:"_id,omitempty"`
    Username string `json:"username" bson:"username"`
    Password string `json:"-" bson:"password"` // The '-' tag means don't return the password field
}

