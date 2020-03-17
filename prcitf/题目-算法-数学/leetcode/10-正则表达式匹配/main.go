// Copyright (c) liuhao. 2012-2050. All rights reserved.
package pp

var dp [][]int

func isMatch(s string, p string) bool {
	//defer func(){
	//	fmt.Printf("%#v \n",dp)
	//}()
	dp = make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(p)+1)
	}
	return subStrIsMatch(s,p)
}

func subStrIsMatch(s, p string)(result bool) {
	defer func(){
		if result{
			dp[len(s)][len(p)]=1
		}else {
			dp[len(s)][len(p)]=-1
		}
	}()

	if dp[len(s)][len(p)]!=0{
		return dp[len(s)][len(p)]==1
	}

	if len(s) == 0 {
		if len(p) == 0 {
			return true
		}
		// s len ==0,p只能是x*x* ...
		if len(p)%2 != 0 ||p[1]!='*'{
			return false
		}
		return subStrIsMatch(s, p[2:])
	}
	if len(p)==0{
		return false
	}

	if len(p)>=2&&p[1]=='*'{
		if s[0]!=p[0]&&p[0]!='.'{
			return subStrIsMatch(s,p[2:])
		}
		return subStrIsMatch(s,p[2:])||subStrIsMatch(s[1:],p[2:])||subStrIsMatch(s[1:],p)
	}

	if s[0]==p[0]||p[0]=='.'{
		return  subStrIsMatch(s[1:],p[1:])
	}
	return false

}
