/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type TxHash struct {
	Key
	Attributes TxHashAttributes `json:"attributes"`
}
type TxHashResponse struct {
	Data     TxHash   `json:"data"`
	Included Included `json:"included"`
}

type TxHashListResponse struct {
	Data     []TxHash `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustTxHash - returns TxHash from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustTxHash(key Key) *TxHash {
	var txHash TxHash
	if c.tryFindEntry(key, &txHash) {
		return &txHash
	}
	return nil
}
