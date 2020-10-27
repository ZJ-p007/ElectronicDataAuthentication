package blockchain

import "github.com/bolt"

const BLOCKCHAIN = "blockchain.db"
const BUCKET_NAME = "blocks"
const LAST_HASH  = "last_hash"

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

//创建一条区块链
func NewBlockChain() BlockChain {
	//创世区块
	genesis := CreateGenesisBlock()
	db, err := bolt.Open("BLOCKCHAIN", 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	bc := BlockChain{
		LastHash: genesis.Hash,
		BoltDB:   db,
	}

	//把创世区块保存到数据库文件当中
	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte("BUCKET_NAME"))
		if err != nil {
			panic(err.Error())
		}

		//序列化
		genesisBytes := genesis.Serialize()
		//把创世区块存储到桶中
		bucket.Put(genesis.Hash, genesisBytes)
		//更新最新区块的哈希值记录
		bucket.Put([]byte(LAST_HASH),genesis.Hash)
		return nil
	})

	return bc
}

//保存数据到区块链中：先生成一个新区块，然后将新区快添加到区块链中
func (bc BlockChain) AddData(data []byte) {
	//1.从文件当中读取到最新的区块
	db := bc.BoltDB
	var lastBlock *Block
	db.View(func(tx *bolt.Tx) error {
       bucket := tx.Bucket([]byte(BUCKET_NAME))
       if bucket == nil{
       	 panic("读取区块链数据失败")
	   }
	   lasthash := bucket.Get([]byte(LAST_HASH))
	   lastBlockBytes := bucket.Get(lasthash)
	   //反序列化

	   lastBlock,_ = DSerialize(lastBlockBytes)
		return nil
	})

	//先新建一个区块
	//newBlock := NewBlock(lastBlock.Height+1,lastBlock.Hash, data)
    //把新区块存储到文件中
   // db.Update()
}
