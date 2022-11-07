package schema

import (
	"errors"
	"github.com/xeipuuv/gojsonschema"
	"os"
	"path/filepath"
)

var (
	InitError error
	TransactionSchema  *gojsonschema.Schema
)
func initTransaction() {
	content, err := os.ReadFile(filepath.Clean("./schema/receipt-01.json"))
	if err != nil {
		InitError = errors.New("could not read schema file: 'receipt-01.json' cleanly.")
	}
	loader := gojsonschema.NewBytesLoader(content)
	TransactionSchema, err = gojsonschema.NewSchema(loader)
	if err != nil {
		InitError = errors.New("could not load schema file: 'receipt-01.json' cleanly.")
	}
}
func init() {
initTransaction()
}
