func maxProfit(prices []int) int {
    var current_min int = 0
    var profit int = 0

    for i:= 1; i < len(prices); i++{
        
        var current_profit int = prices[i] - prices[current_min]

        //If there is a profit sell
        if current_profit > 0{
            profit += current_profit
            current_min = i
        }else 
        
        //If there is a lower min move the min pointer
        if prices[i] < prices[current_min]{
            current_min = i
        }

    }
    return profit
}

