package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"bytes"
	"encoding/binary"
	"log"
	"time"
	"math/rand"
	"strconv"
	"math"
)

func ToJson(obj interface{}) string {
	data, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(data)
}

func ToMD5(str string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(str))
	return hex.EncodeToString(md5Hash.Sum(nil))
}

//整形转换成字节
func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, tmp)
	if err != nil {
		return nil
	}
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	if err != nil {
		return 0
	}
	return int32(tmp)
}

func CatchError() {
	if err := recover(); err != nil {
		log.Println("error:", err.(error).Error())
	}
}

func GetNonce() string {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomInt := rd.Int63n(99999999)
	if randomInt < 10000000 {
		randomInt += 10000000
	}
	return strconv.FormatInt(randomInt, 10)
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

func float64ToString(float float64) string {
	return strconv.FormatFloat(float, 'f', -1, 64)
}