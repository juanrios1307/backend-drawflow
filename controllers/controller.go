package controllers

import(
	"backend/models"
	"backend/configs"
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"golang.org/x/net/context"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/go-chi/chi"
)



func GetAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	dgClient := configs.NewClient()
	txn := dgClient.NewTxn()
	resp , err := txn.Query(context.Background(), queryCode)

	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp.Json)
}

const queryCode string = `
{
	getAll(func: has(Code)) {
		uid
		CodePython
		Code
	}
}
`

func Add(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var rawCode models.Code
	_ = json.NewDecoder(r.Body).Decode(&rawCode)
	p := models.Code { Code: rawCode.Code, CodePython:rawCode.CodePython }
	pb, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	dgClient := configs.NewClient()
	txn := dgClient.NewTxn()

	mutBuyers := &api.Mutation{
		CommitNow: true,
		SetJson: pb,
	}

	resp , err := txn.Mutate(context.Background(), mutBuyers)

	if err != nil {
		log.Fatal(err)
	}
	
	w.Write(resp.Json)

}



func GetOne(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var rawCode models.Code 
	_ = json.NewDecoder(r.Body).Decode(&rawCode)

	if id := chi.URLParam(r, "id"); id != "" {
		query := getQuery(id)
		dgClient := configs.NewClient()
		txn := dgClient.NewTxn()
		resp , err := txn.Query(context.Background(), query)

		if err != nil {
			log.Fatal(err)
		}
		w.Write(resp.Json)
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	
}

func getQuery( uid string )string{
	return fmt.Sprintf(getFileWithId,uid )
}

const getFileWithId string = `
{
	node(func: uid(%s)) {
	  uid
	  Code
	  CodePython
	}
}
  `
