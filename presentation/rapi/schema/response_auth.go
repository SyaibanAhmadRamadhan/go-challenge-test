package schema

type ResponseAuth struct {
	ID          string `json:"id,omitempty"`
	RoleID      int    `json:"role_id,omitempty"`
	Username    string `json:"username,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Token       string `json:"token,omitempty"`
}
