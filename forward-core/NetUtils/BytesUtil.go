package NetUtils

import (
	"bytes"
	"encoding/binary"

	"github.com/axgle/mahonia"
)

func WriteBYTE(data *bytes.Buffer, val uint8) {
	//BYTE 长度：1
	buf := make([]byte, 1)
	buf[0] = byte(val)

	data.Write(buf)
}

func WriteWORD(data *bytes.Buffer, val uint16) {
	//WORD 长度：2
	buf := make([]byte, 2)
	buf[0] = byte(val)
	buf[1] = byte(val >> 8)

	data.Write(buf)
}

func WriteDWORD(data *bytes.Buffer, val uint32) {
	//DWORD 长度：4
	buf := make([]byte, 4)
	buf[0] = byte(val)
	buf[1] = byte(val >> 8)
	buf[2] = byte(val >> 16)
	buf[3] = byte(val >> 24)

	data.Write(buf)
}

func WriteTCHAR(data *bytes.Buffer, size int, val string) {
	//TCHAR 长度：由size指定
	buf := []byte(val)
	data.Write(buf)
	//
	for i := 0; i < size-len(buf); i++ {
		//剩余的补0
		data.WriteByte(0)
	}

}

func WriteUnicodeTCHAR(data *bytes.Buffer, size int, val string) {
	//Unicode TCHAR 长度：size*2
	realSize := size * 2
	buf := []byte(val)
	dataBuffer := make([]byte, realSize)
	k := 0
	for j := 0; j < len(buf); j++ {
		dataBuffer[k] = buf[j]
		dataBuffer[k+1] = byte(0)
		k = k + 2
	}

	data.Write(dataBuffer)
	//

}

func WriteInt(data *bytes.Buffer, val int) {
	//Byte 长度：4
	buf := make([]byte, 4)
	buf[0] = byte(val)
	buf[1] = byte(val >> 8)
	buf[2] = byte(val >> 16)
	buf[3] = byte(val >> 24)

	data.Write(buf)
}

func _ReadInt_(data *bytes.Buffer, val []byte) (int, error) {

	return data.Read(val)
}

func ReadWord(val []byte) uint16 {
	//binary.LittleEndian.Uint16(rData[4:6])
	return binary.LittleEndian.Uint16(val)
}

func ReadDWord(val []byte) uint32 {
	return binary.LittleEndian.Uint32(val)
}

func ReadTCHAR(val []byte) string {
	return string(val)
}

func UTF8_To_GBK(input string) string {
	enc := mahonia.NewEncoder("GBK")
	return enc.ConvertString(input)
}

func GBK_To_UTF8(input string) string {
	dec := mahonia.NewDecoder("UTF-8")
	return dec.ConvertString(input)
}

func U2W(input string) string {
	//网狐荣耀版本的服务端TCHAR编码需要这样转换
	dec := mahonia.NewDecoder("UTF-16LE")
	return dec.ConvertString(input)
}
