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
	// Membuat token baru menggunakan paseto.NewToken().
	token := paseto.NewToken()

	// Menetapkan waktu token diterbitkan (issued).
	token.SetIssuedAt(time.Now())

	// Menetapkan waktu token dapat digunakan (not before).
	token.SetNotBefore(time.Now())

	// Menetapkan waktu token kedaluwarsa (expired).
	token.SetExpiration(time.Now().Add(2 * time.Hour))

	// Menetapkan nilai-nilai yang akan dimasukkan ke dalam token.
	token.Set("id", id)             // Menetapkan nilai ID sebagai objek primitive.ObjectID.
	token.SetString("email", email) // Menetapkan nilai email sebagai string.

	// Mengonversi private key dari format hexadecimal menjadi paseto.V4AsymmetricSecretKey.
	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKey)
	if err != nil {
		return "", err // Mengembalikan error jika terjadi kesalahan dalam pembuatan secret key.
	}

	// Mengenkripsi token menggunakan secret key V4 dan mengembalikan token yang dihasilkan serta error yang terjadi.
	return token.V4Sign(secretKey, nil), err
}

// Decode decodes a PASETO token string using a public key and returns the decoded payload.
func Decode(publicKey string, tokenstring string) (payload Payload, err error) {
	var token *paseto.Token
	var pubKey paseto.V4AsymmetricPublicKey

	// Convert the provided public key from hexadecimal format to a paseto.V4AsymmetricPublicKey.
	pubKey, err = paseto.NewV4AsymmetricPublicKeyFromHex(publicKey)
	if err != nil {
		return payload, fmt.Errorf("Decode NewV4AsymmetricPublicKeyFromHex: %v", err)
	}

	// Create a new PASETO token parser.
	parser := paseto.NewParser()

	// Parse the PASETO token using the public key.
	token, err = parser.ParseV4Public(pubKey, tokenstring, nil)
	if err != nil {
		return payload, fmt.Errorf("Decode ParseV4Public: %v", err)
	}

	// Unmarshal the token's claims (payload) into the `payload` struct.
	err = json.Unmarshal(token.ClaimsJSON(), &payload)
	return payload, err
}

// GenerateKey adalah fungsi yang menghasilkan kunci privat dan kunci publik menggunakan Paseto V4
func GenerateKey() (privateKey, publicKey string) {
	// Membuat kunci privat menggunakan Paseto V4 Asymmetric Secret Key
	secretKey := paseto.NewV4AsymmetricSecretKey() // jangan bagikan kunci ini!!!

	// Mengekspor kunci publik dari kunci privat dalam format hex
	publicKey = secretKey.Public().ExportHex() // Bagikan kunci publik ini

	// Mengekspor kunci privat dalam format hex
	privateKey = secretKey.ExportHex()

	// Mengembalikan kunci privat dan kunci publik sebagai string hex
	return privateKey, publicKey
}
