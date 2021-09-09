package utils

import (
	"encoding/json"
	"github.com/je117er/tfs-03/lec-11/exercises/models"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func MakeRandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Serialize(resp []*models.InfoResponse) (string, error) {
	result, err := json.Marshal(resp)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func Deserialize(b []byte) ([]*models.InfoResponse, error) {
	var resp []*models.InfoResponse
	err := json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
