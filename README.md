# drapper
validate and map object for PDF, JPEG, SVG, HTML files

## Installation

```bash
go get github.com/africhild/drapper
``` 
## Sample
```go
package main

import (
	"fmt"

	"github.com/africhild/drapper" // Import drapper
)

func main() {
	sample := map[string]interface{}{
		"id":                    "3232-4434333-er44333-re332222",
		"amount":                200.0,
		"type":                  "dr",
		"date":                  "2021-01-01T00:00:00Z",
		"status":                "pending",
		"transaction_reference": "1234567890",
		"user": map[string]interface{}{
			"id":           "1234-1234-1234-1234",
			"first_name":   "John",
			"last_name":    "Doe",
			"phone_number": "+233 50 123 4567",
			"email":        "johndoe@gmail.com",
		},
		"products": []string{
			"MacBook Pro",
			"iPhone 12",
			"Apple Watch",
		},
	}
	// err := drapper.Validate(sample, "test") // Validate the request
	// if err != nil {
	// 	// Handle error
	// 	fmt.Println(err)
	// }
	// "test" is the name of the template to validate against
	// (true|false) flag to indicate if the request should be validated
	buf, err := drapper.Convert(sample, "test", true)
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println(buf)
}
```
## Schema
```bash
test:
  id: string
  amount: float
  type: string
  date: string
  status: string
  transaction_reference: string
  user:
    id: string
    first_name: string
    last_name: string
    phone_number: string
    email: string
  products: array
```