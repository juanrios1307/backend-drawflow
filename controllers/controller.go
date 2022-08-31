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
	"os/exec"
)


//Obtener todos los registros en DB
func GetAll(w http.ResponseWriter, r *http.Request){

	//Cabecera del tipo de contenido enviado
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Obtener todos los Programas")
	
	//Creación del nuevo cliente, transacción y envio del query en la transacción
	dgClient := configs.NewClient()
	txn := dgClient.NewTxn()
	resp , err := txn.Query(context.Background(), queryCode)

	if err != nil {
		log.Fatal(err)
	}

	//Escritura de la respuesta
	w.Write(resp.Json)
}

//Query para traer todos los registros
const queryCode string = `
{
	getAll(func: has(Code)) {
		uid
		CodePython
		Code
	}
}
`

//Funcion para crear un nuevo registro
func Add(w http.ResponseWriter, r *http.Request){

	//Cabecera de contenido a responder
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Guardar Programa")

	//Creación de variable de codigo crudo
	var rawCode models.Code

	//Decodificado de json en request y conversion en variable creada
	_ = json.NewDecoder(r.Body).Decode(&rawCode)

	//Creación del modelo estructurado
	p := models.Code { Code: rawCode.Code, CodePython:rawCode.CodePython }
	
	//Codificación del modelo en JSON
	pb, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	//Creación del cliente y la transacción
	dgClient := configs.NewClient()
	txn := dgClient.NewTxn()

	//Creación de la mutacion con datos a crear
	mutBuyers := &api.Mutation{
		CommitNow: true,
		SetJson: pb,
	}

	//Envio de la mutación a la transacción
	resp , err := txn.Mutate(context.Background(), mutBuyers)

	if err != nil {
		log.Fatal(err)
	}
	
	//Escritura de la respuesta
	w.Write(resp.Json)

}


//Función para obtener solo un registro
func GetOne(w http.ResponseWriter, r *http.Request){

	//Seteado de la cabecera para especificar de que tipo se enviara la respuesta al front
	w.Header().Set("Content-Type", "application/json")
	
	fmt.Println("Obtener 1 programa")

	//Obtener data cruda que proviene del json del request
	var rawCode models.Code 
	_ = json.NewDecoder(r.Body).Decode(&rawCode)

	//Si existe un {id} que llego como parametro entonces...
	if id := chi.URLParam(r, "id"); id != "" {

		//Creación del query consulta
		query := getQuery(id)

		//Creación del cliente, transacción y respuesta dada por la BD
		dgClient := configs.NewClient()
		txn := dgClient.NewTxn()
		resp , err := txn.Query(context.Background(), query)

		if err != nil {
			log.Fatal(err)
		}

		//Escritura de respuesta
		w.Write(resp.Json)
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	
}

//Función para obtener query de busqueda 1 con formato correcto
func getQuery( uid string )string{

	//Retorno de query con uid seteado
	return fmt.Sprintf(getFileWithId,uid )
}

//Query para buscar solo 1 registro en la BD
const getFileWithId string = `
{
	node(func: uid(%s)) {
	  uid
	  Code
	  CodePython
	}
}
  `

//Función para ejecutar codigo 
//¡¡CREADO PARA EJECUCIÓN EN UBUNTU!!
func Execute(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("Ejecutar codigo")

	//Llamada a libreria de cmd para ejecutar un comando
	//El comando en este caso es python con paramatro el script en la ubicación dada
	cmd := exec.Command("python","/home/juanesrios/script.py")

	//Espera respuesta
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		json.NewEncoder(w).Encode("Syntaxis error")

		return
	}
	
	//Codificadión de la salida del CMD en formato JSON
	json.NewEncoder(w).Encode(string(out))

	//Escribe lo que envia al front
	w.Write(out)
}
