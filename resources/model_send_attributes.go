/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type SendAttributes struct {
	Recipient SendAttributesRecipient `json:"recipient"`
}

type SendAttributesRecipient struct {
	Address string  `json:"address"`
	Amount  float32 `json:"amount"`
}
