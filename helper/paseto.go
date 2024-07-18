package helper

import (
	"encoding/json" // Package untuk encoding dan decoding data dalam format JSON
	"fmt"           // Package untuk formatting teks dan output
	"time"          // Package untuk manipulasi waktu

	"aidanwoods.dev/go-paseto"                   // Package untuk mengimplementasikan spesifikasi PASETO (Platform-Agnostic Security Tokens)
	"go.mongodb.org/mongo-driver/bson/primitive" // Package untuk tipe data primitive dalam BSON
)

// Payload adalah struktur yang merepresentasikan payload token JWT.
type Payload struct {
	Id    primitive.ObjectID `json:"id"`    // Id adalah ID objek yang biasanya merupakan identitas unik dari entitas pengguna.
	Email string             `json:"email"` // Email adalah alamat email terkait dengan payload.
	Exp   time.Time          `json:"exp"`   // Exp adalah waktu kedaluwarsa (expiration time) dari token.
	Iat   time.Time          `json:"iat"`   // Iat adalah waktu kapan token di-generate (issued at).
	Nbf   time.Time          `json:"nbf"`   // Nbf adalah waktu kapan token mulai berlaku (not before).
}

func Encode(id primitive.ObjectID, email, privateKey string) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))
	token.Set("id", id)
	token.SetString("email", email)
	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKey)
	return token.V4Sign(secretKey, nil), err
}

func Decode(publicKey string, tokenstring string) (payload Payload, err error) {
	var token *paseto.Token
	var pubKey paseto.V4AsymmetricPublicKey
	pubKey, err = paseto.NewV4AsymmetricPublicKeyFromHex(publicKey) // this wil fail if given key in an invalid format
	if err != nil {
		return payload, fmt.Errorf("Decode NewV4AsymmetricPublicKeyFromHex : %v", err)
	}
	parser := paseto.NewParser()                                // only used because this example token has expired, use NewParser() (which checks expiry by default)
	token, err = parser.ParseV4Public(pubKey, tokenstring, nil) // this will fail if parsing failes, cryptographic checks fail, or validation rules fail
	if err != nil {
		return payload, fmt.Errorf("Decode ParseV4Public : %v", err)
	}
	err = json.Unmarshal(token.ClaimsJSON(), &payload)
	return payload, err
}

func GenerateKey() (privateKey, publicKey string) {
	secretKey := paseto.NewV4AsymmetricSecretKey() // don't share this!!!
	publicKey = secretKey.Public().ExportHex()     // DO share this one
	privateKey = secretKey.ExportHex()
	return privateKey, publicKey
}
