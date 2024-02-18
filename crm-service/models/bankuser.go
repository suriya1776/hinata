package models

import "gopkg.in/mgo.v2/bson"

// bankUser represents the user model
type BankUser struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Username    string        `bson:"username" json:"username"`
	Password    string        `bson:"password" json:"password"`
	Designation string        `bson:"designation" json:"designation"`
	Roles       []string      `bson:"roles" json:"roles"`
}

type GrantRolesRequest struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

// SetDefaultRole sets the default role for a bankUser to "user"
