package model

type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type PostCreateRequest struct {
    UserID  uint
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
}

type PostUpdateRequest struct {
    PostID  uint    `json:"id"`
    Title   *string `json:"title"`
    Content *string `json:"content"`
}

type CommentCreateRequest struct {
    UserID  uint   `json:"user_id"`
    PostID  uint   `json:"post_id"`
    Content string `json:"content"`
}
