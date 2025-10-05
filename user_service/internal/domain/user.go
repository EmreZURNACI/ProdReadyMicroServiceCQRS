package domain

import "time"

type User struct {
	ID          string `json:"id" gorm:"PRIMARY KEY;type:uuid" deepcopier:"skip" params:"id" bson:"ID"`
	Firstname   string `json:"name" gorm:"type:VARCHAR(100)"`
	Lastname    string `json:"lastname" gorm:"type:VARCHAR(100)"`
	Nickname    string `json:"nickname" gorm:"type:VARCHAR(255);UNIQUE;NOT NULL;" validate:"required"`
	Email       string `json:"email" gorm:"type:VARCHAR(255);UNIQUE;NOT NULL;" validate:"required"`
	Password    string `json:"password" gorm:"type:VARCHAR(16);NOT NULL;" validate:"required"`
	PhoneNumber string `json:"phone_number" gorm:"type:VARCHAR(20);UNIQUE;NOT NULL;" validate:"required"`
	Address     []byte `json:"address" gorm:"type:jsonb;default:NULL"`
	Avatar      string `json:"avatar" gorm:"type:VARCHAR(255);default:NULL"`

	// U : User , M : Moderator , A : admin
	Role string `json:"role" gorm:"type:CHAR(1);default:U"`
	// Sabit uzunluklu veri varsa CHAR kullanılır bu kaç karakter olarak define edilirse edilsin
	// o kadar yer kaplar örn CHAR(255) yazarsan içine 10 karakterlik veri daha igrsen 255 kaplar.

	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP;NOT NULL;" deepcopier:"skip"`
	UpdatedAt time.Time `json:"updatedAt" deepcopier:"skip"`
	//CreatedAt string `json:"createdAt" gorm:"type:VARCHAR(30);default:NULL" deepcopier:"skip"`
	//UpdatedAt string `json:"updatedAt" gorm:"type:VARCHAR(30);default:NULL" deepcopier:"skip"`
	// DeletedAt string `json:"deletedAt" gorm:"type:VARCHAR(30);default:NULL" deepcopier:"skip"`
}

type Address struct {
	Country    string `json:"country"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Street     string `json:"street"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
}
