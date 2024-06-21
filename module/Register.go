package module

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
	model "github.com/tee-am-ai/backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/argon2"
)


func SignUp(db *mongo.Database, col string, insertedDoc model.User) (string, error) {
	if insertedDoc.NamaLengkap == "" || insertedDoc.Email == "" || insertedDoc.Password == "" || insertedDoc.Confirmpassword == ""{
		return "", fmt.Errorf("mohon untuk melengkapi data")
	}
	if err := checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return "", fmt.Errorf("email tidak valid")
	}
	userExists, _ := GetUserFromEmail(insertedDoc.Email, db)
	if insertedDoc.Email == userExists.Email {
		return "", fmt.Errorf("email sudah terdaftar")
	}
	if strings.Contains(insertedDoc.Password, " ") {
		return "", fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc.Password) < 8 {
		return "", fmt.Errorf("password terlalu pendek")
	}
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("kesalahan server : salt")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	user := bson.M{
		"namalengkap": insertedDoc.NamaLengkap,
		"email": insertedDoc.Email,
		"password": hex.EncodeToString(hashedPassword),
		"conformasiPassword": insertedDoc.Confirmpassword,
		"salt": hex.EncodeToString(salt),
	}
	_, err = InsertOneDoc(db, col, user)
	if err != nil {
		return "", err
	}
	return insertedDoc.Email, nil
}

