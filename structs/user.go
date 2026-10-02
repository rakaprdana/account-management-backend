package structs

type UserResponse struct {
	Id        uint    `json:"id"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	IsDeleted bool    `json:"is_deleted"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Token     *string `json:"token,omitempty"`
}

type UserCreateRequest struct {
	Name      string `json:"name" binding:"required"`
	Username  string `json:"username" binding:"required" gorm:"unique;not null"`
	Email     string `json:"email" binding:"required" gorm:"unique;not null"`
	Password  string `json:"password" binding:"required"`
	IsDeleted bool   `json:"is_deleted"`
}

type UserUpdateRequest struct {
	Name      string `json:"name" binding:"required"`
	Username  string `json:"username" binding:"required" gorm:"unique;not null"`
	Email     string `json:"email" binding:"required" gorm:"unique;not null"`
	Password  string `json:"password,omitempty"`
	IsDeleted bool   `json:"is_deleted"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
