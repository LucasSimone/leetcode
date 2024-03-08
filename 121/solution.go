func maxProfit(prices []int) int {

    var min_index int = 0
    var max_profit int = 0

    for i:= 1; i < len(prices); i++{
        
        var profit int = prices[i] - prices[min_index]
        
        if profit > max_profit{
            max_profit = profit
        }

        if prices[i] < prices[min_index]{
            min_index = i
        }

    }

    return max_profit
}
