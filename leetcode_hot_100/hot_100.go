package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	t := []int{0, 1, 2, 1, 1, 2, 0}
	sortColors(t)
	fmt.Println(t)
}

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow

}

func sortColors(nums []int) {
	ll := len(nums)
	k0 := 0
	k1 := 0
	for i := 0; i < ll; i++ {
		if nums[i] == 0 {
			nums[k0], nums[i] = nums[i], nums[k0]
			if k0 < k1 {
				nums[i], nums[k1] = nums[k1], nums[i]
			}
			k0++
			k1++
		} else if nums[i] == 1 {
			nums[k1], nums[i] = nums[i], nums[k1]
			k1++
		}
	}
}

func majorityElement(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	max := 0
	mk := -1
	for k, v := range m {
		if v > max {
			mk = k
			max = v
		}
	}

	return mk
}

func singleNumber(nums []int) int {
	m := map[int]int{}

	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return 0
}

// todo:1143
func longestCommonSubsequence(text1 string, text2 string) int {
	m1 := map[byte][]int{}
	for i := 0; i < len(text2); i++ {
		if _, ok := m1[text2[i]]; !ok {
			m1[text2[i]] = []int{}
		}
	}

	for i := 0; i < len(text1); i++ {
		if _, ok := m1[text1[i]]; !ok {
			continue
		}
		m1[text1[i]] = append(m1[text1[i]], i)
	}

	res := 0
	for i := 0; i < len(text2); i++ {
		ml := 0
		start := -1
		for j := i; j < len(text2); j++ {
			for _, k := range m1[text2[j]] {
				if k > start {
					start = k
					ml++
					break
				}
			}
		}
		if ml > res {
			res = ml
		}
	}

	return res
}

func longestPalindrome(s string) string {
	ll := len(s)
	if ll <= 1 {
		return s
	}

	mres := map[byte][]int{}
	ml := []int{0, 0}
	for i := 0; i < ll; i++ {
		if _, ok := mres[s[i]]; !ok {
			mres[s[i]] = []int{}
		}
		mres[s[i]] = append(mres[s[i]], i)
	}

	for _, v := range mres {
		a := longestP(v, ml, s)
		if a != nil {
			fmt.Println(s[a[0]:a[1]+1], a)
			if ml[1]-ml[0] < a[1]-a[0] {
				ml = a
			}
		}
	}

	return s[ml[0] : ml[1]+1]
}

func longestP(arr, ml []int, s string) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			m := arr[i]
			n := arr[j]
			if ml[0] < m && ml[1] > n {
				break
			}
			flag := true
			for m < n {
				if s[m] != s[n] {
					flag = false
					break
				}
				m++
				n--
			}
			if flag {
				if ml[1]-ml[0] < arr[j]-arr[i] {
					ml[1] = arr[j]
					ml[0] = arr[i]
				}
			}
		}
	}

	return nil
}

func minPathSum(grid [][]int) int {
	res := make([][]int, len(grid))
	t := 0
	for k := range grid {
		t += grid[k][0]
		res[k] = make([]int, len(grid[0]))
		res[k][0] = t
	}

	t = 0
	for k, v := range grid[0] {
		t += v
		res[0][k] = t
	}

	for i := 1; i < len(grid); i++ {

		for j := 1; j < len(grid[0]); j++ {
			res[i][j] = min(res[i-1][j], res[i][j-1]) + grid[i][j]
		}
	}

	return res[len(res)-1][len(res[0])-1]
}

func uniquePaths(m int, n int) int {
	t1 := (m + n - 2)
	t2 := m - 1
	t3 := n - 1

	t0 := t3
	if t2 > t3 {
		t0 = (t2)
	}

	res := 1
	j := 1
	for i := t0 + 1; i <= t1; i++ {
		res = res * (i) / j
		j++
	}

	return int(res)
}

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	for _, v := range nums {
		if v > sum {
			return false
		}
		if v == sum {
			return true
		}
	}
	sort.Ints(nums)

	ll := len(nums)
	for i := ll - 1; i > ll/2-1; i-- {
		for z := 1; z <= i; z++ {
			t := nums[i]
			arr := []int{}
			j := i - z
			for ; j >= 0; j-- {
				t += nums[j]
				if t == sum {
					return true
				} else if t < sum {
					if j == 0 && len(arr) != 0 {
						t -= nums[j]
						jt := arr[len(arr)-1]
						t -= nums[jt]
						arr = arr[:len(arr)-1]
						j = jt
					} else {
						arr = append(arr, j)
					}
					continue
				} else {
					t -= nums[j]
					if j == 0 && len(arr) != 0 {
						jt := arr[len(arr)-1]
						t -= nums[jt]
						arr = arr[:len(arr)-1]
						j = jt
					}
				}
			}
		}
	}

	return false
}

func maxProduct(nums []int) int {
	res := make([]int, len(nums))
	res2 := make([]int, len(nums))
	res[0] = nums[0]
	res2[0] = nums[0]
	maxRes := nums[0]
	ll := len(nums)
	for i := 1; i < ll; i++ {
		t := nums[i] * res[i-1]
		t1 := nums[i] * res2[i-1]
		res[i] = max(nums[i], t, t1)
		res2[i] = min(nums[i], t, t1)

		maxRes = max(maxRes, res[i], res2[i])
	}

	return maxRes
}

// TODO 300
func lengthOfLIS(nums []int) int {
	res := make([]int, len(nums))
	if len(nums) == 1 {
		return 1
	}
	res[0] = 1
	maxRes := 1
	for i := 1; i < len(nums); i++ {
		res[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				res[i] = max(res[i], res[j]+1)
			}
		}
		maxRes = max(maxRes, res[i])
	}

	return maxRes
}

func wordBreak(s string, wordDict []string) bool {
	i := 0
	res := make([]int, len(s)+1)
	res2 := make([]int, len(s)+1)
	ll := len(s)
	for i < ll {
		t := false
		for j := res[i]; j < len(wordDict); j++ {
			v := wordDict[j]
			if i+len(v) > ll {
				continue
			}
			if s[i:i+len(v)] == v {
				res2[i+len(v)] = i
				res[i] = j
				i += len(v)
				t = true
				break
			}
		}
		if !t {
			i = i - len(wordDict[res[res2[i]]])
			if i < 0 {
				return false
			}
			res[i]++
			if res[i] >= len(wordDict) {
				return false
			}
		}
	}

	return i == ll
}

func coinChange(coins []int, amount int) int {
	res := make([]int, amount+1)
	ll := len(coins)

	res[0] = 0
	for i := 1; i <= amount; i++ {
		t := math.MaxUint32
		for j := 0; j < ll; j++ {
			if coins[j] <= i {
				t = min(t, res[i-coins[j]])
			}

		}
		res[i] = t + 1
	}
	if res[amount] > math.MaxUint32 {
		return -1
	}
	return res[amount]
}

// TODO 279
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			t := i - j*j
			minn = min(minn, f[t])
		}
		f[i] = minn + 1
	}
	return f[n]
}

// TODO 198
func rob(nums []int) int {
	ll := len(nums)
	res := make([]int, ll)
	if ll == 1 {
		return nums[0]
	}
	if ll == 2 {
		if nums[1] > nums[0] {
			return nums[1]
		}
		return nums[0]
	}

	res[0] = nums[0]
	res[1] = max(nums[0], nums[1])

	for i := 2; i < ll; i++ {
		res[i] = max(res[i-2]+nums[i], res[i-1])
	}

	return res[len(res)-1]
}

func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)
	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		t := make([]int, 0)
		t = append(t, 1)
		for j := 0; j < i-1; j++ {
			t = append(t, res[i-1][j]+res[i-1][j+1])
		}
		t = append(t, 1)
		res = append(res, t)
	}

	return res
}

// TODO 70 斐波那契数列
func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func partitionLabels(s string) []int {
	res := []int{}
	m := map[rune][]int{}
	sl := []rune{}
	ll := len(s)
	for k, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = []int{k, k}
			sl = append(sl, v)
		}
		m[v][1] = k
	}

	max := m[sl[0]][1]
	min := m[sl[0]][0]
	sum := 0
	for _, v := range sl {
		if m[v][0] <= max {
			if m[v][1] > max {
				max = m[v][1]
			}
		} else {
			res = append(res, max-min+1)
			min = m[v][0]
			max = m[v][1]
			sum += max - min + 1
		}
		// fmt.Println(string(v), m[v])
	}
	if sum != ll {
		res = append(res, max-min+1)
	}
	return res
}

func jump(nums []int) int {
	res := 0
	ll := len(nums)
	for i := 0; i < ll-1; {
		t := nums[i] + i + 1
		if t >= ll {
			return res + 1
		}
		res++

		tm := t
		ti := 1
		for j := 1; j <= nums[i]; j++ {
			t2 := nums[j+i] + i + j + 1
			if t2 >= ll {
				return res + 1
			}
			if t2 > tm {
				ti = j
				tm = t2
			}
		}

		i += ti
	}

	return res
}

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	ll := len(nums)
	sy := 1
	for i := 0; i < ll-1; i++ {
		sy--
		if sy < 0 {
			return false
		}
		if sy < nums[i] {
			sy = nums[i]
		}

		t := nums[i] + i + 1
		if t >= ll {
			return true
		}
	}

	return false
}

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	res := 0
	min := math.MaxInt64
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}
		if prices[i]-min > res {
			res = prices[i] - min
		}
	}

	return res
}

func topKFrequent(nums []int, k int) []int {
	m := map[int][]int{}
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = []int{v, 0}
		}
		m[v][1]++
	}
	res := make([]int, 0)
	t := [][]int{}
	for _, v := range m {
		t = append(t, v)
	}
	sort.Slice(t, func(i, j int) bool { return t[i][1] > t[j][1] })

	for i := 0; i < k; i++ {
		res = append(res, t[i][0])
	}
	return res
}

func findKthLargest2(nums []int, k int) int {

	quick(nums, 0, len(nums)-1)
	return 0
}

func quick(nums []int, m, n int) {
	if m >= n {
		return
	}

	mid := nums[(m+n)/2]
	l, r := m, n
	for {
		for nums[l] < mid {
			l++
		}
		for nums[r] > mid {
			r--
		}
		if l >= r {
			break
		}
		nums[l], nums[r] = nums[r], nums[l]
	}

	quick(nums, m, l-1)
	quick(nums, r+1, n)
}

// todo 215
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// todo 739
func dailyTemperatures2(temperatures []int) []int {
	stack := []int{}
	res := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {

		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			k := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[k] = i - k
		}

		stack = append(stack, i)
	}

	return res
}

func dailyTemperatures(temperatures []int) []int {
	res := []int{}

	for i := 0; i < len(temperatures)-1; i++ {
		n := temperatures[i]
		k := 1
		j := i + 1
		for ; j < len(temperatures); j++ {
			if n < temperatures[j] {
				break
			}
			k++
		}
		if j == len(temperatures) {
			k = 0
		}

		res = append(res, k)
	}
	res = append(res, 0)
	return res
}

func decodeString(s string) string {
	reg, _ := regexp.Compile(`(\d+)\[(\w+)\]`)

	r1 := reg.FindAllStringSubmatch(s, -1)
	r2 := reg.FindAllStringIndex(s, -1)
	if len(r2) == 0 {
		return s
	}

	res := ""
	i := 0
	for k, v := range r2 {
		res += s[i:v[0]]
		i = v[1]
		n, _ := strconv.ParseInt(r1[k][1], 10, 64)
		res += strings.Repeat(r1[k][2], int(n))
	}
	res += s[i:]
	return decodeString(res)
}

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor3() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func isValid(s string) bool {
	t := make([]rune, 0, len(s))
	m := map[rune]bool{
		rune('('): true,
		rune('['): true,
		rune('{'): true,
	}
	mp := map[rune]rune{
		rune(')'): rune('('),
		rune(']'): rune('['),
		rune('}'): rune('{'),
	}
	for _, v := range s {
		if m[v] {
			t = append(t, v)
			continue
		}
		if len(t) == 0 {
			return false
		}
		if t[len(t)-1] == mp[v] {
			t = t[:len(t)-1]
			continue
		}
		return false
	}

	return len(t) == 0
}

func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}

	return findM(nums, 0, len(nums))
}

func findM(nums []int, m, n int) int {
	if nums[m] <= nums[n-1] {
		return nums[m]
	}
	t := (m + n) / 2
	r0 := findM(nums, m, t)
	r1 := findM(nums, t, n)
	if r0 > r1 {
		return r1
	}
	return r0
}

func search2(nums []int, target int) int {
	if r, ok := sear(nums, target, 0, len(nums)); ok {
		return r
	}
	return -1
}

func sear(nums []int, target, m, n int) (int, bool) {
	if m >= n || n-m == 1 {
		if nums[m] == target {
			return m, true
		}
		return 0, false
	}

	t := (m + n) / 2
	if nums[t] == target {
		return t, true
	}

	if r, ok := sear(nums, target, m, t); ok {
		return r, true
	}
	if r, ok := sear(nums, target, t, n); ok {
		return r, true
	}

	return 0, false
}

func searchRange(nums []int, target int) []int {
	r0 := searchR(nums, target, 0, len(nums))

	if r0 == len(nums) || nums[r0] != target {
		return []int{-1, -1}
	}

	r1 := searchR(nums, target+1, r0, len(nums))
	return []int{r0, r1 - 1}
}

func searchR(nums []int, target, m, n int) int {
	for m < n {
		t := (m + n) / 2
		if nums[t] < target {
			m = t + 1
		} else {
			n = t
		}
	}
	return m
}

func searchMatrix2(matrix [][]int, target int) bool {
	ll := len(matrix)
	lll := len(matrix[0])
	if matrix[0][0] > target || matrix[ll-1][lll-1] < target {
		return false
	}
	ll0 := 0
	for ; ll0 < ll; ll0++ {
		if matrix[ll0][0] > target {
			break
		}
	}
	ll0--
	for lll0 := 0; lll0 < lll; lll0++ {
		if matrix[ll0][lll0] == target {
			return true
		}
	}

	return false
}

func searchInsert(nums []int, target int) int {
	res := 0
	res = search(nums, 0, len(nums)-1, target)
	return res
}

func search(nums []int, m, n, target int) int {
	if n-m == 1 || m >= n {
		if nums[m] >= target {
			return m
		}
		if nums[n] < target {
			return n + 1
		}
		return n
	}
	t := (m + n) / 2
	if nums[t] > target {
		return search(nums, m, t, target)
	} else if nums[t] < target {
		return search(nums, t, n, target)
	}
	return t
}

func solveNQueens(n int) [][]string {
	res := [][]string{}
	m := [][]string{}
	m0 := []string{}
	for i := 0; i < n; i++ {
		m0 = append(m0, ".")
	}
	for i := 0; i < n; i++ {
		mt := make([]string, len(m0))
		copy(mt, m0)
		m = append(m, mt)
	}

	solv(n, 0, 0, 0, m, &res)
	return res
}

func solv(n, i, j, k int, m [][]string, res *[][]string) {
	if i == n {
		mm := make([]string, 0)
		for t := 0; t < n; t++ {
			mm = append(mm, strings.Join(m[t], ""))
		}
		*res = append(*res, mm)
		return
	}

	for ; j < n; j++ {
		for ; k < n; k++ {
			if m[j][k] == "Q" {
				continue
			}
			m[j][k] = "Q"
			if isQueenOk(m, j, k) {
				solv(n, i+1, j, k, m, res)
			}
			m[j][k] = "."
		}
		k = 0
	}
}
func isQueenOk(m [][]string, j, k int) bool {
	for i := 0; i < len(m); i++ {
		for l := 0; l < len(m); l++ {
			if m[i][l] == "Q" {
				if i == j && l == k {
					continue
				}
				if i == j || l == k {
					return false
				}
				tx := i - j
				if tx < 0 {
					tx = -tx
				}
				ty := l - k
				if ty < 0 {
					ty = -ty
				}

				if tx == ty {
					return false
				}
			}
		}
	}

	return true
}

func partition(s string) [][]string {
	res := [][]string{}
	ss := strings.Split(s, "")
	part(ss, []string{}, &res, 0)
	return res
}

func part(ss []string, t []string, res *[][]string, i int) {
	if i == len(ss) {
		tt := make([]string, len(t))
		copy(tt, t)
		*res = append(*res, tt)
		return
	}

	for j := i + 1; j <= len(ss); j++ {
		if isroll(ss[i:j]) {
			t = append(t, strings.Join(ss[i:j], ""))
			part(ss, t, res, j)
			t = t[:len(t)-1]
		}
	}
}

func isroll(ts []string) bool {
	if len(ts) == 0 {
		return false
	}
	for m := 0; m < len(ts)-1; m++ {
		if ts[m] != ts[len(ts)-1-m] {
			return false
		}
	}
	return true
}

func exist(board [][]byte, word string) bool {
	m := make([][]int, 0)
	for i := 0; i < len(board); i++ {
		m = append(m, make([]int, len(board[0])))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != word[0] {
				continue
			}

			if exis(board, m, word, i, j, 0) {
				return true
			}

		}
	}

	return false
}

func exis(board [][]byte, m [][]int, word string, i, j, t int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}
	if m[i][j] == 1 {
		return false
	}

	if board[i][j] == word[t] {
		if t == len(word)-1 {
			return true
		}
		m[i][j] = 1
		f := exis(board, m, word, i+1, j, t+1) ||
			exis(board, m, word, i, j+1, t+1) ||
			exis(board, m, word, i-1, j, t+1) ||
			exis(board, m, word, i, j-1, t+1)
		if f {
			return true
		}
	}
	m[i][j] = 0
	return false
}

func generateParenthesis(n int) []string {
	res := []string{}

	gene(&res, []string{"(", ")"}, n, []string{})
	return res
}

func gene(res *[]string, r []string, n int, t []string) {
	if len(t) == 1 && t[0] == ")" {
		return
	}
	if len(t) == n*2 {
		s := make([]int, 0)
		for _, v := range t {
			if v == "(" {
				s = append(s, 1)
			} else {
				if len(s) == 0 {
					return
				}
				s = s[:len(s)-1]
			}
		}

		if len(s) != 0 {
			return
		}

		*res = append(*res, strings.Join(t, ""))
		return
	}

	for _, v := range r {
		t = append(t, v)
		gene(res, r, n, t)
		t = t[:len(t)-1]
	}

}
func generateParenthesis2(n int) []string {
	res := []string{}

	gene2(&res, []string{"(", ")"}, 0, n, []string{})
	return res
}
func gene2(res *[]string, r []string, l, n int, t []string) {
	if len(t) == 1 && t[0] == ")" {
		return
	}
	if len(t) == n*2 {
		s := make([]int, 0)
		for _, v := range t {
			if v == "(" {
				s = append(s, 1)
			} else {
				if len(s) == 0 {
					return
				}
				s = s[:len(s)-1]
			}
		}

		if len(s) != 0 {
			return
		}

		*res = append(*res, strings.Join(t, ""))
		return
	}

	for _, v := range r {
		if v == "(" {
			l++
			if l > n {
				continue
			}
		} else {
			l--
			if l < 0 {
				continue
			}
		}
		t = append(t, v)
		gene2(res, r, l, n, t)
		t = t[:len(t)-1]
	}

}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	comb(candidates, target, 0, 0, 0, []int{}, &res)
	return res
}

func comb(candidates []int, target, l, n, sum int, t []int, res *[][]int) {
	if n >= 150 || target < sum {
		return
	}
	if target == sum {
		tt := make([]int, len(t))
		copy(tt, t)
		*res = append(*res, tt)
		return
	}

	for ; l < len(candidates); l++ {
		t = append(t, candidates[l])
		comb(candidates, target, l, n+1, sum+candidates[l], t, res)
		t = (t)[:len(t)-1]
	}

}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	res := []string{}
	ll := len(digits)
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	lett(m, digits, &res, []string{}, 0, ll)
	return res
}

func lett(m map[byte][]string, digits string, res *[]string, t []string, l, ll int) {
	if l == ll {
		tt := strings.Join(t, "")
		*res = append(*res, tt)
		return
	}

	for _, w := range m[digits[l]] {
		t = append(t, w)
		lett(m, digits, res, t, l+1, ll)
		t = (t)[:len(t)-1]
	}
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	subs(nums, &res, []int{}, 0)

	return res
}

func subs(nums []int, res *[][]int, t []int, l int) {
	if len(nums) == l {
		tt := make([]int, len(t))
		copy(tt, t)
		*res = append(*res, tt)
		return
	}

	t = append(t, nums[l])
	subs(nums, res, t, l+1)
	t = (t)[:len(t)-1]
	subs(nums, res, t, l+1)
}

func permute(nums []int) [][]int {
	res := [][]int{}
	m := make(map[int]bool)
	perm(&res, m, nums, []int{}, 0, len(nums))
	return res
}

func perm(res *[][]int, m map[int]bool, nums, t []int, l, ll int) {
	if ll == l {
		*res = append(*res, t)
		return
	}

	for _, v := range nums {
		if m[v] {
			continue
		}
		tt := make([]int, len(t))
		copy(tt, t)
		tt = append(tt, v)

		m[v] = true
		perm(res, m, nums, tt, len(tt), ll)
		m[v] = false
	}
}

type Trie struct {
	M map[rune]*Trie
	C rune
	E bool
}

func Constructor2() Trie {
	return Trie{
		M: make(map[rune]*Trie),
		C: 0,
	}
}

func (this *Trie) Insert(word string) {
	for _, v := range word {
		if _, ok := this.M[v]; !ok {
			this.M[v] = &Trie{
				M: make(map[rune]*Trie),
				C: v,
			}
		}
		this = this.M[v]
	}
	this.E = true
}

func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if _, ok := this.M[v]; !ok {
			return false
		}
		this = this.M[v]
	}

	return this.E
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if _, ok := this.M[v]; !ok {
			return false
		}
		this = this.M[v]
	}

	return true
}

// todo 207
func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	m := make(map[int][]int)
	for _, v := range prerequisites {
		if _, ok := m[v[0]]; !ok {
			m[v[0]] = []int{}
		}
		m[v[0]] = append(m[v[0]], v[1])
		if _, ok := m[v[1]]; !ok {
			m[v[1]] = []int{}
		}
	}

	for i := 0; i < numCourses; i++ {
		if _, ok := m[i]; !ok {
			continue
		}
		if len(m[i]) == 0 {
			continue
		}

		d := map[int]int{i: 1}
		f := false
		canf(m, d, m[i], &f, numCourses)
		if f {
			return false
		}
	}

	return true
}

func canf(m map[int][]int, d map[int]int, l []int, f *bool, num int) {
	for _, j := range l {
		for {
			if _, ok := d[j]; ok {
				d[j]++
				if d[j] >= num {
					*f = true
					return
				} else {
					break
				}

			} else {
				d[j] = 1
				canf(m, d, m[j], f, num)
				break
			}
		}
	}
}

func orangesRotting(grid [][]int) int {
	lll := len(grid)
	ll := len(grid[0])

	bad := [][]int{}
	for i := 0; i < lll; i++ {
		for j := 0; j < ll; j++ {
			if grid[i][j] == 2 {
				bad = append(bad, []int{i, j})
			}
		}
	}

	res := oranges(grid, ll, lll, bad)

	for i := 0; i < lll; i++ {
		for j := 0; j < ll; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return res
}

func oranges(grid [][]int, ll, lll int, bad [][]int) int {
	res := 0
	for {
		t := make([][]int, 0)
		for _, v := range bad {
			i, j := v[0], v[1]
			orange(grid, i, j+1, ll, lll, &t)
			orange(grid, i+1, j, ll, lll, &t)
			orange(grid, i, j-1, ll, lll, &t)
			orange(grid, i-1, j, ll, lll, &t)

			grid[i][j] = 0
		}

		if 0 == len(t) {
			return res
		}
		bad = t
		res++
	}
}

func orange(grid [][]int, i, j, ll, lll int, bad *[][]int) {
	if i < 0 || j < 0 || i >= lll || j >= ll || (grid[i][j] != 1) {
		return
	}
	grid[i][j] = 2
	*bad = append(*bad, []int{i, j})
}

// todo 200
func numIslands(grid [][]byte) int {
	res := 0
	lll := len(grid)
	ll := len(grid[0])

	for i := 0; i < lll; i++ {
		for j := 0; j < ll; j++ {
			if grid[i][j] == '1' {
				res++
				clean(grid, i, j, ll, lll)
			}
		}
	}

	return res
}

func clean(grid [][]byte, i, j, ll, lll int) {
	if i < 0 || j < 0 || i >= lll || j >= ll || (grid[i][j] == '0') {
		return
	}

	grid[i][j] = '0'

	clean(grid, i, j+1, ll, lll)
	clean(grid, i+1, j, ll, lll)
	clean(grid, i, j-1, ll, lll)
	clean(grid, i-1, j, ll, lll)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	m := make(map[*TreeNode]*TreeNode)

	r := root
	lowe(root, root, m)

	pp := p
	for {
		t := q
		for {
			if pp == t {
				return t
			}
			if t == r {
				break
			}

			t = m[t]

		}
		if pp == r {
			break
		}
		pp = m[pp]
	}

	return nil
}

func lowe(root, pre *TreeNode, m map[*TreeNode]*TreeNode) {
	if root == nil {
		return
	}

	m[root] = pre

	lowe(root.Left, root, m)
	lowe(root.Right, root, m)
}

func pathSum(root *TreeNode, targetSum int) int {
	r := []int{}
	return pat(root, r, targetSum)
}

func pat(root *TreeNode, r []int, targetSum int) int {
	if root == nil {
		return 0
	}

	res := 0
	sum := 0
	(r) = append((r), root.Val)
	for i := len(r) - 1; i >= 0; i-- {
		sum += (r)[i]
		if targetSum == sum {
			res++
		}
	}

	l := make([]int, len(r))
	copy(l, r)

	return pat(root.Left, l, targetSum) + pat(root.Right, r, targetSum) + res
}

// todo 105
func buildTree(preorder []int, inorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	return nil
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	r := k(root)
	i := 0
	for ; i < len(r)-1; i++ {
		r[i].Right = r[i+1]
		r[i].Left = nil
	}
	r[len(r)-1].Left = nil
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	l1 := []*TreeNode{root}
	l2 := []*TreeNode{}
	l1f := true
	for len(l1) != 0 || len(l2) != 0 {
		var l *[]*TreeNode
		var lc *[]*TreeNode

		if l1f {
			l = &l1
			lc = &l2
		} else {
			l = &l2
			lc = &l1
		}

		root = (*l)[0]
		ll := len(*l)
		if ll == 1 {
			res = append(res, root.Val)
			l1f = !l1f
		}
		if ll > 0 {
			(*l) = (*l)[1:]
		}

		if root.Left != nil {
			(*lc) = append((*lc), root.Left)
		}
		if root.Right != nil {
			(*lc) = append((*lc), root.Right)
		}
	}

	return res
}

func k(root *TreeNode) []*TreeNode {
	nl := []*TreeNode{}
	res := []*TreeNode{}
	for root != nil || len(nl) != 0 {
		for root != nil {
			res = append(res, root)

			nl = append(nl, root)
			root = root.Left
		}

		root = nl[len(nl)-1]
		nl = nl[:len(nl)-1]

		root = root.Right
	}

	return res
}

func kthSmallest2(root *TreeNode, k int) int {
	nl := []*TreeNode{root}
	for root != nil || len(nl) != 0 {
		for root != nil {
			nl = append(nl, root)
			root = root.Left
		}

		root = nl[len(nl)-1]
		nl = nl[:len(nl)-1]

		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}

	return 0
}

func kthSmallest(root *TreeNode, k int) int {
	l := inorderTraversal(root)

	return l[k-1]
}

func isValidBST(root *TreeNode) bool {
	return isvl(root.Left, math.MinInt64, root.Val) && isvr(root.Right, root.Val, math.MaxInt64)
}

func isvl(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return isvl(root.Left, min, root.Val) && isvr(root.Right, root.Val, max)
}
func isvr(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}

	return isvl(root.Left, min, root.Val) && isvr(root.Right, root.Val, max)
}

func sortedArrayToBST(nums []int) *TreeNode {
	return sor(nums, 0, len(nums)-1)
}

func sor(nums []int, i, j int) *TreeNode {
	if j < i {
		return nil
	}
	k := (i + j) / 2

	root := &TreeNode{Val: nums[k]}

	root.Left = sor(nums, i, k-1)
	root.Right = sor(nums, k+1, j)

	return root
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	lev(root, 0, &res)
	return res
}

func lev(root *TreeNode, l int, res *[][]int) {
	if root == nil {
		return
	}

	ll := len(*res)
	if l > ll-1 {
		*res = append(*res, []int{})
	}

	(*res)[l] = append((*res)[l], root.Val)
	lev(root.Left, l+1, res)
	lev(root.Right, l+1, res)
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	res := dia(root, &max)
	if res < max {
		res = max
	}

	return res
}

func dia(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	l := 0
	if root.Left != nil {
		l = 1
		l += dia(root.Left, max)
	}

	r := 0
	if root.Right != nil {
		r = 1
		r += dia(root.Right, max)
	}

	res := l
	if r > l {
		res = r
	}

	if l+r > *max {
		*max = l + r
	}

	return res
}

func isSymmetric(root *TreeNode) bool {
	l := isl(root.Left)
	r := isr(root.Right)

	if len(l) != len(r) {
		return false
	}

	for k, v := range r {
		if l[k] != v {
			return false
		}
	}

	return true
}

func isr(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	if root.Right != nil {
		res = append(res, root.Right.Val)
		res = append(res, isr(root.Right)...)
	} else {
		res = append(res, -9999)
	}

	res = append(res, root.Val)
	if root.Left != nil {
		res = append(res, root.Left.Val)
		res = append(res, isr(root.Left)...)
	} else {
		res = append(res, -9999)
	}

	return res
}

func isl(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	if root.Left != nil {
		res = append(res, root.Left.Val)
		res = append(res, isl(root.Left)...)
	} else {
		res = append(res, -9999)
	}
	res = append(res, root.Val)

	if root.Right != nil {
		res = append(res, root.Right.Val)
		res = append(res, isl(root.Right)...)
	} else {
		res = append(res, -9999)
	}

	return res
}

func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

func invert(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	invert(root.Left)
	invert(root.Right)
}

func maxDepth(root *TreeNode) int {
	res := 0
	sd(root, &res)
	return res
}

func sd(root *TreeNode, n *int) {
	if root == nil {
		return
	}
	*n++

	l := *n
	r := *n
	if root.Left != nil {
		sd(root.Left, &l)
	}

	if root.Right != nil {
		sd(root.Right, &r)
	}

	if l > r {
		*n = l
	} else {
		*n = r
	}

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	l := make([]*TreeNode, 0)

	for root != nil || len(l) != 0 {
		for root != nil {
			l = append(l, root)
			root = root.Left
		}

		root = l[len(l)-1]
		l = l[:len(l)-1]
		res = append(res, root.Val)
		root = root.Right
	}

	return res
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}

type LRUCacheNode struct {
	val       int
	key       int
	next, pre *LRUCacheNode
}
type LRUCache struct {
	m          map[int]*LRUCacheNode
	cap        int
	head, tail *LRUCacheNode
}

func Constructor(capacity int) LRUCache {
	l := &LRUCacheNode{}
	r := &LRUCacheNode{}
	l.next = r
	r.pre = l
	return LRUCache{
		m:    make(map[int]*LRUCacheNode),
		cap:  capacity,
		head: l,
		tail: r,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}

	cur := this.m[key]
	cur.next.pre = cur.pre
	cur.pre.next = cur.next
	cur.pre = this.head
	cur.next = this.head.next
	this.head.next.pre = cur
	this.head.next = cur

	return this.m[key].val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.m[key]; !ok {
		cur := &LRUCacheNode{val: value, key: key}
		this.head.next.pre = cur
		cur.next = this.head.next
		cur.pre = this.head
		this.head.next = cur

		this.m[key] = cur
	} else {
		this.m[key].val = value

		cur := this.m[key]
		cur.next.pre = cur.pre
		cur.pre.next = cur.next
		cur.pre = this.head
		cur.next = this.head.next
		this.head.next.pre = cur
		this.head.next = cur
	}
	if len(this.m) > this.cap {
		k := this.tail.pre.key
		cur := this.tail.pre
		cur.pre.next = cur.next
		cur.next.pre = cur.pre
		delete(this.m, k)
	}
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	l := make([]*ListNode, 0)
	for head != nil {
		l = append(l, head)

		head = head.Next
	}

	sort.Slice(l, func(i, j int) bool {
		return l[i].Val < l[j].Val
	})

	ll := len(l)
	i := 0
	for ; i < ll-1; i++ {
		l[i].Next = l[i+1]
	}
	l[i].Next = nil
	return l[0]
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var res *Node
	var pre *Node
	h := head
	m := make(map[*Node]*Node, 0)
	for head != nil {
		n := &Node{Val: head.Val}

		if pre == nil {
			res = n
		} else {
			pre.Next = n
		}

		m[head] = n
		pre = n
		head = head.Next
	}

	for h != nil {
		m[h].Random = m[h.Random]
		h = h.Next
	}

	return res
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	res := head.Next
	var pre *ListNode

	for {
		if head == nil || head.Next == nil {
			break
		}

		next := head.Next
		head.Next = head.Next.Next
		next.Next = head
		if pre != nil {
			pre.Next = next
		}

		pre = head
		head = head.Next
	}

	return res
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := make([]*ListNode, 0)

	for head != nil {
		l = append(l, head)
		head = head.Next
	}

	ll := len(l)
	m := ll - n
	if m == 0 {
		if ll > 1 {
			return l[1]
		}
		return nil
	}

	l[m-1].Next = l[m].Next

	return l[0]
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	add1 := false
	r := l3
	var pre *ListNode
	pre = l3
	for {
		l1val := 0
		l2val := 0
		if l1 != nil {
			l1val = l1.Val
		}
		if l2 != nil {
			l2val = l2.Val
		}

		if l1 == nil && l2 == nil {
			if add1 {
				l3.Val = 1
			} else {
				pre.Next = nil
			}
			return r
		}

		l3.Val = l1val + l2val
		if add1 {
			l3.Val++
		}
		add1 = false
		if l3.Val > 9 {
			l3.Val -= 10
			add1 = true
		}
		pre = l3
		v := &ListNode{}
		l3.Next = v
		l3 = v

		if l2 != nil {
			l2 = l2.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}

	}
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var pre *ListNode
	ll1 := 0
	ll2 := 0
	for l1 != nil {
		next := l1.Next
		l1.Next = pre
		l1, pre = next, l1
		ll1++
	}
	pre, l1 = l1, pre

	for l2 != nil {
		next := l2.Next
		l2.Next = pre
		l2, pre = next, l2
		ll2++
	}
	l2 = pre

	if ll2 > ll1 {
		l1, l2 = l2, l1
	}

	add1 := false
	r := l1
	for {
		if l1 == nil {
			pre = nil
			for r != nil {
				next := r.Next
				r.Next = pre
				r, pre = next, r
			}
			if add1 {
				return &ListNode{Val: 1, Next: pre}
			}
			return pre
		}

		l1.Val += l2.Val
		if add1 {
			l1.Val++
		}
		add1 = false
		if l1.Val > 9 {
			l1.Val -= 10
			add1 = true
		}

		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}

		l1 = l1.Next
	}
}

func rev(l1 *ListNode) *ListNode {
	if l1 == nil || l1.Next == nil {
		return l1
	} else {
		t := rev(l1.Next)
		l1.Next.Next = l1
		l1.Next = nil
		return t
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var res *ListNode
	var v *ListNode
	var t *ListNode
	if list1.Val < list2.Val {
		res = list1
	} else {
		res = list2
	}
	t = res

	for {
		if list1 == nil && list2 == nil {
			break
		}
		if list1 == nil && list2 != nil {
			t.Next = list2
			break
		} else if list1 != nil && list2 == nil {
			t.Next = list1
			break
		} else {
			if list1.Val < list2.Val {
				v = list1
				list1 = list1.Next
			} else {
				v = list2
				list2 = list2.Next
			}
		}

		t.Next = v
		t = t.Next
	}

	return res
}

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]int, 0)
	for {
		if head == nil {
			break
		}
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = 0
		head = head.Next
	}

	return nil
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	i := 0
	m := head
	n := head.Next
	for {
		if n == nil {
			return false
		}
		if m == n {
			return true
		}
		n = n.Next

		if i%2 == 0 {
			m = m.Next
		}
		i++
	}
}

func isPalindrome(head *ListNode) bool {
	l := make([]int, 0)
	for head != nil {
		l = append(l, head.Val)
		head = head.Next
	}

	ll := len(l)
	for i := 0; i < ll/2; i++ {
		if l[i] != l[ll-1-i] {
			return false
		}
	}

	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]int, 0)

	for {
		m[headA] = 0
		if headA.Next == nil {
			break
		}
		headA = headA.Next
	}

	for {
		if _, ok := m[headB]; ok {
			return headB
		}
		if headB.Next == nil {
			break
		}
		headB = headB.Next
	}

	return nil
}

func searchMatrix(matrix [][]int, target int) bool {
	ll := len(matrix)
	lll := len(matrix[0])
	ll0, lll0 := 0, 0

	if target < matrix[0][0] || target > matrix[ll-1][lll-1] {
		return false
	}

	for ; ll0 < ll-1; ll0++ {
		if target < matrix[ll0][0] {
			break
		} else if target == matrix[ll0][0] {
			return true
		}
	}

	for ; ll0 >= 0; ll0-- {
		for ; lll0 < lll; lll0++ {
			if target < matrix[ll0][lll0] {
				break
			} else if target == matrix[ll0][lll0] {
				return true
			}
		}
	}

	return false
}

func rotate48(matrix [][]int) {
	ll := len(matrix)

	for i := 0; i < ll/2; i++ {
		for j := 0; j < (ll+1)/2; j++ {
			matrix[i][j], matrix[ll-j-1][i], matrix[ll-i-1][ll-j-1], matrix[j][ll-i-1] =
				matrix[ll-j-1][i], matrix[ll-i-1][ll-j-1], matrix[j][ll-i-1], matrix[i][j]
		}
	}
}

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)

	ll := len(matrix)
	ll0 := 0
	lll := len(matrix[0])
	lll0 := 0
	var r []int
	for {
		r = row(matrix, ll0, lll0, lll)
		res = append(res, r...)

		ll0++
		if ll0 >= ll {
			break
		}

		r = line(matrix, lll-1, ll0, ll)
		res = append(res, r...)

		lll--
		if lll0 >= lll {
			break
		}

		r = row(matrix, ll-1, lll-1, lll0)
		res = append(res, r...)

		ll--
		if ll0 >= ll {
			break
		}

		r = line(matrix, lll0, ll-1, ll0)
		res = append(res, r...)

		lll0++
		if lll0 >= lll {
			break
		}
	}

	return res
}

func row(matrix [][]int, l, s, e int) []int {
	res := make([]int, 0)

	if s >= e {
		for i := s; i >= e; i-- {
			res = append(res, matrix[l][i])
		}
	} else {
		for i := s; i < e; i++ {
			res = append(res, matrix[l][i])
		}
	}

	return res
}

func line(matrix [][]int, r, s, e int) []int {
	res := make([]int, 0)

	if s >= e {
		for i := s; i >= e; i-- {
			res = append(res, matrix[i][r])
		}
	} else {
		for i := s; i < e; i++ {
			res = append(res, matrix[i][r])
		}
	}

	return res
}

func setZeroes(matrix [][]int) {
	ll := len(matrix)
	lll := len(matrix[0])
	zl := make([]int, lll)
	zr := make([]int, ll)
	for i, l := range matrix {
		for j, v := range l {
			if v == 0 {
				zl[j]++
				zr[i]++
			}
		}
		if zr[i] > 0 {
			for j := range l {
				l[j] = 0
			}
		}
	}

	for k, v := range zl {
		if v > 0 {
			for _, l := range matrix {
				l[k] = 0
			}
		}
	}
}

func productExceptSelf(nums []int) []int {
	ll := len(nums)
	res := make([]int, len(nums))

	l := make([]int, len(nums))
	r := make([]int, len(nums))

	l[0] = 1
	for i := 0; i < ll-1; i++ {
		l[i+1] = l[i] * nums[i]
	}

	r[ll-1] = 1
	for i := ll - 1; i > 0; i-- {
		r[i-1] = r[i] * nums[i]
	}

	for k, v := range l {
		res[k] = v * r[k]
	}

	return res
}

func rotate(nums []int, k int) {
	ll := len(nums)
	k = k % ll

	t := append(nums[ll-k:], nums[:ll-k]...)
	copy(nums, t)
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	res = append(res, intervals[0])

	t := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[t][1] {
			res = append(res, intervals[i])
			t++
			continue
		}
		if intervals[i][0] <= res[t][1] && intervals[i][1] >= res[t][1] {
			res[t][1] = intervals[i][1]
			continue
		}
	}

	return res
}

// todo:53
func maxSubArray(nums []int) int {
	max := nums[0]

	return max
}

func subarraySum(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
	}

	return res
}

// todo:438
func findAnagrams2(s string, p string) []int {
	pm := make(map[byte]int)
	for _, v := range p {
		pm[byte(v)] = 0
	}

	res := make([]int, 0)
	i, j := 0, 0
	ll := len(s)
	lp := len(pm)
	t := 0
	for ; j < ll; j++ {
		if v, ok := pm[s[j]]; ok {
			if v == 0 {
				pm[s[j]]++
				t++
				if t == lp {
					res = append(res, i)
					pm[s[i]] = 0
					i++
					t = 0
				}
			} else {
				pm[s[i]] = 0
				i++
				t = 0
			}
		} else {
			i = j
			for k := range pm {
				pm[k] = 0
			}
		}
	}

	return res
}

func lengthOfLongestSubstring(s string) int {
	ll := len(s)
	if ll == 1 {
		return 1
	} else if ll == 0 {
		return 0
	}

	res := 0
	t := 0
	m := make(map[byte]int, ll)
	m[s[0]] = 0

	for i, j := 0, 1; j < ll; j++ {
		if n, ok := m[s[j]]; ok {
			t = len(m)
			if res < t {
				res = t
			}
			fmt.Println(s[i:j])
			i = j + 1
			for k, v := range m {
				if v < n {
					delete(m, k)
				}
			}
		}
		m[s[j]] = j
	}

	if res < len(m) {
		res = len(m)
	}
	return res
}

func trap(height []int) int {
	l := make([]int, len(height))
	r := make([]int, len(height))
	ll := len(height)
	l[0] = height[0]
	r[ll-1] = height[ll-1]
	for i := 1; i < ll; i++ {
		l[i] = max(height[i], l[i-1])
	}

	for i := ll - 2; i >= 0; i-- {
		r[i] = max(height[i], r[i+1])
	}

	res := 0
	for i, v := range height {
		res += min(l[i], r[i]) - v
	}
	return res
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	ll := len(nums)

	for i := 0; i < ll; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		k := ll - 1
		for j := i + 1; j < ll; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}

			for k > j && nums[k]+nums[j]+nums[i] > 0 {
				k--
			}

			if k == j {
				break
			}
			if nums[k]+nums[j]+nums[i] == 0 {
				res = append(res, []int{nums[k], nums[j], nums[i]})
			}
		}
	}

	return res
}

func maxArea(height []int) int {
	res := 0
	t := 0
	ll := len(height) - 1
	m := 0
	for k, v := range height {
		if m > v {
			continue
		}
		for i := k + 1; i <= ll; i++ {
			if v < height[i] {
				t = v * (i - k)
			} else {
				t = height[i] * (i - k)
			}

			if t > res {
				res = t
			}
		}
		if m < v {
			m = v
		}
	}

	return res
}

func moveZeroes2(nums []int) {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != 0 {
			nums[slowIndex], nums[fastIndex] = nums[fastIndex], nums[slowIndex]
			slowIndex++
		}
	}
}

func moveZeroes(nums []int) {
	ll := len(nums) - 1
	for k, v := range nums {
		if v != 0 {
			continue
		}
		for i := k + 1; i <= ll; i++ {
			if nums[i] == 0 {
				continue
			}
			nums[k], nums[i] = nums[i], nums[k]
			break
		}
	}
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tnums := make(map[int]bool, len(nums))
	for _, v := range nums {
		tnums[v] = true
	}

	tl := 1
	res := 1

	for v := range tnums {
		if tnums[v-1] {
			continue
		}
		for tnums[v+1] {
			tl++
			v++
		}
		if tl >= res {
			res = tl
		}
		tl = 1
	}

	return res
}

func longestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tnums := make(map[int]bool, len(nums))
	for _, v := range nums {
		tnums[v] = true
	}
	nums = make([]int, 0, len(tnums))
	for k := range tnums {
		nums = append(nums, k)
	}

	sort.Ints(nums)
	t := 1
	l := len(nums)
	res := 1
	for k, v := range nums {
		p := k + 1
		if p >= l {
			if res < t {
				res = t
			}
			break
		}
		if v+1 == nums[p] {
			t++
		} else {
			if res < t {
				res = t
			}
			t = 1
		}
	}

	return res
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}

	res := make(map[string][]string, 0)

	for _, str := range strs {
		ha := []rune(str)
		sort.Slice(ha, func(i, j int) bool { return ha[i] > ha[j] })
		hs := string(ha)

		if _, ok := res[hs]; !ok {
			res[hs] = make([]string, 0)
		}
		res[hs] = append(res[hs], str)
	}

	rr := make([][]string, 0, len(res))
	for _, v := range res {
		rr = append(rr, v)
	}

	return rr
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if v1+v2 == target {
				res[0] = k1
				res[1] = k2
				return res
			}
		}
	}

	return res
}
func twoSum2(nums []int, target int) []int {
	res := make([]int, 2)
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		for j := i + 1; j < numsLen; j++ {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}

	return res
}
