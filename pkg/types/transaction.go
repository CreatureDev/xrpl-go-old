package types

import "encoding/json"

const (
	submitTxReq          string = "submit"
	submitMultisignedReq string = "submit_multisigned"
	txReq                       = "tx"
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

func ParseTx(msg json.RawMessage) Tx {
	base := &BaseTx{}
	json.Unmarshal(msg, base)
	switch base.TransactionType {
	// TODO

	}
	return base
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
	ret := &SubmitTxResponse{}
	json.Unmarshal(res, ret)
	ret.Tx = ParseTx(ret.TxJson)
	return &ret
}

type SubmitTxResponse struct {
	EngineResult             string          `json:"engine_result"`
	EngineResultCode         int             `json:"engine_result_code"`
	EngineResultMessage      string          `json:"engine_result_message"`
	TxBlob                   string          `json:"tx_blob"`
	TxJson                   json.RawMessage `json:"tx_json"`
	Tx                       Tx
	Accepted                 bool   `json:"accepted"`
	AccountSequenceAvailable uint64 `json:"account_sequence_available"`
	AccountSequenceNext      uint64 `json:"account_sequence_next"`
	Applied                  bool   `json:"applied"`
	Broadcast                bool   `json:"broadcast"`
	Kept                     bool   `json:"kept"`
	OpenLedgerCost           string `json:"open_ledger_cost"`
	ValidatedLedgerIndex     uint64 `json:"validated_ledger_index"`
}

type SubmitMultisignedParams struct {
	Tx       Tx   `json:"tx_json"`
	FailHard bool `json:"fail_hard,omitempty"`
}

func (*SubmitMultisignedParams) Validate() error {
	return nil
}

func (*SubmitMultisignedParams) MethodString() string {
	return submitMultisignedReq
}

func (*SubmitMultisignedParams) DecodeResponse(res json.RawMessage) XRPLResponse {
	ret := &SubmitMultisignedResponse{}
	json.Unmarshal(res, ret)
	ret.Tx = ParseTx(ret.TxJson)
	return ret
}

type SubmitMultisignedResponse struct {
	EngineResult        string          `json:"engine_result"`
	EngineResultCode    int             `json:"engine_result_code"`
	EngineResultMessage string          `json:"engine_result_message"`
	TxBlob              string          `json:"tx_blob"`
	TxJson              json.RawMessage `json:"tx_json"`
	Tx                  Tx
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
	tx := ParseTx(res)
	if t.Binary {
		ret := &TxBinaryResponse{}
		json.Unmarshal(res, ret)
		ret.Tx = tx
		return ret
	}
	ret := &TxResponse{}
	json.Unmarshal(res, ret)
	ret.Tx = tx
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
