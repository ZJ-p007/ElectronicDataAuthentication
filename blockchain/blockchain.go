package blockchain

import (
	"errors"
	"github.com/bolt"
	"math/big"
)

const BLOCKCHAIN = "blockchain.db"
const BUCKET_NAME = "blocks"
const LAST_HASH = "lasthash"

//全局的chain对象
var CHAIN *BlockChain

//区块链结构体得定义，代表的是一条区块链
type BlockChain struct {
	LastHash []byte   //表示区块链中最新区块的哈希，用于查找最新的区块的内容
	BoltDB   *bolt.DB //区块链中操作区块数据文件的数据库操作对象
}

/**功能
①：将新区块数据与已有区块进行连接
②：查询某个区块的数据和信息
③: 遍历区块信息
*/

func NewBlockChain() *BlockChain {
	var bc *BlockChain
	//1、先打卡文件
	db, err := bolt.Open(BLOCKCHAIN, 0600, nil)

	//2、查看chain.db文件
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME)) //假设有桶
		if bucket == nil {                       //没有桶，要创建新桶
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		//
		lastHash := bucket.Get([]byte(LAST_HASH))
		if len(lastHash) == 0 { //桶中没有lasthash记录,需要新建创世区块，并保存
			//创世区块
			genesis := CreateGenesisBlock()
			//区块序列化以后的数据
			gensisBytes := genesis.Serialize()
			//创世区块保存到boltdb中
			bucket.Put(genesis.Hash, gensisBytes)
			//更新指向最新块钱的lasthash的值
			bucket.Put([]byte(LAST_HASH), genesis.Hash)
			bc = &BlockChain{
				LastHash: genesis.Hash,
				BoltDB:   db,
			}
		} else { //桶中已有lasthash的记录，不再需要创世区块，只需要读取即可
			lasthash1 := bucket.Get([]byte(LAST_HASH))
			bc = &BlockChain{
				LastHash: lasthash1,
				BoltDB:   db,
			}
		}
		return nil
	})
	CHAIN = bc
	return bc
}

//该方法用于遍历区块链chain.db,并将所有的区块查出，返回
func (bc BlockChain) QueryAllBlocks() ([]*Block, error) {
	blocks := make([]*Block, 0) //blocks是一个容器用于盛放查询到的区块
	db := bc.BoltDB
	var err error
	//从chain.db查询所有的区块
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			err = errors.New("查询区块链数据失败")
			return err
		}
		//bucket存在,获取信息
		eachHash := bc.LastHash
		eachBig := new(big.Int)
		zeroBig := big.NewInt(0) //默认值
		for {
			//根据区块的哈希值获取对应的区块
			eachBlockBytes := bucket.Get(eachHash)
			//反序列化
			eachBlock, _ := DSerialize(eachBlockBytes)
			//将遍历的区块放到容器
			blocks = append(blocks, eachBlock)

			eachBig.SetBytes(eachBlock.PrevHash)
			if eachBig.Cmp(zeroBig) == 0 {
				//找到创世区块
				break //找到创世区块跳出循环
			}
			//不满足条件
			eachHash = eachBlock.PrevHash
		}
		return nil
	})
	return blocks, err
}

//该方法用于完成根据用户输入的区块高度查询对应的区块信息
func (bc BlockChain) QueryBlockRyHeigth(height int64) (*Block, error) {
	if height < 0 {
		return nil, nil
	}
	db := bc.BoltDB
	var errs error
	var eachBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			errs = errors.New("读取区块数据失败")
			return errs
		}
		eachHash := bc.LastHash
		/*if eachBlock.Height < height {
		          break
				}

				eachHash := bc.LastHash*/
		for {
			//获取到最后一个区块的哈希值
			//lashBlockHash :=bucket.Get(bc.LastHash)
			eachBlockHash := bucket.Get(eachHash)
			//最后一个区块的byte类型
			eachBlockHashBytes := bucket.Get(eachBlockHash)
			//反序列
			eachBlock, errs := DSerialize(eachBlockHashBytes)
			if errs != nil {
				//errs = err
				return errs
			}
			if eachBlock.Height < height {
				break
			}

			if eachBlock.Height == height { //t跳出
				break
			}
			//如果高度不匹配，则不满足要求
			eachHash = eachBlock.PrevHash
		}
		return nil
	})
	return eachBlock, errs
}

//保存数据到区块链中：先生成一个新区块，然后将新区快添加到区块链中
func (bc BlockChain) AddData(data []byte) (Block, error) {
	//1.从文件当中读取到最新的区块
	db := bc.BoltDB
	var lastBlock *Block
	//error的自定义
	var err error
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			//panic("读取区块链数据失败")
			err = errors.New("读取区块链数据失败")
			return err
		}
		//lasthash := bucket.Get([]byte(LAST_HASH))
		lastBlockBytes := bucket.Get(bc.LastHash)
		//反序列化
		lastBlock, _ = DSerialize(lastBlockBytes)
		return nil
	})

	//先新建一个区块
	newBlock := NewBlock(lastBlock.Height+1, lastBlock.Hash, data)

	//把新区块存储到文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))

		//把新创建的区块存入到boltdb数据库中
		bucket.Put(newBlock.Hash, newBlock.Serialize())

		//更新LASTHASH对应的值，更新为最新存储的区块的hash值
		bucket.Put([]byte(LAST_HASH), newBlock.Hash)

		//将区块链实例的lasthash值
		bc.LastHash = newBlock.Hash
		return nil
	})

	//1.返回值语句包含newBlock,err,其中err包含信息
	return newBlock, err
}
