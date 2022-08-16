// Copyright 2021 Tencent Inc. All rights reserved.
//
// 境内普通商户退款API
//
// 境内普通商户退款功能涉及的API文档
//
// API version: 1.1.1

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package refunddomestic

import (
	"encoding/json"
	"fmt"
	"time"
)

// Account * `AVAILABLE` - 可用余额, 多账户资金准备退款可用余额出资账户类型 * `UNAVAILABLE` - 不可用余额, 多账户资金准备退款不可用余额出资账户类型
type Account string

func (e Account) Ptr() *Account {
	return &e
}

// Enums of Account
const (
	ACCOUNT_AVAILABLE   Account = "AVAILABLE"
	ACCOUNT_UNAVAILABLE Account = "UNAVAILABLE"
)

// Amount
type Amount struct {
	// 订单总金额，单位为分
	Total *int64 `json:"total"`
	// 退款标价金额，单位为分，可以做部分退款
	Refund *int64 `json:"refund"`
	// 退款出资的账户类型及金额信息
	From []FundsFromItem `json:"from,omitempty"`
	// 现金支付金额，单位为分，只能为整数
	PayerTotal *int64 `json:"payer_total"`
	// 退款给用户的金额，不包含所有优惠券金额
	PayerRefund *int64 `json:"payer_refund"`
	// 去掉非充值代金券退款金额后的退款金额，单位为分，退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	SettlementRefund *int64 `json:"settlement_refund"`
	// 应结订单金额=订单金额-免充值代金券金额，应结订单金额<=订单金额，单位为分
	SettlementTotal *int64 `json:"settlement_total"`
	// 优惠退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为现金，说明详见代金券或立减优惠，单位为分
	DiscountRefund *int64 `json:"discount_refund"`
	// 符合ISO 4217标准的三位字母代码，目前只支持人民币：CNY。
	Currency *string `json:"currency"`
}

func (o Amount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Total == nil {
		return nil, fmt.Errorf("field `Total` is required and must be specified in Amount")
	}
	toSerialize["total"] = o.Total

	if o.Refund == nil {
		return nil, fmt.Errorf("field `Refund` is required and must be specified in Amount")
	}
	toSerialize["refund"] = o.Refund

	if o.From != nil {
		toSerialize["from"] = o.From
	}

	if o.PayerTotal == nil {
		return nil, fmt.Errorf("field `PayerTotal` is required and must be specified in Amount")
	}
	toSerialize["payer_total"] = o.PayerTotal

	if o.PayerRefund == nil {
		return nil, fmt.Errorf("field `PayerRefund` is required and must be specified in Amount")
	}
	toSerialize["payer_refund"] = o.PayerRefund

	if o.SettlementRefund == nil {
		return nil, fmt.Errorf("field `SettlementRefund` is required and must be specified in Amount")
	}
	toSerialize["settlement_refund"] = o.SettlementRefund

	if o.SettlementTotal == nil {
		return nil, fmt.Errorf("field `SettlementTotal` is required and must be specified in Amount")
	}
	toSerialize["settlement_total"] = o.SettlementTotal

	if o.DiscountRefund == nil {
		return nil, fmt.Errorf("field `DiscountRefund` is required and must be specified in Amount")
	}
	toSerialize["discount_refund"] = o.DiscountRefund

	if o.Currency == nil {
		return nil, fmt.Errorf("field `Currency` is required and must be specified in Amount")
	}
	toSerialize["currency"] = o.Currency
	return json.Marshal(toSerialize)
}

func (o Amount) String() string {
	var ret string
	if o.Total == nil {
		ret += "Total:<nil>, "
	} else {
		ret += fmt.Sprintf("Total:%v, ", *o.Total)
	}

	if o.Refund == nil {
		ret += "Refund:<nil>, "
	} else {
		ret += fmt.Sprintf("Refund:%v, ", *o.Refund)
	}

	ret += fmt.Sprintf("From:%v, ", o.From)

	if o.PayerTotal == nil {
		ret += "PayerTotal:<nil>, "
	} else {
		ret += fmt.Sprintf("PayerTotal:%v, ", *o.PayerTotal)
	}

	if o.PayerRefund == nil {
		ret += "PayerRefund:<nil>, "
	} else {
		ret += fmt.Sprintf("PayerRefund:%v, ", *o.PayerRefund)
	}

	if o.SettlementRefund == nil {
		ret += "SettlementRefund:<nil>, "
	} else {
		ret += fmt.Sprintf("SettlementRefund:%v, ", *o.SettlementRefund)
	}

	if o.SettlementTotal == nil {
		ret += "SettlementTotal:<nil>, "
	} else {
		ret += fmt.Sprintf("SettlementTotal:%v, ", *o.SettlementTotal)
	}

	if o.DiscountRefund == nil {
		ret += "DiscountRefund:<nil>, "
	} else {
		ret += fmt.Sprintf("DiscountRefund:%v, ", *o.DiscountRefund)
	}

	if o.Currency == nil {
		ret += "Currency:<nil>"
	} else {
		ret += fmt.Sprintf("Currency:%v", *o.Currency)
	}

	return fmt.Sprintf("Amount{%s}", ret)
}

func (o Amount) Clone() *Amount {
	ret := Amount{}

	if o.Total != nil {
		ret.Total = new(int64)
		*ret.Total = *o.Total
	}

	if o.Refund != nil {
		ret.Refund = new(int64)
		*ret.Refund = *o.Refund
	}

	if o.From != nil {
		ret.From = make([]FundsFromItem, len(o.From))
		for i, item := range o.From {
			ret.From[i] = *item.Clone()
		}
	}

	if o.PayerTotal != nil {
		ret.PayerTotal = new(int64)
		*ret.PayerTotal = *o.PayerTotal
	}

	if o.PayerRefund != nil {
		ret.PayerRefund = new(int64)
		*ret.PayerRefund = *o.PayerRefund
	}

	if o.SettlementRefund != nil {
		ret.SettlementRefund = new(int64)
		*ret.SettlementRefund = *o.SettlementRefund
	}

	if o.SettlementTotal != nil {
		ret.SettlementTotal = new(int64)
		*ret.SettlementTotal = *o.SettlementTotal
	}

	if o.DiscountRefund != nil {
		ret.DiscountRefund = new(int64)
		*ret.DiscountRefund = *o.DiscountRefund
	}

	if o.Currency != nil {
		ret.Currency = new(string)
		*ret.Currency = *o.Currency
	}

	return &ret
}

// AmountReq
type AmountReq struct {
	// 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额。
	Refund *int64 `json:"refund"`
	// 退款需要从指定账户出资时，传递此参数指定出资金额（币种的最小单位，只能为整数）。 同时指定多个账户出资退款的使用场景需要满足以下条件：1、未开通退款支出分离产品功能；2、订单属于分账订单，且分账处于待分账或分账中状态。 参数传递需要满足条件：1、基本账户可用余额出资金额与基本账户不可用余额出资金额之和等于退款金额；2、账户类型不能重复。 上述任一条件不满足将返回错误
	From []FundsFromItem `json:"from,omitempty"`
	// 原支付交易的订单总金额，币种的最小单位，只能为整数。
	Total *int64 `json:"total"`
	// 符合ISO 4217标准的三位字母代码，目前只支持人民币：CNY。
	Currency *string `json:"currency"`
}

func (o AmountReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Refund == nil {
		return nil, fmt.Errorf("field `Refund` is required and must be specified in AmountReq")
	}
	toSerialize["refund"] = o.Refund

	if o.From != nil {
		toSerialize["from"] = o.From
	}

	if o.Total == nil {
		return nil, fmt.Errorf("field `Total` is required and must be specified in AmountReq")
	}
	toSerialize["total"] = o.Total

	if o.Currency == nil {
		return nil, fmt.Errorf("field `Currency` is required and must be specified in AmountReq")
	}
	toSerialize["currency"] = o.Currency
	return json.Marshal(toSerialize)
}

func (o AmountReq) String() string {
	var ret string
	if o.Refund == nil {
		ret += "Refund:<nil>, "
	} else {
		ret += fmt.Sprintf("Refund:%v, ", *o.Refund)
	}

	ret += fmt.Sprintf("From:%v, ", o.From)

	if o.Total == nil {
		ret += "Total:<nil>, "
	} else {
		ret += fmt.Sprintf("Total:%v, ", *o.Total)
	}

	if o.Currency == nil {
		ret += "Currency:<nil>"
	} else {
		ret += fmt.Sprintf("Currency:%v", *o.Currency)
	}

	return fmt.Sprintf("AmountReq{%s}", ret)
}

func (o AmountReq) Clone() *AmountReq {
	ret := AmountReq{}

	if o.Refund != nil {
		ret.Refund = new(int64)
		*ret.Refund = *o.Refund
	}

	if o.From != nil {
		ret.From = make([]FundsFromItem, len(o.From))
		for i, item := range o.From {
			ret.From[i] = *item.Clone()
		}
	}

	if o.Total != nil {
		ret.Total = new(int64)
		*ret.Total = *o.Total
	}

	if o.Currency != nil {
		ret.Currency = new(string)
		*ret.Currency = *o.Currency
	}

	return &ret
}

// Channel * `ORIGINAL` - 原路退款, 退款渠道 * `BALANCE` - 退回到余额, 退款渠道 * `OTHER_BALANCE` - 原账户异常退到其他余额账户, 退款渠道 * `OTHER_BANKCARD` - 原银行卡异常退到其他银行卡, 退款渠道
type Channel string

func (e Channel) Ptr() *Channel {
	return &e
}

// Enums of Channel
const (
	CHANNEL_ORIGINAL       Channel = "ORIGINAL"
	CHANNEL_BALANCE        Channel = "BALANCE"
	CHANNEL_OTHER_BALANCE  Channel = "OTHER_BALANCE"
	CHANNEL_OTHER_BANKCARD Channel = "OTHER_BANKCARD"
)

// CreateRequest
type CreateRequest struct {
	// 子商户的商户号，由微信支付生成并下发。服务商模式下必须传递此参数
	SubMchid *string `json:"sub_mchid,omitempty"`
	// 原支付交易对应的微信订单号
	TransactionId *string `json:"transaction_id,omitempty"`
	// 原支付交易对应的商户订单号
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo *string `json:"out_refund_no"`
	// 若商户传入，会在下发给用户的退款消息中体现退款原因
	Reason *string `json:"reason,omitempty"`
	// 异步接收微信支付退款结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。 如果参数中传了notify_url，则商户平台上配置的回调地址将不会生效，优先回调当前传的这个地址。
	NotifyUrl *string `json:"notify_url,omitempty"`
	// 若传递此参数则使用对应的资金账户退款，否则默认使用未结算资金退款（仅对老资金流商户适用）  枚举值： - AVAILABLE：可用余额账户    * `AVAILABLE` - 可用余额
	FundsAccount *ReqFundsAccount `json:"funds_account,omitempty"`
	// 订单金额信息
	Amount *AmountReq `json:"amount"`
	// 指定商品退款需要传此参数，其他场景无需传递
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

func (o CreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.SubMchid != nil {
		toSerialize["sub_mchid"] = o.SubMchid
	}

	if o.TransactionId != nil {
		toSerialize["transaction_id"] = o.TransactionId
	}

	if o.OutTradeNo != nil {
		toSerialize["out_trade_no"] = o.OutTradeNo
	}

	if o.OutRefundNo == nil {
		return nil, fmt.Errorf("field `OutRefundNo` is required and must be specified in CreateRequest")
	}
	toSerialize["out_refund_no"] = o.OutRefundNo

	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	if o.NotifyUrl != nil {
		toSerialize["notify_url"] = o.NotifyUrl
	}

	if o.FundsAccount != nil {
		toSerialize["funds_account"] = o.FundsAccount
	}

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in CreateRequest")
	}
	toSerialize["amount"] = o.Amount

	if o.GoodsDetail != nil {
		toSerialize["goods_detail"] = o.GoodsDetail
	}
	return json.Marshal(toSerialize)
}

func (o CreateRequest) String() string {
	var ret string
	if o.SubMchid == nil {
		ret += "SubMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}

	if o.TransactionId == nil {
		ret += "TransactionId:<nil>, "
	} else {
		ret += fmt.Sprintf("TransactionId:%v, ", *o.TransactionId)
	}

	if o.OutTradeNo == nil {
		ret += "OutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutTradeNo:%v, ", *o.OutTradeNo)
	}

	if o.OutRefundNo == nil {
		ret += "OutRefundNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutRefundNo:%v, ", *o.OutRefundNo)
	}

	if o.Reason == nil {
		ret += "Reason:<nil>, "
	} else {
		ret += fmt.Sprintf("Reason:%v, ", *o.Reason)
	}

	if o.NotifyUrl == nil {
		ret += "NotifyUrl:<nil>, "
	} else {
		ret += fmt.Sprintf("NotifyUrl:%v, ", *o.NotifyUrl)
	}

	if o.FundsAccount == nil {
		ret += "FundsAccount:<nil>, "
	} else {
		ret += fmt.Sprintf("FundsAccount:%v, ", *o.FundsAccount)
	}

	ret += fmt.Sprintf("Amount:%v, ", o.Amount)

	ret += fmt.Sprintf("GoodsDetail:%v", o.GoodsDetail)

	return fmt.Sprintf("CreateRequest{%s}", ret)
}

func (o CreateRequest) Clone() *CreateRequest {
	ret := CreateRequest{}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	if o.OutTradeNo != nil {
		ret.OutTradeNo = new(string)
		*ret.OutTradeNo = *o.OutTradeNo
	}

	if o.OutRefundNo != nil {
		ret.OutRefundNo = new(string)
		*ret.OutRefundNo = *o.OutRefundNo
	}

	if o.Reason != nil {
		ret.Reason = new(string)
		*ret.Reason = *o.Reason
	}

	if o.NotifyUrl != nil {
		ret.NotifyUrl = new(string)
		*ret.NotifyUrl = *o.NotifyUrl
	}

	if o.FundsAccount != nil {
		ret.FundsAccount = new(ReqFundsAccount)
		*ret.FundsAccount = *o.FundsAccount
	}

	if o.Amount != nil {
		ret.Amount = o.Amount.Clone()
	}

	if o.GoodsDetail != nil {
		ret.GoodsDetail = make([]GoodsDetail, len(o.GoodsDetail))
		for i, item := range o.GoodsDetail {
			ret.GoodsDetail[i] = *item.Clone()
		}
	}

	return &ret
}

// FundsAccount * `UNSETTLED` - 未结算资金, 退款所使用资金对应的资金账户类型 * `AVAILABLE` - 可用余额, 退款所使用资金对应的资金账户类型 * `UNAVAILABLE` - 不可用余额, 退款所使用资金对应的资金账户类型 * `OPERATION` - 运营户, 退款所使用资金对应的资金账户类型 * `BASIC` - 基本账户（含可用余额和不可用余额）, 退款所使用资金对应的资金账户类型
type FundsAccount string

func (e FundsAccount) Ptr() *FundsAccount {
	return &e
}

// Enums of FundsAccount
const (
	FUNDSACCOUNT_UNSETTLED   FundsAccount = "UNSETTLED"
	FUNDSACCOUNT_AVAILABLE   FundsAccount = "AVAILABLE"
	FUNDSACCOUNT_UNAVAILABLE FundsAccount = "UNAVAILABLE"
	FUNDSACCOUNT_OPERATION   FundsAccount = "OPERATION"
	FUNDSACCOUNT_BASIC       FundsAccount = "BASIC"
)

// FundsFromItem
type FundsFromItem struct {
	// 下面枚举值多选一。 枚举值： AVAILABLE : 可用余额 UNAVAILABLE : 不可用余额 * `AVAILABLE` - 可用余额 * `UNAVAILABLE` - 不可用余额
	Account *Account `json:"account"`
	// 对应账户出资金额
	Amount *int64 `json:"amount"`
}

func (o FundsFromItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Account == nil {
		return nil, fmt.Errorf("field `Account` is required and must be specified in FundsFromItem")
	}
	toSerialize["account"] = o.Account

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in FundsFromItem")
	}
	toSerialize["amount"] = o.Amount
	return json.Marshal(toSerialize)
}

func (o FundsFromItem) String() string {
	var ret string
	if o.Account == nil {
		ret += "Account:<nil>, "
	} else {
		ret += fmt.Sprintf("Account:%v, ", *o.Account)
	}

	if o.Amount == nil {
		ret += "Amount:<nil>"
	} else {
		ret += fmt.Sprintf("Amount:%v", *o.Amount)
	}

	return fmt.Sprintf("FundsFromItem{%s}", ret)
}

func (o FundsFromItem) Clone() *FundsFromItem {
	ret := FundsFromItem{}

	if o.Account != nil {
		ret.Account = new(Account)
		*ret.Account = *o.Account
	}

	if o.Amount != nil {
		ret.Amount = new(int64)
		*ret.Amount = *o.Amount
	}

	return &ret
}

// GoodsDetail
type GoodsDetail struct {
	// 由半角的大小写字母、数字、中划线、下划线中的一种或几种组成
	MerchantGoodsId *string `json:"merchant_goods_id"`
	// 微信支付定义的统一商品编号（没有可不传）
	WechatpayGoodsId *string `json:"wechatpay_goods_id,omitempty"`
	// 商品的实际名称
	GoodsName *string `json:"goods_name,omitempty"`
	// 商品单价金额，单位为分
	UnitPrice *int64 `json:"unit_price"`
	// 商品退款金额，单位为分
	RefundAmount *int64 `json:"refund_amount"`
	// 对应商品的退货数量
	RefundQuantity *int64 `json:"refund_quantity"`
}

func (o GoodsDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.MerchantGoodsId == nil {
		return nil, fmt.Errorf("field `MerchantGoodsId` is required and must be specified in GoodsDetail")
	}
	toSerialize["merchant_goods_id"] = o.MerchantGoodsId

	if o.WechatpayGoodsId != nil {
		toSerialize["wechatpay_goods_id"] = o.WechatpayGoodsId
	}

	if o.GoodsName != nil {
		toSerialize["goods_name"] = o.GoodsName
	}

	if o.UnitPrice == nil {
		return nil, fmt.Errorf("field `UnitPrice` is required and must be specified in GoodsDetail")
	}
	toSerialize["unit_price"] = o.UnitPrice

	if o.RefundAmount == nil {
		return nil, fmt.Errorf("field `RefundAmount` is required and must be specified in GoodsDetail")
	}
	toSerialize["refund_amount"] = o.RefundAmount

	if o.RefundQuantity == nil {
		return nil, fmt.Errorf("field `RefundQuantity` is required and must be specified in GoodsDetail")
	}
	toSerialize["refund_quantity"] = o.RefundQuantity
	return json.Marshal(toSerialize)
}

func (o GoodsDetail) String() string {
	var ret string
	if o.MerchantGoodsId == nil {
		ret += "MerchantGoodsId:<nil>, "
	} else {
		ret += fmt.Sprintf("MerchantGoodsId:%v, ", *o.MerchantGoodsId)
	}

	if o.WechatpayGoodsId == nil {
		ret += "WechatpayGoodsId:<nil>, "
	} else {
		ret += fmt.Sprintf("WechatpayGoodsId:%v, ", *o.WechatpayGoodsId)
	}

	if o.GoodsName == nil {
		ret += "GoodsName:<nil>, "
	} else {
		ret += fmt.Sprintf("GoodsName:%v, ", *o.GoodsName)
	}

	if o.UnitPrice == nil {
		ret += "UnitPrice:<nil>, "
	} else {
		ret += fmt.Sprintf("UnitPrice:%v, ", *o.UnitPrice)
	}

	if o.RefundAmount == nil {
		ret += "RefundAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundAmount:%v, ", *o.RefundAmount)
	}

	if o.RefundQuantity == nil {
		ret += "RefundQuantity:<nil>"
	} else {
		ret += fmt.Sprintf("RefundQuantity:%v", *o.RefundQuantity)
	}

	return fmt.Sprintf("GoodsDetail{%s}", ret)
}

func (o GoodsDetail) Clone() *GoodsDetail {
	ret := GoodsDetail{}

	if o.MerchantGoodsId != nil {
		ret.MerchantGoodsId = new(string)
		*ret.MerchantGoodsId = *o.MerchantGoodsId
	}

	if o.WechatpayGoodsId != nil {
		ret.WechatpayGoodsId = new(string)
		*ret.WechatpayGoodsId = *o.WechatpayGoodsId
	}

	if o.GoodsName != nil {
		ret.GoodsName = new(string)
		*ret.GoodsName = *o.GoodsName
	}

	if o.UnitPrice != nil {
		ret.UnitPrice = new(int64)
		*ret.UnitPrice = *o.UnitPrice
	}

	if o.RefundAmount != nil {
		ret.RefundAmount = new(int64)
		*ret.RefundAmount = *o.RefundAmount
	}

	if o.RefundQuantity != nil {
		ret.RefundQuantity = new(int64)
		*ret.RefundQuantity = *o.RefundQuantity
	}

	return &ret
}

// Promotion
type Promotion struct {
	// 券或者立减优惠id
	PromotionId *string `json:"promotion_id"`
	// 枚举值： - GLOBAL- 全场代金券 - SINGLE- 单品优惠 * `GLOBAL` - 全场代金券 * `SINGLE` - 单品优惠
	Scope *Scope `json:"scope"`
	// 枚举值： - COUPON- 代金券，需要走结算资金的充值型代金券 - DISCOUNT- 优惠券，不走结算资金的免充值型优惠券 * `COUPON` - 代金券 * `DISCOUNT` - 优惠券
	Type *Type `json:"type"`
	// 用户享受优惠的金额（优惠券面额=微信出资金额+商家出资金额+其他出资方金额 ），单位为分
	Amount *int64 `json:"amount"`
	// 优惠退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为用户支付的现金，说明详见代金券或立减优惠，单位为分
	RefundAmount *int64 `json:"refund_amount"`
	// 优惠商品发生退款时返回商品信息
	GoodsDetail []GoodsDetail `json:"goods_detail,omitempty"`
}

func (o Promotion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.PromotionId == nil {
		return nil, fmt.Errorf("field `PromotionId` is required and must be specified in Promotion")
	}
	toSerialize["promotion_id"] = o.PromotionId

	if o.Scope == nil {
		return nil, fmt.Errorf("field `Scope` is required and must be specified in Promotion")
	}
	toSerialize["scope"] = o.Scope

	if o.Type == nil {
		return nil, fmt.Errorf("field `Type` is required and must be specified in Promotion")
	}
	toSerialize["type"] = o.Type

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in Promotion")
	}
	toSerialize["amount"] = o.Amount

	if o.RefundAmount == nil {
		return nil, fmt.Errorf("field `RefundAmount` is required and must be specified in Promotion")
	}
	toSerialize["refund_amount"] = o.RefundAmount

	if o.GoodsDetail != nil {
		toSerialize["goods_detail"] = o.GoodsDetail
	}
	return json.Marshal(toSerialize)
}

func (o Promotion) String() string {
	var ret string
	if o.PromotionId == nil {
		ret += "PromotionId:<nil>, "
	} else {
		ret += fmt.Sprintf("PromotionId:%v, ", *o.PromotionId)
	}

	if o.Scope == nil {
		ret += "Scope:<nil>, "
	} else {
		ret += fmt.Sprintf("Scope:%v, ", *o.Scope)
	}

	if o.Type == nil {
		ret += "Type:<nil>, "
	} else {
		ret += fmt.Sprintf("Type:%v, ", *o.Type)
	}

	if o.Amount == nil {
		ret += "Amount:<nil>, "
	} else {
		ret += fmt.Sprintf("Amount:%v, ", *o.Amount)
	}

	if o.RefundAmount == nil {
		ret += "RefundAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundAmount:%v, ", *o.RefundAmount)
	}

	ret += fmt.Sprintf("GoodsDetail:%v", o.GoodsDetail)

	return fmt.Sprintf("Promotion{%s}", ret)
}

func (o Promotion) Clone() *Promotion {
	ret := Promotion{}

	if o.PromotionId != nil {
		ret.PromotionId = new(string)
		*ret.PromotionId = *o.PromotionId
	}

	if o.Scope != nil {
		ret.Scope = new(Scope)
		*ret.Scope = *o.Scope
	}

	if o.Type != nil {
		ret.Type = new(Type)
		*ret.Type = *o.Type
	}

	if o.Amount != nil {
		ret.Amount = new(int64)
		*ret.Amount = *o.Amount
	}

	if o.RefundAmount != nil {
		ret.RefundAmount = new(int64)
		*ret.RefundAmount = *o.RefundAmount
	}

	if o.GoodsDetail != nil {
		ret.GoodsDetail = make([]GoodsDetail, len(o.GoodsDetail))
		for i, item := range o.GoodsDetail {
			ret.GoodsDetail[i] = *item.Clone()
		}
	}

	return &ret
}

// QueryByOutRefundNoRequest
type QueryByOutRefundNoRequest struct {
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo *string `json:"out_refund_no"`
	// 子商户的商户号，由微信支付生成并下发。服务商模式下必须传递此参数
	SubMchid *string `json:"sub_mchid,omitempty"`
}

func (o QueryByOutRefundNoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.OutRefundNo == nil {
		return nil, fmt.Errorf("field `OutRefundNo` is required and must be specified in QueryByOutRefundNoRequest")
	}
	toSerialize["out_refund_no"] = o.OutRefundNo

	if o.SubMchid != nil {
		toSerialize["sub_mchid"] = o.SubMchid
	}
	return json.Marshal(toSerialize)
}

func (o QueryByOutRefundNoRequest) String() string {
	var ret string
	if o.OutRefundNo == nil {
		ret += "OutRefundNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutRefundNo:%v, ", *o.OutRefundNo)
	}

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>"
	} else {
		ret += fmt.Sprintf("SubMchid:%v", *o.SubMchid)
	}

	return fmt.Sprintf("QueryByOutRefundNoRequest{%s}", ret)
}

func (o QueryByOutRefundNoRequest) Clone() *QueryByOutRefundNoRequest {
	ret := QueryByOutRefundNoRequest{}

	if o.OutRefundNo != nil {
		ret.OutRefundNo = new(string)
		*ret.OutRefundNo = *o.OutRefundNo
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	return &ret
}

// Refund
type Refund struct {
	// 微信支付退款号
	RefundId *string `json:"refund_id"`
	// 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo *string `json:"out_refund_no"`
	// 微信支付交易订单号
	TransactionId *string `json:"transaction_id"`
	// 原支付交易对应的商户订单号
	OutTradeNo *string `json:"out_trade_no"`
	// 枚举值： - ORIGINAL—原路退款 - BALANCE—退回到余额 - OTHER_BALANCE—原账户异常退到其他余额账户 - OTHER_BANKCARD—原银行卡异常退到其他银行卡 * `ORIGINAL` - 原路退款 * `BALANCE` - 退回到余额 * `OTHER_BALANCE` - 原账户异常退到其他余额账户 * `OTHER_BANKCARD` - 原银行卡异常退到其他银行卡
	Channel *Channel `json:"channel"`
	// 取当前退款单的退款入账方，有以下几种情况： 1）退回银行卡：{银行名称}{卡类型}{卡尾号} 2）退回支付用户零钱:支付用户零钱 3）退还商户:商户基本账户商户结算银行账户 4）退回支付用户零钱通:支付用户零钱通
	UserReceivedAccount *string `json:"user_received_account"`
	// 退款成功时间，退款状态status为SUCCESS（退款成功）时，返回该字段。遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。
	SuccessTime *time.Time `json:"success_time,omitempty"`
	// 退款受理时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。
	CreateTime *time.Time `json:"create_time"`
	// 退款到银行发现用户的卡作废或者冻结了，导致原路退款银行卡失败，可前往商户平台（pay.weixin.qq.com）-交易中心，手动处理此笔退款。 枚举值： - SUCCESS—退款成功 - CLOSED—退款关闭 - PROCESSING—退款处理中 - ABNORMAL—退款异常 * `SUCCESS` - 退款成功 * `CLOSED` - 退款关闭 * `PROCESSING` - 退款处理中 * `ABNORMAL` - 退款异常
	Status *Status `json:"refund_status"`
	// 退款所使用资金对应的资金账户类型 枚举值： - UNSETTLED : 未结算资金 - AVAILABLE : 可用余额 - UNAVAILABLE : 不可用余额 - OPERATION : 运营户 - BASIC : 基本账户（含可用余额和不可用余额） * `UNSETTLED` - 未结算资金 * `AVAILABLE` - 可用余额 * `UNAVAILABLE` - 不可用余额 * `OPERATION` - 运营户 * `BASIC` - 基本账户（含可用余额和不可用余额）
	FundsAccount *FundsAccount `json:"funds_account,omitempty"`
	// 金额详细信息
	Amount *Amount `json:"amount"`
	// 优惠退款信息
	PromotionDetail []Promotion `json:"promotion_detail,omitempty"`
}

func (o Refund) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.RefundId == nil {
		return nil, fmt.Errorf("field `RefundId` is required and must be specified in Refund")
	}
	toSerialize["refund_id"] = o.RefundId

	if o.OutRefundNo == nil {
		return nil, fmt.Errorf("field `OutRefundNo` is required and must be specified in Refund")
	}
	toSerialize["out_refund_no"] = o.OutRefundNo

	if o.TransactionId == nil {
		return nil, fmt.Errorf("field `TransactionId` is required and must be specified in Refund")
	}
	toSerialize["transaction_id"] = o.TransactionId

	if o.OutTradeNo == nil {
		return nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in Refund")
	}
	toSerialize["out_trade_no"] = o.OutTradeNo

	if o.Channel == nil {
		return nil, fmt.Errorf("field `Channel` is required and must be specified in Refund")
	}
	toSerialize["channel"] = o.Channel

	if o.UserReceivedAccount == nil {
		return nil, fmt.Errorf("field `UserReceivedAccount` is required and must be specified in Refund")
	}
	toSerialize["user_received_account"] = o.UserReceivedAccount

	if o.SuccessTime != nil {
		toSerialize["success_time"] = o.SuccessTime.Format(time.RFC3339)
	}

	if o.CreateTime == nil {
		return nil, fmt.Errorf("field `CreateTime` is required and must be specified in Refund")
	}
	toSerialize["create_time"] = o.CreateTime.Format(time.RFC3339)

	if o.Status == nil {
		return nil, fmt.Errorf("field `Status` is required and must be specified in Refund")
	}
	toSerialize["status"] = o.Status

	if o.FundsAccount != nil {
		toSerialize["funds_account"] = o.FundsAccount
	}

	if o.Amount == nil {
		return nil, fmt.Errorf("field `Amount` is required and must be specified in Refund")
	}
	toSerialize["amount"] = o.Amount

	if o.PromotionDetail != nil {
		toSerialize["promotion_detail"] = o.PromotionDetail
	}
	return json.Marshal(toSerialize)
}

func (o Refund) String() string {
	var ret string
	if o.RefundId == nil {
		ret += "RefundId:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundId:%v, ", *o.RefundId)
	}

	if o.OutRefundNo == nil {
		ret += "OutRefundNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutRefundNo:%v, ", *o.OutRefundNo)
	}

	if o.TransactionId == nil {
		ret += "TransactionId:<nil>, "
	} else {
		ret += fmt.Sprintf("TransactionId:%v, ", *o.TransactionId)
	}

	if o.OutTradeNo == nil {
		ret += "OutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutTradeNo:%v, ", *o.OutTradeNo)
	}

	if o.Channel == nil {
		ret += "Channel:<nil>, "
	} else {
		ret += fmt.Sprintf("Channel:%v, ", *o.Channel)
	}

	if o.UserReceivedAccount == nil {
		ret += "UserReceivedAccount:<nil>, "
	} else {
		ret += fmt.Sprintf("UserReceivedAccount:%v, ", *o.UserReceivedAccount)
	}

	if o.SuccessTime == nil {
		ret += "SuccessTime:<nil>, "
	} else {
		ret += fmt.Sprintf("SuccessTime:%v, ", *o.SuccessTime)
	}

	if o.CreateTime == nil {
		ret += "CreateTime:<nil>, "
	} else {
		ret += fmt.Sprintf("CreateTime:%v, ", *o.CreateTime)
	}

	if o.Status == nil {
		ret += "Status:<nil>, "
	} else {
		ret += fmt.Sprintf("Status:%v, ", *o.Status)
	}

	if o.FundsAccount == nil {
		ret += "FundsAccount:<nil>, "
	} else {
		ret += fmt.Sprintf("FundsAccount:%v, ", *o.FundsAccount)
	}

	ret += fmt.Sprintf("Amount:%v, ", o.Amount)

	ret += fmt.Sprintf("PromotionDetail:%v", o.PromotionDetail)

	return fmt.Sprintf("Refund{%s}", ret)
}

func (o Refund) Clone() *Refund {
	ret := Refund{}

	if o.RefundId != nil {
		ret.RefundId = new(string)
		*ret.RefundId = *o.RefundId
	}

	if o.OutRefundNo != nil {
		ret.OutRefundNo = new(string)
		*ret.OutRefundNo = *o.OutRefundNo
	}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	if o.OutTradeNo != nil {
		ret.OutTradeNo = new(string)
		*ret.OutTradeNo = *o.OutTradeNo
	}

	if o.Channel != nil {
		ret.Channel = new(Channel)
		*ret.Channel = *o.Channel
	}

	if o.UserReceivedAccount != nil {
		ret.UserReceivedAccount = new(string)
		*ret.UserReceivedAccount = *o.UserReceivedAccount
	}

	if o.SuccessTime != nil {
		ret.SuccessTime = new(time.Time)
		*ret.SuccessTime = *o.SuccessTime
	}

	if o.CreateTime != nil {
		ret.CreateTime = new(time.Time)
		*ret.CreateTime = *o.CreateTime
	}

	if o.Status != nil {
		ret.Status = new(Status)
		*ret.Status = *o.Status
	}

	if o.FundsAccount != nil {
		ret.FundsAccount = new(FundsAccount)
		*ret.FundsAccount = *o.FundsAccount
	}

	if o.Amount != nil {
		ret.Amount = o.Amount.Clone()
	}

	if o.PromotionDetail != nil {
		ret.PromotionDetail = make([]Promotion, len(o.PromotionDetail))
		for i, item := range o.PromotionDetail {
			ret.PromotionDetail[i] = *item.Clone()
		}
	}

	return &ret
}

// ReqFundsAccount * `AVAILABLE` - 可用余额, 仅对老资金流商户适用，指定从可用余额账户出资
type ReqFundsAccount string

func (e ReqFundsAccount) Ptr() *ReqFundsAccount {
	return &e
}

// Enums of ReqFundsAccount
const (
	REQFUNDSACCOUNT_AVAILABLE ReqFundsAccount = "AVAILABLE"
)

// Scope * `GLOBAL` - 全场代金券, 全场优惠类型 * `SINGLE` - 单品优惠, 单品优惠类型
type Scope string

func (e Scope) Ptr() *Scope {
	return &e
}

// Enums of Scope
const (
	SCOPE_GLOBAL Scope = "GLOBAL"
	SCOPE_SINGLE Scope = "SINGLE"
)

// Status * `SUCCESS` - 退款成功, 退款状态 * `CLOSED` - 退款关闭, 退款状态 * `PROCESSING` - 退款处理中, 退款状态 * `ABNORMAL` - 退款异常, 退款状态
type Status string

func (e Status) Ptr() *Status {
	return &e
}

// Enums of Status
const (
	STATUS_SUCCESS    Status = "SUCCESS"
	STATUS_CLOSED     Status = "CLOSED"
	STATUS_PROCESSING Status = "PROCESSING"
	STATUS_ABNORMAL   Status = "ABNORMAL"
)

// Type * `COUPON` - 代金券, 代金券类型，需要走结算资金的充值型代金券 * `DISCOUNT` - 优惠券, 优惠券类型，不走结算资金的免充值型优惠券
type Type string

func (e Type) Ptr() *Type {
	return &e
}

// Enums of Type
const (
	TYPE_COUPON   Type = "COUPON"
	TYPE_DISCOUNT Type = "DISCOUNT"
)
