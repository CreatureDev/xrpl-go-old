package types

type AccountRoot struct {
}

type Amendments struct {
}

type Check struct {
}

type DepositPreauth struct {
}

type DirectoryNode struct {
}

type Escrow struct {
}

type FeeSettings struct {
}

type LedgerHashes struct {
}

type NegativeUNL struct {
}

type NFTokenOffer struct {
}

type NFTokenPage struct {
}

type OfferFlags uint

const (
	OfferPassive OfferFlags = 0x00010000
	OfferSell               = 0x00020000
)

type Offer struct {
}

type PayChannel struct {
}

type RippleState struct {
}

type SignerListFlags uint

const (
	SLOneOwnerCount SignerListFlags = 0x00010000
)

type SignerList struct {
	LedgerEntryType   string          `json:"LedgerEntryType"`
	Flags             SignerListFlags `json:"Flags"`
	PreviousTxnID     string          `json:"PreviousTxnID"`
	PreviousTxnLgrSeq uint64          `json:"PreviousTxnLgrSeq"`
	OwnerNode         string          `json:"OwnerNode"`
	SignerEntries     []SignerEntry   `json:"SignerEntries"`
	SignerListId      uint64          `json:"SignerListID"`
	SignerQuorum      uint64          `json:"SignerQuorum"`
}

type SignerEntry struct {
	Account       string `json:"Account"`
	SignerWeight  uint64 `json:"SignerWeight"`
	WalletLocator string `json:"WalletLocator"`
}

type Ticket struct {
}

type NFTokenFlag uint

const (
	NftBurnable     NFTokenFlag = 0x1
	NftOnlyXRP                  = 0x2
	NftTrustLine                = 0x4
	NftTransferable             = 0x8
	NftReservedFlag             = 0x8000
)

type NFT struct {
	Flags        NFTokenFlag `json:"Flags"`
	Issuer       string      `json:"Issuer"`
	NFTokenID    string      `json:"NFTokenID"`
	NFTokenTaxon uint        `json:"NFTokenTaxon"`
	URI          string      `json:"URI"`
	NFTSerial    uint64      `json:"nft_serial"`
}

type TransactionMeta struct {
	AffectedNodes          []AffectedNode `json:"AffectedNodes"`
	PartialDeliveredAmount interface{}    `json:"DeliveredAmount"`
	TransactionIndex       uint64         `json:"TransactionIndex"`
	TransactionResult      string         `json:"TransactionResult"`
	DeliveredAmount        interface{}    `json:"delivered_amount"`
}

type AffectedNode struct {
	CreatedNode  *CreatedNode  `json:"CreatedNode,omitempty"`
	ModifiedNode *ModifiedNode `json:"ModifiedNode,omitempty"`
	DeletedNode  *DeletedNode  `json:"DeletedNode,omitempty"`
}

type CreatedNode struct {
	LedgerEntryType string      `json:"LedgerEntryType,omitempty"`
	LedgerIndex     string      `json:"LedgerIndex,omitempty"`
	NewFields       interface{} `json:"NewFields,omitempty"`
}

type ModifiedNode struct {
	LedgerEntryType   string      `json:"LedgerEntryType"`
	LedgerIndex       string      `json:"LedgerIndex"`
	FinalFields       interface{} `json:"FinalFields"`
	PreviousFields    interface{} `json:"PreviousFields"`
	PreviousTxnID     string      `json:"PreviousTxnID,omitempty"`
	PreviousTxnLgrSeq uint64      `json:"PreviousTxnLgrSeq,omitempty"`
}

type DeletedNode struct {
	LedgerEntryType string      `json:"LedgerEntryType"`
	LedgerIndex     string      `json:"LedgerIndex"`
	FinalFields     interface{} `json:"FinalFields"`
}

type TxFlags uint32

type Memo struct {
	MemoData   string `json:"MemoData"`
	MemoFormat string `json:"MemoFormat"`
	MemoType   string `json:"MemoType"`
}

type MemoWrapper struct {
	Memo Memo `json:"Memo"`
}
