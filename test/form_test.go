/*
 * @Author: GG
 * @Date: 2023-03-18 16:13:54
 * @LastEditTime: 2023-03-18 19:15:30
 * @LastEditors: GG
 * @Description:
 * @FilePath: \go-oms\test\form_test.go
 *
 */
package test

import (
	"encoding/json"
	"fmt"
	"oms/internal/request"
	"reflect"
	"strings"
	"testing"
)

func DeepCopyByJson(src, dst interface{}) error {
	if tmp, err := json.Marshal(&src); err != nil {
		return err
	} else {
		err = json.Unmarshal(tmp, dst)
		return err
	}
}
func getPostMapForm(formData map[string][]string, fieldName string) (map[string]interface{}, int, int) {
	var key string
	var i, j int
	dicts := make(map[string]interface{})
	for k, _ := range formData {
		if i = strings.IndexByte(k, '['); i >= 1 && k[0:i] == fieldName {
			if j = strings.IndexByte(k[i+1:], ']'); j >= 1 {
				key = k[i+1:][:j]
				if dicts[key] == nil {
					dicts[key] = new(interface{})
				}
			}
		}
	}

	return dicts, i, j
}

func getPostMapFormItem(formData map[string][]string, i, j int, fieldName string, entity interface{}) {
	var key string
	var i2, j2 int
	for k, v := range formData {
		if i2 = strings.IndexByte(k[i+1:][j:], '['); i2 >= 1 && k[i+1:][:j] == fieldName {
			if j2 = strings.IndexByte(k[i+1:][j:][i2+1:], ']'); j2 >= 1 {
				key = k[i+1:][j:][i2+1:][:j2]

				refEntityV := reflect.ValueOf(entity).Elem()
				if refEntityV.FieldByName(key).IsValid() {
					refEntityV.FieldByName(key).Set(reflect.ValueOf(v[0]))
					continue
				}

				for i := 0; i < refEntityV.NumField(); i++ {
					field := refEntityV.Type().Field(i)
					if field.Tag.Get("form") == key {
						refEntityV.Field(i).Set(reflect.ValueOf(v[0]))
						break
					}
				}
			}
		}
	}
}

// go test -v -run TestGetPostMapForm .\test\form_test.go
func TestGetPostMapForm(t *testing.T) {
	data := make(map[string][]string)
	data["product[1][name]"] = []string{"product_name"}
	data["product[1][sku]"] = []string{"product_sku"}
	data["product[1][image]"] = []string{"product_image"}
	data["product[1][attribute]"] = []string{"product_attribute"}
	data["product[2][name]"] = []string{"product_name2"}
	data["product[2][sku]"] = []string{"product_sku2"}
	data["product[2][image]"] = []string{"product_image2"}
	data["product[2][attribute]"] = []string{"product_attribute2"}
	data["product[3][name]"] = []string{"product_name3"}
	data["product[3][sku]"] = []string{"product_sku3"}
	data["product[3][image]"] = []string{"product_image3"}
	data["product[3][attribute]"] = []string{"product_attribute3"}

	p, i, j := getPostMapForm(data, "product")
	fmt.Printf("p: %v\n", p)

	for k, _ := range p {
		Product := &request.CreateOrderProductRequest{}
		getPostMapFormItem(data, i, j, k, Product)
		p[k] = Product
	}
	for k, v := range p {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}

}
