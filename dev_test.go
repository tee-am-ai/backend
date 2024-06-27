package backend_test

import (
	"testing"

	"github.com/tee-am-ai/backend/config"
	helper "github.com/tee-am-ai/backend/helper"
	// module "github.com/tee-am-ai/backend/module"
)

// var db = module.MongoConnect("MONGOSTRING", "team_ai")

func TestGenerateKey(t *testing.T) {
	privateKey, publicKey := helper.GenerateKey()
	t.Logf("PrivateKey : %v", privateKey)
	t.Logf("PublicKey : %v", publicKey)
}

// TestInsertOneDoc
func TestInsertOneDoc(t *testing.T) {
	var data = map[string]interface{}{
		"username": "teeamai",
		"password": "12345",
	}
	insertedDoc, err := helper.InsertOneDoc(config.Mongoconn, "users", data)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	t.Logf("InsertedDoc : %v", insertedDoc)
}

// func TestSignUp(t *testing.T) {
// 	var doc model.User
// 	doc.NamaLengkap = "Fedhira Syaila"
// 	doc.Email = "pedped@gmail.com"
// 	doc.Password = "pedi12345"
// 	doc.Confirmpassword = "pedi12345"
// 	email, err := module.SignUp(db, "user", doc)
// 	if err != nil {
// 		t.Errorf("Error inserting document: %v", err)
// 	} else {
// 		fmt.Println("Data berhasil disimpan dengan email:", email)
// 	}
// }

// func TestLogIn(t *testing.T) {
// 	var user model.User
// 	user.Email = "pedped@gmail.com"
// 	user.Password = "pedi12345"
// 	user, err := module.LogIn(db, user)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Berhasil LogIn : ", user.Email)
// 	}
// }
