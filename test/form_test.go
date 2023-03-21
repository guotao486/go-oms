/*
 * @Author: GG
 * @Date: 2023-03-18 16:13:54
 * @LastEditTime: 2023-03-21 11:47:41
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\test\form_test.go
 *
 */
package test

import (
	"encoding/json"
	"fmt"
	"oms/internal/request"
	"oms/pkg/app"
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
func getPostMapForm(formData map[string][]string, fieldName string) (map[string]interface{}, int) {
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

	return dicts, i
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
	data["product[1][name]"] = []string{"oms_222"}
	data["product[1679368693300-0][sku]"] = []string{"oms_111"}
	data["product[1679368693300-0][image]"] = []string{"9abcfe885065270eb7f6791574aad63f.png"}
	data["product[1679368693300-0][name]"] = []string{"oms_111"}
	data["product[1679368693300-0][attribute]"] = []string{"oms_111"}
	data["product[1][image]"] = []string{"f5d805ec3ad3e8c0cf099d91c3ec83e1.png"}
	data["product[1][sku]"] = []string{"oms_222"}
	data["product[1][attribute]"] = []string{"oms_222"}

	p := app.GetPostMapForm(data, "product")
	fmt.Printf("p: %v\n", p)

	for k, v := range p {
		Product := &request.CreateOrderProductRequest{}
		app.GetPostMapFormItem(data, v["i"], v["j"], "product", k, Product)
		fmt.Printf("Product: %v\n", Product)
	}
	for k, v := range p {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}

}
