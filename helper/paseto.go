package helper

import (
	"encoding/json"
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payload struct {
	Id    primitive.ObjectID `json:"id"`
	Email string             `json:"email"`
	Exp   time.Time          `json:"exp"`
	Iat   time.Time          `json:"iat"`
	Nbf   time.Time          `json:"nbf"`
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

