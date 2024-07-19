package helper

import (
	"context" // Package untuk mengelola konteks dalam operasi-operasi non-blocking
	"errors"  // Package untuk menangani error dengan cara yang lebih baik
	"fmt"     // Package untuk formatting teks dan output

	"github.com/tee-am-ai/backend/model"         // Package yang mungkin berisi definisi-definisi struktur data (model)
	"go.mongodb.org/mongo-driver/bson"           // Package untuk encoding dan decoding data dalam format BSON yang digunakan dalam MongoDB
	"go.mongodb.org/mongo-driver/bson/primitive" // Package untuk tipe data primitive dalam BSON
	"go.mongodb.org/mongo-driver/mongo"          // Package untuk melakukan operasi terkait MongoDB seperti menyimpan, mengambil, dan memperbarui data
	"go.mongodb.org/mongo-driver/mongo/options"  // Package untuk mengatur opsi-opsi dalam operasi MongoDB
)

// DBInfo adalah struktur untuk menyimpan informasi koneksi database.
type DBInfo struct {
	DBString string // DBString adalah string yang berisi informasi koneksi ke database, seperti URI atau alamat.
	DBName   string // DBName adalah string yang menentukan nama database yang akan digunakan.
}

func MongoConnect(mconn DBInfo) (db *mongo.Database, err error) {
	// Membuat koneksi ke MongoDB menggunakan string koneksi yang diberikan (mconn.DBString).
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mconn.DBString))
	if err != nil {
		return nil, err // Mengembalikan nil dan error jika terjadi kesalahan saat koneksi.
	}

	// Mengembalikan objek *mongo.Database yang merepresentasikan database yang dipilih.
	// Objek ini digunakan untuk melakukan operasi-operasi database di MongoDB.
	return client.Database(mconn.DBName), nil
}

// InsertOneDoc inserts a document into the specified MongoDB collection.
func InsertOneDoc(db *mongo.Database, col string, doc any) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		return
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

// GetUserFromEmail retrieves a user document from the "users" collection in MongoDB based on the provided email address.
func GetUserFromEmail(email string, db *mongo.Database) (doc model.User, err error) {
	collection := db.Collection("users")
	filter := bson.M{"email": email}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return doc, fmt.Errorf("email tidak ditemukan")
		}
		return doc, fmt.Errorf("kesalahan server")
	}
	return doc, nil
}

// GetAllDocs adalah fungsi generik yang mengambil semua dokumen dari koleksi MongoDB berdasarkan filter yang diberikan
func GetAllDocs[T any](db *mongo.Database, col string, filter bson.M) (docs T, err error) {
	// Membuat context tanpa batas waktu (timeout)
	ctx := context.TODO()

	// Mendapatkan koleksi dari database berdasarkan nama koleksi yang diberikan
	collection := db.Collection(col)

	// Mencari dokumen dalam koleksi berdasarkan filter yang diberikan
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		// Jika terjadi kesalahan saat pencarian, kembalikan nilai error
		return
	}

	// Menutup cursor setelah selesai digunakan
	defer cursor.Close(ctx)

	// Menyalin semua dokumen yang ditemukan ke dalam variabel docs
	err = cursor.All(context.TODO(), &docs)
	if err != nil {
		// Jika terjadi kesalahan saat penyalinan, kembalikan nilai error
		return
	}

	// Mengembalikan dokumen yang ditemukan dan nil
	return
}

func GetUserFromID(_id primitive.ObjectID, db *mongo.Database) (doc model.User, err error) {
	collection := db.Collection("users")
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return doc, fmt.Errorf("no data found for ID %s", _id)
		}
		return doc, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return doc, nil
}
