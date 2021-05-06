package stoploss

import (
	"context"
	"strconv"
	"strings")

// Connecting to the Binance exchange
type Binance struct {
	api *binance.Client
	ctx context.Context
    
}

// Creating an instanse module for exchange
func NewExchange(ctx context.Context, api *binance.Client) *Binance {
	return &Binance{api, ctx}
}

// Returning the balance of a given pair
func (exchange *Binance) GetBalance(coin string) (string, error) {
    // Changing the pair format
	coin = strings.ToUpper(coin)
    // sending request 	 
	account, err := exchange.api.NewGetAccountService().Do(exchange.ctx)
	if err != nil {
		return "0", err}
    
    //  find the coin 
	for _, balance := range account.Balances {
		if strings.ToUpper(balance.Asset) == coin {
			return balance.Free, nil
		}
	}
	return "0", nil
}

// Return the last trading price of a given pair
func (exchange *Binance) GetMarketPrice(market string) (float64, error) {
	prices, err := exchange.api.NewListPricesService().Symbol(market).Do(exchange.ctx)
	if err != nil {
		return 0, err
	}   
    // Find the pair 	
	for _, p := range prices {
		if p.Symbol == market {
		    //  changing the type of the price
			return strconv.ParseFloat(p.Price, 64)
		}
	}

	return 0, nil
}

// Get the bestBid for selling
func (exchange *Binance) Sell(market string, quantity string) (string, error) {
	order, err := exchange.api.NewCreateOrderService().Symbol(market).
		Side(binance.SideTypeSell).Type(binance.OrderTypeMarket).
		Quantity(quantity).Do(exchange.ctx)

	if err != nil {
		return "", err
	}
    // Returning the last 10 orders
	return strconv.FormatInt(order.OrderID, 10), nil
}

// Get the bestASk to buy
func (exchange *Binance) Buy(market string, quantity string) (string, error) {
	order, err := exchange.api.NewCreateOrderService().Symbol(market).
		Side(binance.SideTypeBuy).Type(binance.OrderTypeMarket).
		Quantity(quantity).Do(exchange.ctx)

	if err != nil {
		return "", err
	}

	return strconv.FormatInt(order.OrderID, 10), nil
}
