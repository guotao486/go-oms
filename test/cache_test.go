package test

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"oms/pkg/cache"
	"oms/pkg/convert"
	"oms/pkg/setting"
	"testing"
	"time"

	"github.com/allegro/bigcache"
)

type User struct {
	Name string
	Age  int
}

//struct转换为[]byte
func Encode(data interface{}, order binary.ByteOrder) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, order, data)
	if err != nil {
		panic(err)
	}
	return buf.Bytes(), nil
}

//将[]byte转换为struct
func Decode(b []byte, order binary.ByteOrder, data interface{}) error {
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, order, data)
	return err
}

// 创建自定义缓存 go test -v -run TestInitCustom .\test\cache_test.go
func TestInitCustom(t *testing.T) {
	// 指定创建属性
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(time.Minute))
	if err != nil {
		t.Error(err)
	}
	defer cache.Close()

	// 设置缓存
	_ = cache.Set("key1", []byte("hello word"))

	buf := new(bytes.Buffer)
	//gob编码
	enc := gob.NewEncoder(buf)
	if err := enc.Encode([]User{{Name: "tom", Age: 18}}); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", buf.Bytes())
	_ = cache.Set("list", buf.Bytes())
	// 验证CleanWindow是否生效
	time.Sleep(10 * time.Second)
	// 获取缓存
	data, err := cache.Get("key1")
	if err != nil {
		t.Errorf("获取缓存失败:%v", err)
	}
	fmt.Printf("获取结果:%s\n", data)
	list, err := cache.Get("list")
	buf1 := new(bytes.Buffer)
	buf1.Write(list)
	dec := gob.NewDecoder(buf1)
	var info2 []User
	if err := dec.Decode(&info2); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("info2: %v\n", info2)
	fmt.Println("运行结束！")
}

// go test -v -run TestCache .\test\cache_test.go
func TestCache(t *testing.T) {
	setting := &setting.CacheSettingS{
		CacheStore: "bigCache",
	}
	cacheStore, err := cache.NewCacheStore(setting)
	if err != nil {
		fmt.Printf("NewCacheStore err: %v\n", err)
	}
	currency, _ := cacheStore.Engine.Get("currency")
	fmt.Printf("currency: %v\n", currency)
	if currency == nil {
		fmt.Println("currency empty")
	}
	cacheStore.Engine.Set("name", []byte("guotao"))

	// set
	list := []User{{Name: "tom", Age: 18}, {Name: "jer", Age: 25}}
	st := &convert.StructTo{V: list}
	buf, err := st.StructToBytes()
	if err != nil {
		fmt.Printf("StructToBytes err: %v\n", err)
	}
	fmt.Printf("list: %v\n", list)
	fmt.Printf("st.V: %v\n", st.V)
	fmt.Printf("buf: %v\n", buf)
	cacheStore.Engine.Set("list", buf)

	// get
	name, err := cacheStore.Engine.Get("name")
	if err != nil {
		fmt.Printf("get name err: %v\n", err)
	}
	fmt.Printf("name: %v\n", name)

	list2, err := cacheStore.Engine.Get("list")
	if err != nil {
		fmt.Printf("get list err: %v\n", err)
	}
	fmt.Printf("list2: %v\n", list2)
	var list3 []User

	bs := &convert.ByteTo{V: &list3}
	err = bs.ByteToStruct(list2)
	if err != nil {
		fmt.Printf("ByteToStruct err: %v\n", err)
	}
	fmt.Printf("list3: %v\n", list3)
	fmt.Printf("bs.V: %v\n", bs.V)
}
