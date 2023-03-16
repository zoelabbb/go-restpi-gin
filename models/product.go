package models

type Product struct {
	Id          int64  `gorm:"primarykey" json:"id"`
	NamaProduct string `gorm:"type:varchar(300)" json:"nama_produk"`
	Deskripsi   string `gorm:"type:text(300)" json:"deskripsi"`
}
