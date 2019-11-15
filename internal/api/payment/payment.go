package payment

import (
	"context"
	"gitlab.com/MXCFoundation/cloud/conn_checker/pkg/api/payment"
)

type PaymentAPI struct{}

// PaymentAPI returns a new M2MServerAPI.
func NewPaymentServerAPI() *PaymentAPI {
	return &PaymentAPI{}
}

func (*PaymentAPI) TokenTxReq (ctx context.Context, req *payment.TxReqType) (*payment.TxReqReplyType, error) {
	req.PaymentClientEnum = 1
	req.ReqIdClient = 20
	req.ReceiverAdr = ""
	req.Amount = "0.000001"
	req.TokenNameEnum = 0

	return &payment.TxReqReplyType{}, nil
}

func (*PaymentAPI) CheckTxStatus (ctx context.Context, req *payment.CheckTxStatusType) (*payment.CheckTxStatusReplyType, error)  {

	return &payment.CheckTxStatusReplyType{}, nil
}