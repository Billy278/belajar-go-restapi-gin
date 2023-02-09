package model

type Product struct {
	Id        int64  `gorm:"primarykey" json:"id"`
	Name      string `gorm:"type:varchar(200)" json:"name"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
}
