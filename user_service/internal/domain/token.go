package domain

type RefreshToken struct {
	ID           string  `json:"id" gorm:"PRIMARY KEY;type:uuid"`
	UserUUID     User    `json:"user_uuid" gorm:"type:UUID;NOT NULL;references:ID;"`
	User         User    `gorm:"foreignKey:UserUUID;"`
	Token        string  `json:"token" gorm:"type:VARCHAR(255);NOT NULL;UNIQUE"`
	CreatedAt    float64 `json:"created_at" gorm:"NOT NULL;"`
	ExpirationAt float64 `json:"expiration_at" gorm:"NOT NULL;"`
}
type AccessToken struct {
	UUID         string  `json:"uuid"`
	Nickname     string  `json:"nickname"`
	ExpirationAt float64 `json:"expiration_at"`
}
