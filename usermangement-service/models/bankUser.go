package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// bankUser represents the user model
type BankUser struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username           string             `bson:"username" json:"username"`
	Password           string             `bson:"password" json:"password"`
	LastPasswordUpdate time.Time          `bson:"lastPasswordUpdate" json:"lastPasswordUpdate"`
	Designation        string             `bson:"designation" json:"designation"`
	Roles              []string           `bson:"roles" json:"roles"`
}

type GrantRolesRequest struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

// SetDefaultRole sets the default role for a bankUser to "user"
