package model

import "time"

// Message represents the structure of a message document in the database.
type Message struct {
    ID        string    `json:"id" bson:"_id,omitempty"`
    Sender    string    `json:"sender" bson:"sender"`
    Receiver  string    `json:"receiver" bson:"receiver"`
    Content   string    `json:"content" bson:"content"`
    Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

