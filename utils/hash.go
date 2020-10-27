package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

//对一个字符串数据进行md5哈希计算
func Md5HashString(data string) string {
	md5Hsah := md5.New()
	md5Hsah.Write([]byte(data))
	bytes := md5Hsah.Sum(nil)
	return hex.EncodeToString(bytes)
}


func MD5HashReader(reader io.Reader) (string,error) {
	md5Hash := md5.New()
	readerBytes,err :=ioutil.ReadAll(reader)

	//fmt.Println("读取到的文件",readerBytes)

	if err !=nil{
		return "",err
	}
	md5Hash.Write(readerBytes)
	hashBytes :=md5Hash.Sum(nil)
    return hex.EncodeToString(hashBytes),nil
}

//读取io流中的数据，并对数据进行hash计算
func SHA256HashReader(reader io.Reader) (string,error) {
	sha256Hash:= sha256.New()
	readerBytes,err :=ioutil.ReadAll(reader)
	if err != nil{
		return "", err
	}
	sha256Hash.Write(readerBytes)
	hashBytes := sha256Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}

//对区块数据进行SHA256Hash计算
//func SHA256HashBlock(block blockchain.Block) []byte {
func SHA256HashBlock(bs []byte) []byte {
	//1.将block结构体数据转换成[]byte类型
	/**heightBytes,_ := Int64ToByte(block.Height)
	timeStampBytes,_ := Int64ToByte(block.TimeStamp)
	versionBytes := StringToBytes(block.Version)

	var blockBytes []byte
	//byes.Jion 拼接
	bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		block.PrevHash,
		block.Data,
		versionBytes,
	},[]byte{})

	 */

	//2.将转换后的[]byte字节切片输入write方法
	sha256Hash := sha256.New()
	sha256Hash.Write(bs)
	hash :=sha256Hash.Sum(nil)
	return hash
}
