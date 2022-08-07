package types

import "encoding/json"

const (
	submitTxReq string = "submit"
	txReq              = "tx"
)

type BaseTx struct {
	Account            Account       `json:"Account"`
	TransactionType    string        `json:"TransactionType"`
	Fee                string        `json:"Fee"`
	Sequence           uint64        `json:"Sequence"`
	AccountTxnID       string        `json:"AccountTxnID"`
	Flags              TxFlags       `json:"Flags"`
	LastLedgerSequence uint64        `json:"LastLedgerSequence"`
	Memos              []MemoWrapper `json:"Memos"`
	// TODO verify this field is not wrapped like Memo field
	Signers []SignerEntry `json:"Signers"`
}

func (*BaseTx) IsTx() {
}

type Tx interface {
	IsTx()
}

type SubmitTxParams struct {
	TxBlob   string `json:"tx_blob"`
	FailHard bool   `json:"fail_hard,omitempty"`
}

func (*SubmitTxParams) Validate() error {
	return nil
}

func (*SubmitTxParams) MethodString() string {
	return submitTxReq
}

func (*SubmitTxParams) DecodeResponse(res json.RawMessage) XRPLResponse {
	// TODO determine type of transaction and return appropriate data
	return &SubmitTxResponse{}
}

type SubmitTxResponse struct {
	EngineResult             string `json:"engine_result"`
	EngineResultCode         int    `json:"engine_result_code"`
	EngineResultMessage      string `json:"engine_result_message"`
	TxBlob                   string `json:"tx_blob"`
	TxJson                   Tx     `json:"tx_json"`
	Accepted                 bool   `json:"accepted"`
	AccountSequenceAvailable uint64 `json:"account_sequence_available"`
	AccountSequenceNext      uint64 `json:"account_sequence_next"`
	Applied                  bool   `json:"applied"`
	Broadcast                bool   `json:"broadcast"`
	Kept                     bool   `json:"kept"`
	OpenLedgerCost           string `json:"open_ledger_cost"`
	ValidatedLedgerIndex     uint64 `json:"validated_ledger_index"`
}

type TxParams struct {
	Transaction string `json:"transaction"`
	Binary      bool   `json:"binary,omitempty"`
	MinLedger   uint64 `json:"min_ledger,omitempty"`
	MaxLedger   uint64 `json:"max_ledger,omitempty"`
}

func (*TxParams) Validate() error {
	return nil
}

func (*TxParams) MethodString() string {
	return txReq
}

func (t *TxParams) DecodeResponse(res json.RawMessage) XRPLResponse {
	if t.Binary {
		ret := &TxBinaryResponse{}
		json.Unmarshal(res, ret)
		return ret
	}
	ret := &TxResponse{}
	json.Unmarshal(res, ret)
	return ret
}

type TxResponse struct {
	Tx
	Hash        string          `json:"hash"`
	LedgerIndex uint64          `json:"ledger_index"`
	Meta        TransactionMeta `json:"meta"`
	Validated   bool            `json:"validated"`
}

type TxBinaryResponse struct {
	Tx
	Hash        string `json:"hash"`
	LedgerIndex uint64 `json:"ledger_index"`
	Meta        string `json:"meta"`
	Validated   bool   `json:"validated"`
}