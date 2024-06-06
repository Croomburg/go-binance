package binance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Croomburg/go-binance/v2/futures"
)

// CreateFutureAlgoTwapOrderService create future algo order
type CreateFutureAlgoTwapOrderService struct {
	c            *Client
	symbol       string
	side         SideType
	positionSide *futures.PositionSideType
	quantity     *string
	duration     int64
	clientAlgoId *string
	reduceOnly   *bool
	limitPrice   *string
}

// Symbol set symbol
func (s *CreateFutureAlgoTwapOrderService) Symbol(symbol string) *CreateFutureAlgoTwapOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateFutureAlgoTwapOrderService) Side(side SideType) *CreateFutureAlgoTwapOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateFutureAlgoTwapOrderService) PositionSide(positionSide futures.PositionSideType) *CreateFutureAlgoTwapOrderService {
	s.positionSide = &positionSide
	return s
}

// Quantity set quantity
func (s *CreateFutureAlgoTwapOrderService) Quantity(quantity string) *CreateFutureAlgoTwapOrderService {
	s.quantity = &quantity
	return s
}

// Duration set duration
func (s *CreateFutureAlgoTwapOrderService) Duration(duration int64) *CreateFutureAlgoTwapOrderService {
	s.duration = duration
	return s
}

// ClientAlgoId set clientAlgoId
func (s *CreateFutureAlgoTwapOrderService) ClientAlgoId(clientAlgoId string) *CreateFutureAlgoTwapOrderService {
	s.clientAlgoId = &clientAlgoId
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateFutureAlgoTwapOrderService) ReduceOnly(reduceOnly bool) *CreateFutureAlgoTwapOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// LimitPrice set limitPrice
func (s *CreateFutureAlgoTwapOrderService) LimitPrice(limitPrice string) *CreateFutureAlgoTwapOrderService {
	s.limitPrice = &limitPrice
	return s
}

func (s *CreateFutureAlgoTwapOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":   s.symbol,
		"side":     s.side,
		"quantity": s.quantity,
		"duration": s.duration,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.clientAlgoId != nil {
		m["clientAlgoId"] = *s.clientAlgoId
	}
	if s.limitPrice != nil {
		m["limitPrice"] = *s.limitPrice
	}
	r.setFormParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *CreateFutureAlgoTwapOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateFutureAlgoTwapOrderResponse, err error) {
	data, err := s.createOrder(ctx, "/sapi/v1/algo/futures/newOrderTwap", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateFutureAlgoTwapOrderResponse)
	err = json.Unmarshal(data, res)
	fmt.Printf("response is %v", data)

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateFutureAlgoTwapOrderResponse define create future algo twap order response
type CreateFutureAlgoTwapOrderResponse struct {
	ClientAlgoId string `json:"clientAlgoId"`
	Success      bool   `json:"success"`
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
}

// ListOpenFutureAlgoOrderService list current open future algo orders
type ListOpenFutureAlgoOrderService struct {
	c *Client
}

// Do send request
func (s *ListOpenFutureAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (res *ListOpenFutureAlgoOrderResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/algo/futures/openOrders",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	fmt.Println("data is ", data)
	fmt.Println("err is ", err)
	res = new(ListOpenFutureAlgoOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type FutureAlgoOrder struct {
	//策略订单ID
	AlgoId       int64                    `json:"algoId"`
	Symbol       string                   `json:"symbol"`
	Side         SideType                 `json:"side"`
	PositionSide futures.PositionSideType `json:"positionSide"`
	TotalQty     string                   `json:"totalQty"`
	ExecutedQty  string                   `json:"executedQty"`
	ExecutedAmt  string                   `json:"executedAmt"`
	AvgPrice     string                   `json:"avgPrice"`
	ClientAlgoId string                   `json:"clientAlgoId"`
	BookTime     int64                    `json:"bookTime"`
	EndTime      int64                    `json:"endTime"`
	// 策略订单状态 WORKING
	AlgoStatus FutureAlgoOrderStatusType `json:"algoStatus"`
	AlgoType   FutureAlgoType            `json:"algoType"`
	Urgency    FutureAlgoUrgencyType     `json:"urgency"`
}

// ListOpenFutureAlgoOrderResponse define response of list open future algo orders
type ListOpenFutureAlgoOrderResponse struct {
	Total  int64              `json:"total"`
	Orders []*FutureAlgoOrder `json:"orders"`
}

// ListFutureAlgoOrderHistoryService list future algo historical orders
type ListFutureAlgoOrderHistoryService struct {
	c         *Client
	symbol    string
	side      *SideType
	startTime *int64
	endTime   *int64
	page      *int
	pageSize  *int
}

// Symbol set symbol
func (s *ListFutureAlgoOrderHistoryService) Symbol(symbol string) *ListFutureAlgoOrderHistoryService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *ListFutureAlgoOrderHistoryService) Side(side SideType) *ListFutureAlgoOrderHistoryService {
	s.side = &side
	return s
}

// StartTime set startTime
func (s *ListFutureAlgoOrderHistoryService) StartTime(startTime int64) *ListFutureAlgoOrderHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListFutureAlgoOrderHistoryService) EndTime(endTime int64) *ListFutureAlgoOrderHistoryService {
	s.endTime = &endTime
	return s
}

// Page set page
func (s *ListFutureAlgoOrderHistoryService) Page(page int) *ListFutureAlgoOrderHistoryService {
	s.page = &page
	return s
}

// PageSize set pageSize
func (s *ListFutureAlgoOrderHistoryService) PageSize(pageSize int) *ListFutureAlgoOrderHistoryService {
	s.pageSize = &pageSize
	return s
}

// Do send request
func (s *ListFutureAlgoOrderHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *ListFutureAlgoOrderHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/algo/futures/historicalOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	if s.side != nil {
		r.setParam("side", s.side)
	}
	if s.startTime != nil {
		r.setParam("startTime", s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", s.endTime)
	}
	if s.page != nil {
		r.setParam("page", s.page)
	}
	if s.pageSize != nil {
		r.setParam("pageSize", s.pageSize)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	fmt.Println("data is ", data)
	fmt.Println("err is ", err)
	res = new(ListFutureAlgoOrderHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListFutureAlgoOrderHistoryResponse defines response of list future algo historical orders
type ListFutureAlgoOrderHistoryResponse struct {
	Total  int64              `json:"total"`
	Orders []*FutureAlgoOrder `json:"orders"`
}

// CancelFutureAlgoOrderService cancel future algo twap order
type CancelFutureAlgoOrderService struct {
	c      *Client
	algoId *int64
}

func (s *CancelFutureAlgoOrderService) AlgoId(algoId int64) *CancelFutureAlgoOrderService {
	s.algoId = &algoId
	return s
}

// Do send request
func (s *CancelFutureAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelFutureAlgoOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/algo/futures/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("algoId", s.algoId)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	fmt.Println("data is ", data)
	fmt.Println("err is ", err)
	res = new(CancelFutureAlgoOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelFutureAlgoOrderResponse define response of cancel future algo twap order
type CancelFutureAlgoOrderResponse struct {
	AlgoId  int64  `json:"algoId"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
}

// GetFutureAlgoSubOrderService get future algo sub orders
type GetFutureAlgoSubOrderService struct {
	c        *Client
	algoId   int64
	page     *int
	pageSize *int
}

// AlgoId set algoId
func (s *GetFutureAlgoSubOrderService) AlgoId(algoId int64) *GetFutureAlgoSubOrderService {
	s.algoId = algoId
	return s
}

// Page set page
func (s *GetFutureAlgoSubOrderService) Page(page int) *GetFutureAlgoSubOrderService {
	s.page = &page
	return s
}

// PageSize set pageSize
func (s *GetFutureAlgoSubOrderService) PageSize(pageSize int) *GetFutureAlgoSubOrderService {
	s.pageSize = &pageSize
	return s
}

// Do send request
func (s *GetFutureAlgoSubOrderService) Do(ctx context.Context, opts ...RequestOption) (res *GetFutureAlgoSubOrderResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/algo/futures/subOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("algoId", s.algoId)
	if s.page != nil {
		r.setFormParam("page", s.page)
	}
	if s.pageSize != nil {
		r.setFormParam("pageSize", s.pageSize)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	fmt.Println("data is ", data)
	fmt.Println("err is ", err)
	res = new(GetFutureAlgoSubOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// FutureAlgoSubOrder definen sub order of future algo order
type FutureAlgoSubOrder struct {
	AlgoId  int64    `json:"algoId"`
	OrderId int64    `json:"orderId"`
	Symbol  string   `json:"symbol"`
	Side    SideType `json:"side"`
	// TODO 改成enum
	OrderStatus string          `json:"orderStatus"`
	ExecutedQty string          `json:"executedQty"`
	ExecutedAmt string          `json:"executedAmt"`
	FeeAmt      string          `json:"feeAmt"`
	FeeAsset    string          `json:"feeAsset"`
	AvgPrice    string          `json:"avgPrice"`
	BookTime    int64           `json:"bookTime"`
	SubId       int64           `json:"subId"`
	TimeInForce TimeInForceType `json:"timeInForce"`
	OrigQty     string          `json:"origQty"`
}

type GetFutureAlgoSubOrderResponse struct {
	Total       int64                 `json:"total"`
	ExecutedQty string                `json:"executedQty"`
	ExecutedAmt string                `json:"executedAmt"`
	SubOrders   []*FutureAlgoSubOrder `json:"subOrders"`
}
