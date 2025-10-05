package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/domain"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/infra"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// retryablePing attempts to ping the database with simple retry logic
func retryablePing(db *sql.DB, maxRetries int) error {
	for i := range maxRetries {
		if err := db.Ping(); err == nil {
			return nil
		}
		zap.L().Info("Ping attempt failed, retrying...\n", zap.Int("attempt", i+1))
		if i < maxRetries-1 {
			time.Sleep(time.Second * 2) // Wait 2 seconds between retries
		}
	}
	return fmt.Errorf("database ping failed after %d attempts", maxRetries)
}

func NewDBRepository(host, port, user, password, dbname string) (*Repository, error) {

	zap.L().Info("Database informations: ",
		zap.String("host", host),
		zap.String("port", port),
		zap.String("user", user),
		zap.String("password", password),
		zap.String("dbname", dbname),
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Use simple retryable ping with 3 attempts
	if err := retryablePing(db, 3); err != nil {
		return nil, err
	}

	con, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("gorm ile bağlantı kurulamadı : %v", err.Error())
	}

	migrator := con.Migrator()

	if !migrator.HasTable(&domain.RefreshToken{}) {
		if err := con.AutoMigrate(&domain.RefreshToken{}); err != nil {
			return nil, fmt.Errorf("db init edilemedi : %v", err.Error())
		}
	}

	if !migrator.HasTable(&domain.User{}) {
		if err := con.AutoMigrate(&domain.User{}); err != nil {
			return nil, fmt.Errorf("db init edilemedi : %v", err.Error())
		}
	}

	return &Repository{
		db: con,
	}, nil
}

// WRITE REPO
func (h *Repository) CreateUser(ctx context.Context, user domain.User) error {
	tx := h.db.WithContext(ctx).Model(&domain.User{}).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return infra.ErrTransaction
	}

	if err := tx.Create(&user).Error; err != nil {
		return TransformError(err)
	}

	if err := tx.Commit().Error; err != nil {
		return infra.ErrCommit
	}

	return nil
}

func (h *Repository) UpdateUser(ctx context.Context, user domain.User) error {
	tx := h.db.WithContext(ctx).Model(&domain.User{}).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return infra.ErrTransaction
	}

	// transaction := tx.Where("id = ?", user.ID).Where("deleted_at IS NULL").Omit("id").Updates(&user)
	transaction := tx.Where("id = ?", user.ID).Omit("id").Updates(&user)
	if err := transaction.Error; err != nil {
		return TransformError(err)
	}

	if transaction.RowsAffected == 0 {
		return infra.ErrUserNotExists
	}

	if err := transaction.Commit().Error; err != nil {
		return infra.ErrCommit
	}

	return nil
}

func (h *Repository) DeleteUser(ctx context.Context, id string) error {

	tx := h.db.WithContext(ctx).Model(&domain.User{}).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return infra.ErrTransaction
	}

	// Where("id = ? AND deleted_at IS NULL", id).
	transaction := tx.Model(&domain.User{}).Where("id = ?", id).Delete(&domain.User{})

	if err := transaction.Error; err != nil {
		return infra.ErrQueryError
	}

	if transaction.RowsAffected == 0 {
		return infra.ErrUserNotExists
	}

	if err := transaction.Commit().Error; err != nil {
		return infra.ErrCommit
	}

	return nil

}

func TransformError(err error) error {
	if pgerr, ok := err.(*pq.Error); ok {
		if pgerr.Code == "23505" {
			switch pgerr.Constraint {
			case "uni_users_nickname":
				return infra.ErrNicknameInUse
			case "uni_users_email":
				return infra.ErrEmailInUse
			case "uni_users_phone_number":
				return infra.ErrPhoneNumberInUse
			}
		}
	}
	return infra.ErrQueryError
}

func (h *Repository) SaveRefreshToken(ctx context.Context, refreshToken domain.RefreshToken) error {

	tx := h.db.WithContext(ctx).Model(&domain.RefreshToken{}).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return infra.ErrTransaction
	}

	if err := tx.Create(&refreshToken).Error; err != nil {
		return infra.ErrModelCreate
	}

	if err := tx.Commit().Error; err != nil {
		return infra.ErrCommit
	}

	return nil
}

//  wrong case
// tx := xx.Begin()
// tx.Update(xx).Error
// fmt.Println(tx.RowsAffected)   // result is 0

// correct case
// tx := xx.Begin()
// db := tx.Update(xx)
// fmt.Println(db.RowsAffected)   // result is n
