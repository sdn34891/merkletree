package merkletree

import (
	"merkletree/hash"
	"merkletree/kvstore"
)

type SampleMerkleTree struct {
	db kvstore.KVStore
}

func (sample *SampleMerkleTree) Root() []byte {
	return nil
}

func (sample *SampleMerkleTree) NewNode(content []byte) {
	key := hash.Sha3Sum256(content)
	sample.db.Put(key, content)
}

func (sample *SampleMerkleTree) Exist(content []byte) bool {
	key := hash.Sha3Sum256(content)
	return false
}

func (sample *SampleMerkleTree) DeleteNode(content []byte) {
	key := hash.Sha3Sum256(content)
	sample.db.Delete(key)
}

func (sample *SampleMerkleTree) UpdateNode(oldContent, newContent []byte) {
	oldKey := hash.Sha3Sum256(oldContent)
	newKey := hash.Sha3Sum256(newContent)
	sample.db.Delete(oldKey)
	sample.db.Put(newKey, newContent)
}

func (sample *SampleMerkleTree) GetProof(content []byte) [][]byte {
	return nil
}

var _ MerkleTree = (*SampleMerkleTree)(nil)
