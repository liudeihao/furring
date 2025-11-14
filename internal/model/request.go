package model

type UpdateUserRequest struct {
    Username *string `json:"username"`
}

type RegisterRequest struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required;min=6"`
}

type LoginRequest struct {
    Username *string `json:"username"`
    Email    *string `json:"email"`
    Password string  `json:"password" binding:"required"`
}
