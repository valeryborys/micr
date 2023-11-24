package model

import (
	"crypto/rand"
	"encoding/base64"
)

type Id []byte

func GenerateId() Id {
	id := make(Id, 20)
	if _, err := rand.Read(id); err != nil {
		panic(err)
	}
	return id
}

func (id Id) String() string {
	return base64.RawURLEncoding.EncodeToString(id)
}

func DecodeId(s string) (Id, error) {
	decoded, err := base64.RawURLEncoding.DecodeString(s)
	return Id(decoded), err
}
