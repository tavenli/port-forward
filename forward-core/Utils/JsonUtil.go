package Utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/json-iterator/go"
	"reflect"

	"github.com/astaxie/beego/logs"
	"github.com/ugorji/go/codec"
)

var jsonIter = jsoniter.ConfigCompatibleWithStandardLibrary

func ToJson(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func FromJson(data []byte, t interface{}) error {
	return json.Unmarshal(data, t)
}

func ToJsonIterator(obj interface{}) ([]byte, error) {
	return jsonIter.Marshal(obj)
}

func FromJsonIterator(data []byte, t interface{}) error {
	return jsonIter.Unmarshal(data, t)
}

func MsgpEncode(obj interface{}) []byte {
	var mh codec.MsgpackHandle
	mh.MapType = reflect.TypeOf(obj)
	var buf bytes.Buffer
	enc := codec.NewEncoder(&buf, &mh)
	err := enc.Encode(obj)
	if err == nil {
		return buf.Bytes()
	} else {
		logs.Error(err)
		return nil
	}

}

func MsgpDecode(data []byte, obj interface{}) error {
	var mh codec.MsgpackHandle
	mh.MapType = reflect.TypeOf(obj)
	dec := codec.NewDecoder(bytes.NewReader(data), &mh)
	err := dec.Decode(&obj)
	return err

}

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func ShowObjAllProp(obj interface{}) {
	value_method := reflect.ValueOf(obj)
	obj_type := value_method.Type()

	fmt.Printf("输出对象的属性和方法\t%v\n", obj)

	fmt.Println("\tMethods...")

	for i := 0; i < value_method.NumMethod(); i++ {
		fmt.Printf("\t%d\t%s\n", i, obj_type.Method(i).Name)
	}

	value_element := reflect.ValueOf(obj).Elem()
	obj_element := value_element.Type()

	fmt.Println("\tFields...")
	for i := 0; i < value_element.NumField(); i++ {
		fmt.Printf("\t%d\t%s\n", i, obj_element.Field(i).Name)

	}
}
