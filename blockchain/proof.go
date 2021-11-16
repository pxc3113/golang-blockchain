package blockchain

// take the date from the block
//creates counter (nonce) which starts at -
//create a has of the data plus the counter
//check the hash to see if it meets a set of requirementss
//requirements:
//the first few bytes must contain 0s

const Difficulty = 12

type ProofOfWork struct {
	Block *Block
	Target *big.Int
}
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte{
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int65(nonce)),
			TOhex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte){
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce <math.maxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x",hash)
		intHash.SetBytes(hash[:])
		if intHash.Cmp(pow.Target)==-1{
			break
		}else{
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

func (pow *ProofOfWork) validate() bool{
	var intHash big.Int
	data := pow.InitData(pow.block>nonce)
	hash:= sha256.SUm256(data)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target)==-1
}

func ToHex(num int64) []byute{
	butff:= new(bytes.Buffer)
	err:=binary.Write(buff, binary.BIgEndina, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}