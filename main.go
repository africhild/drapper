package drapper

import (
	"bytes"
	"encoding/json"
	"github.com/africhild/drapper/schema"
	"github.com/xeipuuv/gojsonschema"
	"html/template"
	"log"
	"path"
)
//func main() {
//	if schema.InitError != nil {
//		log.Fatal(schema.InitError)
//	}
//	sample := map[string]interface{}{
//			"id": "3232-4434333-er44333-re332222",
//			"amount": 200.0,
//			"type": "dr",
//			"date": "2021-01-01T00:00:00Z",
//			"status": "pending",
//			"transaction_reference": "1234567890",
//		    "user": map[string]interface{}{
//				"id": "1234-1234-1234-1234",
//				"first_name": "John",
//				"last_name": "Doe",
//				"phone_number": "+233 50 123 4567",
//				"email": "johndoe@gmail.com",
//			},
//			"products": []string{
//				"MacBook Pro",
//				"iPhone 12",
//				"Apple Watch",
//			},
//	}
//	validate(sample)
//}

func Validate(obj interface{}) {
	if schema.InitError != nil {
		log.Fatal(schema.InitError)
	}

	//jsonContent, err := ioutil.ReadFile("document.json")
	//schema, err := os.ReadFile(filepath.Clean("./schema/receipt-01.json"))
	//if err != nil {
	//	log.Fatalf("Could not read schema file: '%s' cleanly.", schema)
	//}
	//loadedSchema := gojsonschema.NewBytesLoader(schema)

	//schemaLoader := gojsonschema.NewReferenceLoader(schemaLocation[schemaVersion])
	data, err := json.Marshal(obj); if err != nil {
		log.Fatalf("Could not marshal obj: '%v'", obj)
	}
	documentLoader := gojsonschema.NewBytesLoader(data)
	// Validate the JSON data against the loaded JSON Schema
	result, err := schema.TransactionSchema.Validate(documentLoader)
	if err != nil {
		log.Printf("There was a problem validating %s", obj)
		log.Fatalf(err.Error())

	}
	// Check the validity of the result and throw a message if the document is valid or if it's not with errors.
	if result.Valid() {
		layout := path.Join("templates", "layout.html")
		body := path.Join("templates", "receipt-01.html")
		t := template.Must(template.ParseFiles(layout, body))
		var content bytes.Buffer

		err := t.Execute(&content, data)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("The document is valid. %s", content.String())
	} else {
		log.Printf("%s is not a valid document...\n", obj)
		for _, desc := range result.Errors() {
			log.Printf("--- %s\n", desc)
		}
		log.Fatalf("Exiting...")
		//return errors.New("document not valid")
	}
}

