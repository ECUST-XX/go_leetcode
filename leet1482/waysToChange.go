package leet1482

var resMap = []struct{
	num int
	r [4]int
}{
	{
		5,
		[4]int{0,0,0,5},
	},
	{
		10,
		[4]int{0,0,2,0},
	},
	{
		10,
		[4]int{0,0,1,5},
	},
	{
		10,
		[4]int{0,0,0,10},
	},
	{
		25,
		[4]int{0,2,1,0},
	},
}



// todo:动态规划我不会。。。。方程偷鸡也搞不出来。。。

var res = 0
func WaysToChange(n int) int {
	r := []int{0, 0, 0, 0}

	r25 := n / 25
	p25 := n % 25

	r10 := p25 / 10
	p10 := p25 % 10

	r5 := p10 / 5
	p5 := p10 % 5

	r[0] = r25
	r[1] = r10
	r[2] = r5
	r[3] = p5

    res++






	return 0
}
