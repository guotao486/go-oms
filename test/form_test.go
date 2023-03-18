package test

import (
	"fmt"
	"oms/internal/request"
	"reflect"
	"strings"
	"testing"
)

func getPostMapForm(formData map[string][]string, fieldName string, entity interface{}) map[string]interface{} {
	key := ""
	key2 := ""
	dicts := make(map[string]interface{})
	for k, v := range formData {
		if i := strings.IndexByte(k, '['); i >= 1 && k[0:i] == fieldName {
			if j := strings.IndexByte(k[i+1:], ']'); j >= 1 {
				key = k[i+1:][:j]
				if dicts[key] == nil {
					var entityV = entity
					dicts[key] = &entityV
				}

				if i2 := strings.IndexByte(k[i+1:][j:], '['); i2 >= 1 {
					if j2 := strings.IndexByte(k[i+1:][j:][i2+1:], ']'); j2 >= 1 {
						key2 = k[i+1:][j:][i2+1:][:j2]
						refEntityV := reflect.ValueOf(dicts[key]).Elem()
						if refEntityV.FieldByName(key2).IsValid() {
							refEntityV.FieldByName(key2).Set(reflect.ValueOf(v[0]))
							continue
						}

						for i := 0; i < refEntityV.NumField(); i++ {
							refEntityT := refEntityV.Type()
							field := refEntityT.Field(i)
							if field.Tag.Get("form") == key2 {
								fmt.Printf("refEntityV.FieldByName(field.Name): %v\n", refEntityV.FieldByName(field.Name))
								refEntityV.FieldByName(field.Name).Set(reflect.ValueOf(v[0]))
								break
							}
						}
						fmt.Printf("dicts[key]: %v\n", dicts[key])
					}
				}
			}
		}
	}
	return dicts
}

// go test -v -run TestGetPostMapForm .\test\form_test.go
func TestGetPostMapForm(t *testing.T) {
	data := make(map[string][]string)
	data["product[1][name]"] = []string{"product_name"}
	data["product[1][sku]"] = []string{"product_sku"}
	data["product[1][image]"] = []string{"product_image"}
	data["product[1][attribute]"] = []string{"product_attribute"}

	var Product request.CreateOrderProductRequest
	p := getPostMapForm(data, "product", Product)
	fmt.Printf("p: %v\n", p)

	for _, v := range p {
		fmt.Printf("v: %v\n", v)
	}
}
