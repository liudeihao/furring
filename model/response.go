package model

type UserPublicInfoResponse struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
}

type UserPrivateInfoResponse struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

type UserLoginResponse struct {
    ID    uint   `json:"id"`
    Token string `json:"token"`
}

type UserRegisterResponse struct {
    ID    uint   `json:"id"`
    Token string `json:"token"`
}

type UserInfo struct {
    UserID   uint   `json:"user_id"`
    Username string `json:"username"`
}

type PostResponse struct {
    UserInfo
    ID      uint   `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

type PostBrief struct {
    ID    uint   `json:"id"`
    Title string `json:"title"`
}

type UserPostsResponse struct {
    Posts []PostBrief
}

type CommentResponse struct {
    CommentID uint   `json:"comment_id"`
    UserID    uint   `json:"user_id"`
    Username  string `json:"username"`
    Content   string `json:"content"`
}

type PostCommentsResponse struct {
    Comments []CommentResponse `json:"comments"`
}
