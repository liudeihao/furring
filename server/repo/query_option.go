package repo

import "gorm.io/gorm"

type QueryOption = func(*gorm.DB) *gorm.DB

func FilterByFieldEqual(field, value any) QueryOption {
    return func(db *gorm.DB) *gorm.DB {
        return db.Where("? = ?", field, value)
    }
}

func FilterByUsername(username string) QueryOption {
    return FilterByFieldEqual("username", username)
}
func FilterByPostID(postid uint) QueryOption {
    return FilterByFieldEqual("post_id = ?", postid)
}
func FilterByUserID(userid uint) QueryOption {
    return FilterByFieldEqual("user_id = ?", userid)
}
