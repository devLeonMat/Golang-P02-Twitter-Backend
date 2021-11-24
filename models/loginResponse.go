package models

/*LoginResponse response from jwt */
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
