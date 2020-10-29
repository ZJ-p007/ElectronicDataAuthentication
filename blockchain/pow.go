package blockchain

import (
	"DataCertPlatform/utils"
	"bytes"
	"math/big"
)

const DIFFICULTY  = 14

//工作量证明算法结构体
type ProofOfWork struct {
	Target *big.Int //系统的目标值
	Block  Block    //要找的nonce值对应的区块
}

//实例化一个pow算法
func NewPow(block Block) ProofOfWork {
	t := big.NewInt(1)
	t = t.Lsh(t, 255 - DIFFICULTY)
	pow := ProofOfWork{
		Target: t,
		Block:  block,
	}
	return pow
}

//run方法用于寻找合适的nonce值
func (p ProofOfWork) Run() ([]byte,int64) {

	var nonce int64
	nonce = 0
    var blockHash []byte

	for { //使用死循环
		block := p.Block
		heightBytes, _ := utils.Int64ToByte(block.Height)
		timeBytes, _ := utils.Int64ToByte(block.TimeStamp)
		versionBytes := utils.StringToBytes(block.Version)

		nonceBytes, _ := utils.Int64ToByte(nonce)
		//已有的区块信息和尝试的nonce值的拼接
		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.PrevHash,
			block.Data,
			versionBytes,
			nonceBytes,
		}, []byte{})

		//区块和尝试的nonce值拼接后得到的hash值
		blockHash = utils.SHA256HashBlock(blockBytes)
		target := p.Target     //目标值
		var hashBig *big.Int   //声明和定义
		hashBig = new(big.Int) //分配内存空间，为变量分配地址

		hashBig = hashBig.SetBytes(blockHash)
		//hashBig = hashBig.SetBytes(blockHash)
		//当前nonce值
		//fmt.Println("当前尝试的Nonce值",nonce)
		if hashBig.Cmp(target) == -1 {
			//停止寻找
			break
		}
		nonce++ //自增继续寻找
	}
	//将找到的符合规则的nonce返回
	return blockHash,nonce
}


