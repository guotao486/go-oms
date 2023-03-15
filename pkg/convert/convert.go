/*
 * @Author: GG
 * @Date: 2023-02-28 08:57:36
 * @LastEditTime: 2023-03-14 17:35:38
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\pkg\convert\convert.go
 *
 */
package convert

import (
	"bytes"
	"encoding/gob"
	"strconv"
)

/**
* @Author $
* @Description //TODO $
* @Date $ $
* @Param $
* @return $
**/

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}

type StructTo struct {
	V interface{}
}

// struct 转 []byte
func (st *StructTo) StructToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	//gob编码
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(st.V); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type ByteTo struct {
	V interface{}
}

// []byte 转 struct
func (bt *ByteTo) ByteToStruct(b []byte) error {
	buf := new(bytes.Buffer)
	buf.Write(b)
	dec := gob.NewDecoder(buf)

	if err := dec.Decode(bt.V); err != nil {
		return err
	}
	return nil
}
