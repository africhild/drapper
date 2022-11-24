package drapper

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"log"

	"github.com/africhild/drapper/config"
	"github.com/africhild/drapper/schema"
	"github.com/xeipuuv/gojsonschema"
)

func Validate(obj map[string]interface{}, tmpt string) error {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Fatalf("Could not marshal obj: '%v'", obj)
		return err
	}
	documentLoader := gojsonschema.NewBytesLoader(data)
	transactionSchema, err := schema.InitTransaction(tmpt)
	if err != nil {
		log.Fatal(err)
		return err
	}
	result, err := transactionSchema.Validate(documentLoader)
	if err != nil {
		log.Printf("There was a problem validating %s", obj)
		return err
	}

	if !result.Valid() {
		log.Printf("%s is not a valid document...\n", obj)
		for _, desc := range result.Errors() {
			log.Printf("--- %s\n", desc)
		}
		log.Fatalf("Exiting...")
		return errors.New("document not valid")
	}
	return nil
}

func Convert(obj map[string]interface{}, tmpt string, flag bool) (string, error) {
	if flag {
		err := Validate(obj, tmpt)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
	}

	c := config.Config{
		Root:    "github.com/africhild/drapper",
		FileDir: "template",
	}
	// t := template.Must(template.ParseFiles("../drapper/test.html"))
	t := template.Must(template.ParseFiles(c.GetFile(tmpt, "html")))

	var buf bytes.Buffer
	err := t.Execute(&buf, obj)
	if err != nil {
		log.Printf("--- %s\n", err)
		return "", err
	}
	return buf.String(), nil
}
