package database

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"log"
)

type UserDatabaseService struct {
	db *badger.DB
}

type UserDatabase interface {
	GetUserNonce(address string) (string, error)
	SetUserNonce(address string, nonce string) error
}

func NewDatabase() *UserDatabaseService {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}

	return &UserDatabaseService{db}
}

func (u *UserDatabaseService) GetUserNonce(address string) (string, error) {
	key := []byte(fmt.Sprintf("user/%s", address))
	bytes := readFromDatabase(key, u.db)
	return string(bytes), nil
}

func (u *UserDatabaseService) SetUserNonce(address string, nonce string) error {
	return u.db.Update(func(txn *badger.Txn) error {
		key := []byte(fmt.Sprintf("user/%s", address))
		value := []byte(fmt.Sprintf("%s", nonce))

		err := txn.Set(key, value)
		return err
	})
}

func (u *UserDatabaseService) Close() {
	u.db.Close()
}

func readFromDatabase(key []byte, db *badger.DB) []byte {
	var bytes []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				print("Key not found")
			}
			return err
		}

		err = item.Value(func(val []byte) error {
			bytes = append([]byte{}, val...)
			return nil
		})
		return err
	})

	if err != nil {
		panic(err)
	}
	return bytes
}
