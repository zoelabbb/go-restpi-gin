package models

type Pasien struct {
	Id           int64  `gorm:"primarykey" json:"id"`
	NamaLengkap  string `gorm:"type:varchar(150)" json:"nama_lengkap"`
	NIK          string `gorm:"type:varchar(16)" json:"nik"`
	JenisKelamin string `gorm:"type:tinyint(1)" json:"jenis_kelamin"`
	TempatLahir  string `gorm:"type:varchar(100)" json:"tempat_lahir"`
	TanggalLahir string `gorm:"type:date" json:"tanggal_lahir"`
	Alamat       string `gorm:"type:text" json:"alamat"`
	NoHp         string `gorm:"type:varchar(15)" json:"no_hp"`
}
