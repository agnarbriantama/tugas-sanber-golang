package request

type ChangePasswordRequest struct {
	OldPassword string `form:"old_password" `
	NewPassword string `form:"new_password" `
}
