// models/bankuser.go
package models

import "gopkg.in/mgo.v2/bson"

// BankUser represents the user model
type BankUser struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Username    string        `bson:"username" json:"username"`
	Password    string        `bson:"password" json:"password"`
	Designation string        `bson:"designation" json:"designation"`
	Role        string        `bson:"role" json:"role"`
}

// SetDefaultRole sets the default role for a BankUser to "user"
