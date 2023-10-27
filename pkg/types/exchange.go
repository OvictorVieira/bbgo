package types

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/OvictorVieira/bbgo/pkg/fixedpoint"
)

const DateFormat = "2006-01-02"

type ExchangeId int

func (eid *ExchangeId) Value() (driver.Value, error) {
	return eid.String(), nil
}

func (eid *ExchangeId) UnmarshalJSON(data []byte) error {
	var exchangeId int
	if err := json.Unmarshal(data, &exchangeId); err != nil {
		return err
	}

	switch exchangeId {
	case 1, 2, 3, 4:
		*eid = ExchangeId(exchangeId)
		return nil

	}

	return fmt.Errorf("unknown or unsupported exchange name: %d, valid names are: max, binance, okex, kucoin", exchangeId)
}

func (eid ExchangeId) String() string {
	return string(rune(eid))
}

const (
	ExchangeMax      ExchangeId = 1
	ExchangeBinance  ExchangeId = 2
	ExchangeOKEx     ExchangeId = 3
	ExchangeKucoin   ExchangeId = 4
	ExchangeBitget   ExchangeId = 5
	ExchangeBacktest ExchangeId = 6
	ExchangeBybit    ExchangeId = 7
)

var SupportedExchanges = []ExchangeId{
	ExchangeMax,
	ExchangeBinance,
	ExchangeOKEx,
	ExchangeKucoin,
	ExchangeBitget,
	ExchangeBybit,
	// note: we are not using "backtest"
}

func ValidExchangeId(a int) (ExchangeId, error) {
	for _, n := range SupportedExchanges {
		if n == ExchangeId(a) { // Convert 'a' to 'ExchangeId' before comparing
			return n, nil
		}
	}

	return 0, fmt.Errorf("invalid exchange id: %d", a) // use %d for integer formatting
}

type ExchangeMinimal interface {
	Name() ExchangeId
	PlatformFeeCurrency() string
}

//go:generate mockgen -destination=mocks/mock_exchange.go -package=mocks . Exchange
type Exchange interface {
	ExchangeMinimal
	ExchangeMarketDataService
	ExchangeAccountService
	ExchangeTradeService
}

// ExchangeBasic is the new type for replacing the original Exchange interface
type ExchangeBasic = Exchange

// ExchangeOrderQueryService provides an interface for querying the order status via order ID or client order ID
//
//go:generate mockgen -destination=mocks/mock_exchange_order_query.go -package=mocks . ExchangeOrderQueryService
type ExchangeOrderQueryService interface {
	QueryOrder(ctx context.Context, q OrderQuery) (*Order, error)
	QueryOrderTrades(ctx context.Context, q OrderQuery) ([]Trade, error)
}

type ExchangeAccountService interface {
	QueryAccount(ctx context.Context) (*Account, error)

	QueryAccountBalances(ctx context.Context) (BalanceMap, error)
}

type ExchangeTradeService interface {
	SubmitOrder(ctx context.Context, order SubmitOrder) (createdOrder *Order, err error)

	QueryOpenOrders(ctx context.Context, symbol string) (orders []Order, err error)

	CancelOrders(ctx context.Context, orders ...Order) error
}

type ExchangeDefaultFeeRates interface {
	DefaultFeeRates() ExchangeFee
}

type ExchangeAmountFeeProtect interface {
	SetModifyOrderAmountForFee(ExchangeFee)
}

//go:generate mockgen -destination=mocks/mock_exchange_trade_history.go -package=mocks . ExchangeTradeHistoryService
type ExchangeTradeHistoryService interface {
	QueryTrades(ctx context.Context, symbol string, options *TradeQueryOptions) ([]Trade, error)
	QueryClosedOrders(ctx context.Context, symbol string, since, until time.Time, lastOrderID uint64) (orders []Order, err error)
}

type ExchangeMarketDataService interface {
	NewStream() Stream

	QueryMarkets(ctx context.Context) (MarketMap, error)

	QueryTicker(ctx context.Context, symbol string) (*Ticker, error)

	QueryTickers(ctx context.Context, symbol ...string) (map[string]Ticker, error)

	QueryKLines(ctx context.Context, symbol string, interval Interval, options KLineQueryOptions) ([]KLine, error)
}

type CustomIntervalProvider interface {
	SupportedInterval() map[Interval]int
	IsSupportedInterval(interval Interval) bool
}

type ExchangeTransferService interface {
	QueryDepositHistory(ctx context.Context, asset string, since, until time.Time) (allDeposits []Deposit, err error)
	QueryWithdrawHistory(ctx context.Context, asset string, since, until time.Time) (allWithdraws []Withdraw, err error)
}

type ExchangeWithdrawalService interface {
	Withdraw(ctx context.Context, asset string, amount fixedpoint.Value, address string, options *WithdrawalOptions) error
}

type ExchangeRewardService interface {
	QueryRewards(ctx context.Context, startTime time.Time) ([]Reward, error)
}

type TradeQueryOptions struct {
	StartTime   *time.Time
	EndTime     *time.Time
	Limit       int64
	LastTradeID uint64
}
