package basics


func solve(i int,j int,s string,t string) int{
	if i>=len(s) || j>=len(t) {
		return 0;		
	}
	ans :=1
	if s[i]==t[j] {
		ans = max(ans,1+solve(i+1,j+1,s,t))
	} else {
		val:= max(solve(i,j+1,s,t),solve(i+1,j,s,t))
		ans = max(ans,val)
	}
	return ans

}

func longestCommonSubsequence(s string, t string) int {

    m:=len(s)
    n:=len(t)
    dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := n-1; i >= 0; i-- {
		
		for j := m-1; j >=0; j-- {
			if s[i]==t[j] {
				dp[i][j]= max(dp[i][j],dp[i+1][j+1])
			}else{
				val:=max(dp[i+1][j],dp[i][j+1])
				dp[i][j] = max(dp[i][j],val)
			}
			
		}
		
	}
	return dp[0][0]
    


	// return solve(0,0,s,t)
    
    
}