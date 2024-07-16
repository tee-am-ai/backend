package controller

import (
	"encoding/json"
	"net/http"

	"github.com/tee-am-ai/backend/helper"
	model "github.com/tee-am-ai/backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Fungsi untuk menambahkan ulasan
func AddUlasan(db *mongo.Database, col string, respw http.ResponseWriter, req *http.Request) {
	var ulasan model.Ulasan

		// Decode request body menjadi struct Ulasan
		err := json.NewDecoder(req.Body).Decode(&ulasan)
		if err != nil {
			helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
			return
		}

		// Validasi input
		if ulasan.NamaLengkap == "" || ulasan.Email == "" || ulasan.Rating == "" || ulasan.Message == "" {
			helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
			return
		}

			// Masukkan data ulasan ke database
		ulasanData := bson.M{
			"namalengkap": ulasan.NamaLengkap,
			"email":       ulasan.Email,
			"rating":      ulasan.Rating,
			"message":     ulasan.Message,
		}

		insertedID, err := helper.InsertOneDoc(db, col, ulasanData)
		if err != nil {
			helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : insert data, "+err.Error())
			return
		}

				// Response sukses
		resp := map[string]any{
			"message":    "ulasan berhasil ditambahkan",
			"insertedID": insertedID,
		}
		helper.WriteJSON(respw, http.StatusCreated, resp)
	}

	// Fungsi untuk mendapatkan semua ulasan
func GetAllUlasan(db *mongo.Database, col string, respw http.ResponseWriter, req *http.Request) {
	// Ambil semua data ulasan dari database
	ulasans, err := helper.GetAllDocs(db, col)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : get data, "+err.Error())
		return
	}
		// Response dengan data ulasan
		resp := map[string]any{
			"ulasan": ulasans,
		}
		helper.WriteJSON(respw, http.StatusOK, resp)
	}