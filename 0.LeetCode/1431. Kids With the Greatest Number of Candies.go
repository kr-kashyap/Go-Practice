/*
    https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
    250822
*/
func kidsWithCandies(candies []int, extraCandies int) []bool {
    i:=0
    var ans = make([]bool, len(candies))
    max:=candies[0]
    for i=1;i< len(candies); i++ {
        if candies[i] > max {
            max = candies[i]
        }
    }
    for i=0;i< len(candies); i++ {
        if candies[i] + extraCandies >= max {
            ans[i]=true    
        } else {
            ans[i]=false    
        }
    }
    return ans
}