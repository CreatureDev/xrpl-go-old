package types

import (
	"fmt"
	"math/big"
)

const (
	accountChannelsReq   string = "account_channels"
	accountCurrenciesReq        = "account_currencies"
	accountInfoReq              = "account_info"
	accountLinesReq             = "account_lines"
	accountNFTsReq              = "account_nfts"
	accountObjectsReq           = "account_objects"
	accountOffersReq            = "account_offers"
	accountTxReq                = "account_tx"
	gatewayBalancesReq          = "gateway_balances"
	norippleCheckReq            = "noripple_check"
)

type Account string

// TODO Validate account string format
func (a Account) Valid() error {
	if a == "" {
		return fmt.Errorf("Account argument is required")
	}
	return nil
}

type BaseAccountParams struct {
	Account     Account     `json:"account"`
	LedgerHash  string      `json:"ledger_hash,omitempty"`
	LedgerIndex interface{} `json:"ledger_index,omitempty"`
}

func (b BaseAccountParams) Valid() error {
	if err := b.Account.Valid(); err != nil {
		return err
	}
	if b.LedgerIndex == nil {
		return nil
	}
	switch b.LedgerIndex.(type) {
	// Possible TODO ensure non-negative number for signed ints
	case int, uint, int32, uint32, int64, uint64, big.Int:
		return nil
	case string:
		return nil
	}
	return fmt.Errorf("Ledger Index provided in invalid format")
}

// LedgerIndex parameters are expected as a string or integer type
type AccountChannelsParams struct {
	BaseAccountParams
	DestinationAccount string      `json:"destination_account,omitempty"`
	Limit              int         `json:"limit,omitempty"`
	Marker             interface{} `json:"marker,omitempty"`
}

func (*AccountChannelsParams) MethodString() string {
	return accountChannelsReq
}

func (*AccountChannelsParams) ResponseContainer() XRPLResponse {
	return &AccountChannelsResponse{}
}

type AccountChannelsResponse struct {
	Account     string      `json:"account"`
	Channels    []Channel   `json:"channels"`
	LedgerHash  string      `json:"ledger_hash,omitempty"`
	LedgerIndex uint64      `json:"ledger_index"`
	Validated   bool        `json:"validated,omitempty"`
	Limit       int         `json:"limit,omitempty"`
	Marker      interface{} `json:"marker,omitempty"`
}

type Channel struct {
	Account            string `json:"account"`
	Amount             string `json:"amount"`
	Balance            string `json:"balance"`
	ChannelID          string `json:"channel_id"`
	DestinationAccount string `json:"destination_account"`
	SettleDelay        uint64 `json:"settle_delay"`
	PublicKey          string `json:"public_key,omitempty"`
	PublicKeyHex       string `json:"public_key_hex,omitempty"`
	Expiration         uint64 `json:"expiration,omitempty"`
	CancelAfter        uint64 `json:"cancel_after,omitempty"`
	SourceTag          uint64 `json:"source_tag,omitempty"`
	DestinationTag     uint64 `json:"destination_tag,omitempty"`
}

type AccountCurrenciesParams struct {
	BaseAccountParams
	Strict bool `json:"strict,omitempty"`
}

func (*AccountCurrenciesParams) MethodString() string {
	return accountCurrenciesReq
}

func (*AccountCurrenciesParams) ResponseContainer() XRPLResponse {
	return &AccountCurrenciesResponse{}
}

type AccountCurrenciesResponse struct {
	LedgerHash        string   `json:"ledger_hash,omitempty"`
	LedgerIndex       uint64   `json:"ledger_index"`
	RecieveCurrencies []string `json:"recieve_currencies"`
	SendCurrencies    []string `json:"send_currencies"`
	Validated         bool     `json:"validated"`
}

type AccountInfoParams struct {
	BaseAccountParams
	Queue      bool `json:"queue,omitempty"`
	SignerList bool `json:"signer_list,omitempty"`
	Strict     bool `json:"strict,omitempty"`
}

func (*AccountInfoParams) MethodString() string {
	return accountInfoReq
}

func (*AccountInfoParams) ResponseContainer() XRPLResponse {
	return &AccountInfoResponse{}
}

type AccountInfoResponse struct {
	AccData            AccountData  `json:"account_data"`
	SignerLists        []SignerList `json:"signer_lists,omitempty"`
	LedgerCurrentIndex uint64       `json:"ledger_current_index"`
	QData              QueueData    `json:"queue_data"`
	Validated          bool         `json:"validated"`
}

type AccountData struct {
	Account           string `json:"Account"`
	Balance           string `json:"Balance"`
	Flags             uint64 `json:"Flags"`
	LedgerEntryType   string `json:"LedgerEntryType"`
	OwnerCount        uint64 `json:"OwnerCount"`
	PreviousTxnID     string `json:"PreviousTxnID"`
	PreviousTxnLgrSeq uint64 `json:"PreviousTxnLgrSeq"`
	Sequence          uint64 `json:"Sequence"`
	Index             string `json:"index"`
}

type QueueData struct {
	TxnCount           uint64              `json:"txn_count"`
	AuthChangeQueued   bool                `json:"auth_change_queued"`
	LowestSequence     uint64              `json:"lowest_sequence"`
	HighestSequence    uint64              `json:"highest_sequence"`
	MaxSpendDropsTotal string              `json:"max_spend_drops_total"`
	Transactions       []QueuedTransaction `json:"transactions"`
}

type QueuedTransaction struct {
	AuthChange    bool   `json:"auth_change"`
	Fee           string `json:"fee"`
	FeeLevel      string `json:"fee_level"`
	MaxSpendDrops string `json:"max_spend_drops"`
	Seq           uint64 `json:"seq"`
}
type AccountLinesParams struct {
	BaseAccountParams
	Peer   string `json:"peer,omitempty"`
	Limit  uint64 `json:"limit,omitempty"`
	Marker interface{}
}

func (*AccountLinesParams) MethodString() string {
	return accountLinesReq
}

func (*AccountLinesParams) ResponseContainer() XRPLResponse {
	return &AccountLinesResponse{}
}

type AccountLinesResponse struct {
	Account            string        `json:"account"`
	Lines              []AccountLine `json:"lines"`
	LedgerCurrentIndex uint64        `json:"ledger_current_index"`
	LedgerIndex        uint64        `json:"ledger_index"`
	LedgerHash         string        `json:"ledger_hash"`
	Marker             interface{}   `json:"marker"`
}

type AccountLine struct {
	Account        string `json:"account"`
	Balance        string `json:"balance"`
	Currency       string `json:"currency"`
	Limit          string `json:"limit"`
	LimitPeer      string `json:"limit_peer"`
	QualityIn      uint64 `json:"quality_in"`
	QualityOut     uint64 `json:"quality_out"`
	NoRipple       bool   `json:"no_ripple"`
	NoRipplePeer   bool   `json:"no_ripple_peer"`
	Authorized     bool   `json:"authorized"`
	PeerAuthorized bool   `json:"peer_authorized"`
	Freeze         bool   `json:"freeze"`
	FreezePeer     bool   `json:"freeze_peer"`
}

type AccountNFTsParams struct {
	BaseAccountParams
	Limit  int         `json:"limit,omitempty"`
	Marker interface{} `json:"marker,omitempty"`
}

func (*AccountNFTsParams) MethodString() string {
	return accountNFTsReq
}

func (*AccountNFTsParams) ResponseContainer() XRPLResponse {
	return &AccountNFTsResponse{}
}

type AccountNFTsResponse struct {
	Account            string      `json:"account"`
	AccountNFTs        []NFT       `json:"account_nfts"`
	LedgerHash         string      `json:"ledger_hash"`
	LedgerIndex        uint64      `json:"ledger_index"`
	LedgerCurrentIndex uint64      `json:"ledger_current_index"`
	Marker             interface{} `json:"marker"`
	Validated          bool        `json:"validated"`
}

type AccountObjectsParams struct {
	BaseAccountParams
	Type                  string      `json:"type,omitempty"`
	DelectionBlockersOnly bool        `json:"deletion_blockers_only,omitempty"`
	Limit                 int         `json:"limit,omitempty"`
	Marker                interface{} `json:"Marker"`
}

func (*AccountObjectsParams) MethodString() string {
	return accountObjectsReq
}

func (*AccountObjectsParams) ResponseContainer() XRPLResponse {
	return &AccountObjectsResponse{}
}

type AccountObjectsResponse struct {
	Account            string        `json:"account"`
	AccountObjects     []interface{} `json:"account_objects"`
	LedgerHash         string        `json:"ledger_hash"`
	LedgerIndex        uint64        `json:"ledger_index"`
	LedgerCurrentIndex uint64        `json:"ledger_current_index"`
	Limit              int           `json:"limit"`
	Marker             interface{}   `json:"marker"`
	Validated          bool          `json:"validated"`
}

type AccountOffersParams struct {
	BaseAccountParams
	Limit  int         `json:"limit,omitempty"`
	Marker interface{} `json:"marker,omitempty"`
	Strict bool        `json:"strict,omitempty"`
}

func (*AccountOffersParams) MethodString() string {
	return accountOffersReq
}

func (*AccountOffersParams) ResponseContainer() XRPLResponse {
	return &AccountOffersResponse{}
}

type AccountOffersResponse struct {
	Account            string         `json:"account"`
	Offers             []AccountOffer `json:"offers"`
	LedgerCurrentIndex uint64         `json:"ledger_current_index"`
	LedgerIndex        uint64         `json:"ledger_index"`
	LedgerHash         string         `json:"ledger_hash"`
	Marker             interface{}    `json:"marker"`
}

type AccountOffer struct {
	Flags      OfferFlags  `json:"flags"`
	Seq        uint64      `json:"seq"`
	TakerGets  interface{} `json:"taker_gets"`
	TakerPays  interface{} `json:"taker_pays"`
	Quality    string      `json:"quality"`
	Expiration uint64      `json:"expiration"`
}

type AccountTxParams struct {
	BaseAccountParams
	LedgerIndexMin int64       `json:"ledger_index_min,omitempty"`
	LedgerIndexMax int64       `json:"ledger_index_max,omitempty"`
	Binary         bool        `json:"binary,omitempty"`
	Forward        bool        `json:"forward,omitempty"`
	Limit          int         `json:"limit,omitempty"`
	Marker         interface{} `json:"marker,omitempty"`
}

func (*AccountTxParams) MethodString() string {
	return accountTxReq
}

func (*AccountTxParams) ResponseContainer() XRPLResponse {
	return &AccountTxResponse{}
}

type AccountTxResponse struct {
	Account        string               `json:"account"`
	LedgerIndexMin uint64               `json:"ledger_index_min"`
	LedgerIndexMax uint64               `json:"ledger_index_max"`
	Limit          int                  `json:"limit"`
	Marker         interface{}          `json:"marker"`
	Transactions   []AccountTransaction `json:"transactions"`
	Validated      bool                 `json:"validated"`
}

type AccountTransaction struct {
	LedgerIndex uint64          `json:"ledger_index"`
	Meta        TransactionMeta `json:"meta"`
	Tx          Tx              `json:"tx"`
	TxBlob      string          `json:"tx_blob"`
	Validated   bool            `json:"validated"`
}

type GatewayBalancesParams struct {
	BaseAccountParams
	Struct    bool     `json:"strict,omitempty"`
	HotWallet []string `json:"hotwallet,omitempty"`
}

func (*GatewayBalancesParams) MethodString() string {
	return gatewayBalancesReq
}

func (*GatewayBalancesParams) ResponseContainer() XRPLResponse {
	return &GatewayBalancesResponse{}
}

type GatewayBalancesResponse struct {
	Account     string                 `json:"account"`
	Obligations map[string]interface{} `json:"obligations"`
	// TODO change fron interface{} to Currency{"currency", "value"}
	Balances           map[string][]interface{} `json:"balances"`
	Assets             map[string][]interface{} `json:"assets"`
	LedgerHash         string                   `json:"ledger_hash"`
	LedgerIndex        uint64                   `json:"ledger_index"`
	LedgerCurrentIndex uint64                   `json:"ledger_current_index"`
}

type NorippleCheckParams struct {
	BaseAccountParams
	Role         string `json:"role"`
	Transactiosn bool   `json:"transactions,omitempty"`
	Limit        int    `json:"limit,omitempty"`
}

func (*NorippleCheckParams) MethodString() string {
	return norippleCheckReq
}

func (*NorippleCheckParams) ResponseContainer() XRPLResponse {
	return &NorippleCheckResponse{}
}

type NorippleCheckResponse struct {
	LedgerCurrentIndex uint64   `json:"ledger_current_index"`
	Problems           []string `json:"problems"`
	Transactions       []Tx     `json:"transactions"`
}
