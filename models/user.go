package models

type UserModels struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	RoleID    int    `json:"roleID"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserRegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	RoleID    int    `json:"roleID"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	UserID       int64                  `json:"userID"`
	RoleID       int                    `json:"roleID"`
	AccessToken  string                 `json:"accessToken"`
	RefreshToken string                 `json:"refreshToken"`
	Permission   []UserPermissionModels `json:"permission"`
}

type UserPermissionModels struct {
	ID     int64  `json:"id:"`
	Group  string `json:"group"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}
