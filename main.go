mas2
auto-arbitrage-bot
A Go-based crypto arbitrage bot scanning multiple DEXs for profitable opportunities and executing trades automatically.
package main

import (
    "fmt"
    "math/big"
    "time"
)

type Dex struct {
    Name string
    Price *big.Float
}

func getPrice(d Dex) *big.Float {
    // Mock: Replace with real API calls
    return d.Price
}

func main() {
    dexes := []Dex{
        {"Uniswap", big.NewFloat(1820.50)},
        {"Sushiswap", big.NewFloat(1832.20)},
        {"PancakeSwap", big.NewFloat(1815.75)},
    }

    for {
        for i := 0; i < len(dexes); i++ {
            for j := i+1; j < len(dexes); j++ {
                priceDiff := new(big.Float).Sub(dexes[i].Price, dexes[j].Price)
                if priceDiff.Abs(priceDiff).Cmp(big.NewFloat(5.0)) > 0 {
                    fmt.Printf("Arbitrage found: Buy on %s, sell on %s (Diff: %s USD)\n", dexes[j].Name, dexes[i].Name, priceDiff.Text('f', 2))
                }
            }
        }
        time.Sleep(10 * time.Second)
    }
}
