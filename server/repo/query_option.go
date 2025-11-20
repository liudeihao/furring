package repo

import "gorm.io/gorm"

type QueryOption = func(*gorm.DB) *gorm.DB

func FilterByUsername(username string) QueryOption {
    return FilterByFieldEqual("username", username)
}

func FilterByFieldEqual(field, value string) QueryOption {
    return func(db *gorm.DB) *gorm.DB {
        return db.Where("? = ?", field, value)
    }
}

func FilterByUserid(userid uint) QueryOption {
    return func(db *gorm.DB) *gorm.DB {
        return db.Where("user_id = ?", userid)
    }
}
