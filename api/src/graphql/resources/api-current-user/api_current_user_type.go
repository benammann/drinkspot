package api_current_user

type CurrentUser struct {
	EmailAddress *string `json:"email_address"`
}

type CurrentUserResolver struct {
	CurrentUser
}

func (vr *CurrentUserResolver) EmailAddress() *string {
	return vr.CurrentUser.EmailAddress
}
