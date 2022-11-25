package schema

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/africhild/drapper/config"
	"github.com/xeipuuv/gojsonschema"
)

var (
	InitError         error
	TransactionSchema *gojsonschema.Schema
)

func InitTransaction(temp string) (*gojsonschema.Schema, error) {
	// pwd, _ := os.Getwd()
	// fmt.Println(pwd)
	// if _, err := os.Stat(pwd + "/github.com/africhild/drapper/schema/receipt-01.json"); err != nil {
	// 	log.Printf(err.Error())
	// }
	// files, err := ioutil.ReadDir(pwd)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return nil, err
	// }

	// for _, file := range files {
	// 	fmt.Println(file.Name(), file.IsDir())
	// }
	c := config.Config{
		Root:    "github.com/africhild/drapper",
		FileDir: "schema",
	}
	content, err := os.ReadFile(path.Clean(c.GetFile(temp, "json")))
	if err != nil {
		log.Printf(err.Error())
		InitError = errors.New("could not read schema file: 'receipt-01.json' cleanly")
		return nil, InitError
	}
	// fmt.Println(string(content))
	loader := gojsonschema.NewBytesLoader(content)
	TransactionSchema, err = gojsonschema.NewSchema(loader)
	if err != nil {
		InitError = errors.New("could not load schema file: 'receipt-01.json' cleanly")
		return nil, InitError
	}
	return TransactionSchema, nil
}

// func init() {
// initTransaction()
// }
