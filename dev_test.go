package backend_test

import (
	"testing"

	"github.com/tee-am-ai/backend/config"
	helper "github.com/tee-am-ai/backend/helper"
)

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