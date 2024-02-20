// models/reset_request.go

package models

// ResetRequest represents the request structure for initiating a password reset

type ResetRequest struct {
	Username        string `json:"username" binding:"required"`
	OldPassword     string `json:"old_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}
