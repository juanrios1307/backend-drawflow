package controllers

import (
	"backend/configs"
	"context"
	"log"
	"net/http"
)

type Code struct {
    Code []string `json:"Code"`
	Uid string `json:"uid"`
}

const queryCode string = `
{
	getAll(func: has(Code)) {
		uid
		Code
	}
}`

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dgClient := configs.NewClient()
	txn := dgClient.NewTxn()
	resp, err := txn.Query(context.Background(), queryCode)

	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp.Json)
}