package mongo

import (
	"errors"
	"fmt"

	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/domain"
	"github.com/EmreZURNACI/ProdFullReadyApp_User/internal/infra"
	"go.uber.org/zap"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Repository struct {
	collection *mongo.Collection
}

func NewDBRepository(user, password, appname, dbname, collection string) (*Repository, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s.gwkifwi.mongodb.net/?retryWrites=true&w=majority&appName=%s", user, password, dbname, appname)).
		SetServerAPIOptions(serverAPI)

	zap.L().Info("Database informations...",
		zap.String("user", user),
		zap.String("password", password),
		zap.String("appname", appname),
		zap.String("dbname", dbname),
		zap.String("collection", collection),
	)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, infra.ErrDBConnection
	}
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, infra.ErrDBPing
	}
	collect := client.Database(dbname).Collection(collection)
	return &Repository{
		collection: collect,
	}, nil

}

// READ REPO ASLINDA
// func (h *Repository) LoginWPhoneNumber(ctx context.Context, phonenumber, password string) (*domain.User, error) {

// 	filter := bson.M{"PhoneNumber": phonenumber}

// 	var user domain.User

// 	if err := h.collection.FindOne(ctx, filter).Decode(&user); err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return nil, ErrUserNotExists
// 		}
// 		return nil, ErrQueryError

// 	}

// 	if user.Password != password {
// 		return nil, ErrInvalidPhoneOrPassword
// 	}

// 	return &user, nil
// }

// func (h *Repository) LoginWEmail(ctx context.Context, email, password string) (*domain.User, error) {

// 	filter := bson.M{"Email": email}

// 	var user domain.User

// 	if err := h.collection.FindOne(ctx, filter).Decode(&user); err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return nil, ErrUserNotExists
// 		}
// 		return nil, ErrQueryError

// 	}

// 	if user.Password != password {
// 		return nil, ErrInvalidEmailOrPassword
// 	}

// 	return &user, nil
// }

// func (h *Repository) LoginWNickName(ctx context.Context, nickname, password string) (*domain.User, error) {

// 	filter := bson.M{"NickName": nickname}

// 	var user domain.User

// 	if err := h.collection.FindOne(ctx, filter).Decode(&user); err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return nil, ErrUserNotExists
// 		}
// 		return nil, ErrQueryError

// 	}

// 	if user.Password != password {
// 		return nil, ErrInvalidNicknameOrPassword
// 	}

// 	return &user, nil
// }

func (h *Repository) GetUser(ctx context.Context, id string) (*domain.User, error) {

	var user domain.User

	err := h.collection.FindOne(ctx, bson.M{"ID": "78025f78-e7d8-4396-bc7b-7839d6fa45f2"}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, infra.ErrUserNotExists
		}
		return nil, infra.ErrQueryError
	}

	return &user, err
}

func (h *Repository) GetUsers(ctx context.Context) ([]domain.User, error) {

	cursor, err := h.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []domain.User
	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
func (h *Repository) SignIn(ctx context.Context, field, password string) (*domain.User, error) {

	filter := bson.M{"NickName": field, "PhoneNumber": field, "Email": field}

	var user domain.User

	if err := h.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, infra.ErrUserNotExists
		}
		return nil, infra.ErrQueryError

	}

	if user.Password != password {
		return nil, infra.ErrInvalidNicknameOrPassword
	}

	return &user, nil
}
