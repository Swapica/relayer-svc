/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type EvmTransactionAttributes struct {
	ChainId string `json:"chain_id"`
	// Transaction call data
	Data string `json:"data"`
	// The address of the sender
	From string `json:"from"`
	// The address of the transaction recipient
	To string `json:"to"`
	// The amount of Wei to send
	Value string `json:"value"`
}
