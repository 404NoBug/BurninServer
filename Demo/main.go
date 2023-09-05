package main

import (
	"fmt"
	"sort"
	"sync"
)

var (
	Cnt int
	mu  sync.Mutex
)

func Add(iter int) {
	mu.Lock()
	for i := 0; i < iter; i++ {
		Cnt++
	}
	mu.Unlock()
}

type MaxHeap []int

func main() {
	// 1. 创建trace持久化的文件句柄
	//f, err := os.Create("trace.out")
	//if err != nil {
	//	log.Fatalf("failed to create trace output file: %v", err)
	//}
	//defer func() {
	//	if err := f.Close(); err != nil {
	//		log.Fatalf("failed to close trace file: %v", err)
	//	}
	//}()
	//// 2. trace绑定文件句柄
	//if err := trace.Start(f); err != nil {
	//	log.Fatalf("failed to start trace: %v", err)
	//}
	//defer trace.Stop()
	//
	//wg := &sync.WaitGroup{}
	//for i := 0; i < 2; i++ {
	//	wg.Add(1)
	//	go func() {
	//		Add(100000)
	//		defer wg.Done()
	//	}()
	//}
	//wg.Wait()
	//fmt.Println(Cnt)
	//wg := &sync.WaitGroup{}
	//wg.Add(2)
	//go func() {
	//	fmt.Println("go run")
	//	wg.Done()
	//}()
	//wg.Wait()
	//fmt.Println(twoSum2([]int{1, 6, 5, 7, 9}, 10))
	//fmt.Println(findMedianSortedArrays1([]int{1, 3, 5}, []int{2, 4, 6})) // 3.5
	//fmt.Println(maxArea1([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	//fmt.Println(maxArea2([]int{1, 8, 6, 2, 9, 4, 8, 3, 7})) // 49
	//fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1 -1 2] [-1 0 1]]
	//fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	//fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	//fmt.Println(removeDuplicates1([]int{1, 1, 2}))                      // 2
	//fmt.Println(removeDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})) // 5
	//fmt.Println(removeElement2([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))

	//nums := []int{1, 2, 7, 4, 3, 1}
	//nextPermutation(nums)
	//fmt.Println(nums) // [1 3 1 2 4 7] // bingo

	//fmt.Println(reverse([]int{1, 2, 3, 4, 5, 6}))

	//fmt.Println(searchRange2([]int{5, 7, 7, 8, 8, 10}, 8)) // [3, 4]

	//fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	//fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	//fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))

	//fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))

	//fmt.Println(firstMissingPositive([]int{3, 4, -1, 1})) // 2
	//fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6

	//matrix := [][]int{
	//	{5, 1, 9, 11},
	//	{2, 4, 8, 10},
	//	{13, 3, 6, 7},
	//	{15, 14, 12, 16},
	//}
	//matrix = rotate(matrix)
	//for _, nums := range matrix {
	//	for _, num := range nums {
	//		fmt.Print(num, " ")
	//	}
	//	fmt.Println()
	//}
	//fmt.Println(maxSubArray([]int{-3, 4, -1, 3, -5}))
	//n := 10
	//depth := 0
	//for i := n; i > 0; i >>= 1 {
	//	fmt.Println("i = ", i)
	//	depth++
	//}
	//fmt.Println(depth)
	//fmt.Println(spiralOrder([][]int{
	//	// {1, 2, 3, 4},
	//	// {5, 6, 7, 8},
	//	// {9, 10, 11, 12},
	//
	//	// {1, 2},
	//	// {3, 4},
	//
	//	{2, 3, 4},
	//	{5, 6, 7},
	//	{8, 9, 10},
	//	{11, 12, 13},
	//	{14, 15, 16},
	//}))

	//a := [][]int{
	//	{2, 3, 1, 4},
	//	{1, 2, 0, 3},
	//	{4, 2, 1, 7},
	//	{3, 1, 4, 2},
	//}
	//maxCol, maxRow := len(a[0]), len(a)
	//maxColArr := make([]int, maxCol)
	//maxRowArr := make([]int, maxRow)
	//sum := 0
	//for i := 0; i < maxRow; i++ {
	//	curRow := 0
	//	for j := 0; j < maxCol; j++ {
	//		curRow += a[i][j]
	//	}
	//	fmt.Println(curRow)
	//	maxRowArr[i] = curRow
	//	sum += curRow
	//}
	//for j := 0; j < maxCol; j++ {
	//	curCol := 0
	//	for i := 0; i < maxRow; i++ {
	//		curCol += a[i][j]
	//	}
	//	fmt.Println(curCol)
	//	maxColArr[j] = curCol
	//}
	//maxNum := 0
	//for i, RowNum := range maxRowArr {
	//	for j, colNum := range maxColArr {
	//		fmt.Println(RowNum, colNum, a[i][j])
	//		maxNum = max(maxNum, RowNum+colNum-a[i][j])
	//	}
	//}
	//fmt.Println(maxNum)

	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4, 7}, 2)) // 5
}

func findKthLargest(nums []int, k int) int {
	maxHeap := NewMaxHeap(nums)
	for k > 1 {
		maxHeap.Pop()
		k--
	}
	return maxHeap.Pop()
}

func NewMaxHeap(nums []int) *MaxHeap {
	hs := MaxHeap(nums)
	n := len(hs)
	h := &hs
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
	return h
}

func (h *MaxHeap) Push(v int) {
	*h = append(*h, v)
	h.up(len(*h) - 1)
}

func (h *MaxHeap) Pop() int {
	hs := *h
	max := hs[0]
	n := len(hs)

	h.swap(0, n-1)
	h.down(0, n-1) // 除最后一个元素外全体下滤
	*h = hs[:n-1]

	return max
}

// 上滤
func (h *MaxHeap) up(i int) {
	for {
		parent := (i - 1) / 2
		if h.more(parent, i) || parent == i {
			break
		}
		h.swap(parent, i)
		i = parent
	}
}

// 下滤
func (h *MaxHeap) down(mid, n int) {
	for {
		l := 2*mid + 1
		if l >= n || l < 0 {
			fmt.Println(*h)
			break
		}
		max := l
		if r := l + 1; r < n && h.more(r, max) {
			max = r
		}
		if !h.more(max, mid) {
			fmt.Println(*h)
			break
		}

		h.swap(mid, max)
		mid = max
	}
}

func (h *MaxHeap) swap(i, j int) {
	hs := *h
	hs[i], hs[j] = hs[j], hs[i]
}

func (h *MaxHeap) more(i, j int) bool {
	hs := *h
	return hs[i] > hs[j]
}

func canJump(nums []int) bool {
	len := len(nums)
	maxPos := 0
	if len <= 1 {
		return true
	}
	for i := 0; i < len-1; i++ {
		if i <= maxPos {
			maxPos = max(maxPos, i+nums[i])
			if maxPos >= len-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return nil
	}
	if len(matrix) == 1 {
		return matrix[0]
	}
	return order(matrix, 0, len(matrix[0])-1, 0, len(matrix)-1, []int{})
}

func order(matrix [][]int, start, end int, up, down int, nums []int) []int {
	if start > end && up > down {
		return nums
	}
	//向右走
	for i := start; i <= end; i++ {
		nums = append(nums, matrix[start][i])
	}

	//向下走
	stop := true
	for i := up + 1; i <= down; i++ {
		nums = append(nums, matrix[i][end])
		stop = false
	}
	if stop {
		return nums
	}

	//向左走
	stop = true
	for i := end - 1; i >= start; i-- {
		nums = append(nums, matrix[down][i])
		stop = false
	}
	if stop {
		return nums
	}

	//向上走
	stop = true
	for i := down - 1; i > up; i-- {
		nums = append(nums, matrix[i][up])
		stop = false
	}
	if stop {
		return nums
	}
	return order(matrix, start+1, end-1, up+1, down-1, nums)
}

//func order(matrix [][]int, start, end int, up, down int, nums []int) []int {
//	if start > end || up > down {
//		return nums
//	}
//	// 向右走
//	for i := start; i <= end; i++ {
//		nums = append(nums, matrix[up][i])
//	}
//
//	// 向下走
//	stop := true
//	for i := up + 1; i <= down; i++ {
//		nums = append(nums, matrix[i][end])
//		stop = false
//	}
//	if stop {
//		return nums // 无路可走
//	}
//
//	// 向左走
//	stop = true
//	for i := end - 1; i >= start; i-- {
//		nums = append(nums, matrix[down][i])
//		stop = false
//	}
//	if stop {
//		return nums // 无路可走
//	}
//
//	// 向上走
//	for i := down - 1; i > up; i-- {
//		if end > 0 { // 单列的情况，下一趟不能往上走
//			nums = append(nums, matrix[i][start])
//		}
//	}
//	return order(matrix, start+1, end-1, up+1, down-1, nums)
//}

//
// 跨步遍历的感觉
//
func maxSubArray(nums []int) int {
	maxNum := nums[0]
	for _, n := range nums {
		if n > maxNum {
			maxNum = n
		}
	}
	sum, maxSum := 0, 0
	for i := 0; i < len(nums); i++ {
		switch {
		case nums[i] > 0:
			sum += nums[i]
		case nums[i] < 0:
			if sum+nums[i] > 0 {
				sum += nums[i]
			} else {
				sum = 0
			}
		}
		maxSum = max(sum, maxSum)
	}
	if maxNum < 0 && maxSum > maxNum {
		return maxNum
	}
	return maxSum
}

// 正方形任意旋转可改为多重折叠
func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	if n <= 1 {
		return matrix
	}
	// 右侧对角线折叠
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	//左右折叠
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
	return matrix
}

func trap(nums []int) int {
	area := 0
	l, r := 0, len(nums)-1
	lTop, rTop := 0, 0
	for l < r {
		lTop = max(lTop, nums[l])
		rTop = max(rTop, nums[r])
		if lTop < rTop {
			area += lTop - nums[l] // 右侧更高些，往右侧走。一边向右移，一边计算高度差来累计面积（最高点处高度差为 0）
			l++
		} else {
			area += rTop - nums[r] // 向左侧移，用目前右侧最高高度来计算高度差，累计面积
			r--
		}
	}
	return area
}

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// 对数组进行归位预处理
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			swap(&nums[i], &nums[nums[i]-1])
		}
	}
	// fmt.Println(nums)    // [1 -1 3 4]

	// 向后检查
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// 理想正整数数组
	return n + 1
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combine(candidates, target)
}

func combine(nums []int, target int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	subSum := target - nums[0]
	switch {
	case subSum < 0: //向后不存在更大的值，递归结束
		return res
	case subSum == 0: //第一个数就是target
		res = append(res, []int{nums[0]})
	case subSum > 0:
		remains := combine(nums, subSum) //寻找所有的子集合
		for _, v := range remains {
			way := append([]int{nums[0]}, v...)
			res = append(res, way)
		}
	}
	res = append(res, combine(nums[1:], target)...) // 向后查找其他组合，避免重复
	return res
}

// 二分查找
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		switch {
		case target < nums[mid]:
			right = mid - 1
		case target > nums[mid]:
			left = mid + 1
		default:
			return mid
		}
	}
	return left
}

// 改进的二分搜索
func searchRange2(nums []int, target int) []int {
	l := edgeBinSearch(nums, true, target)
	r := edgeBinSearch(nums, false, target)
	return []int{l, r}
}

// 普通二分搜索在匹配到 target 时直接 return
// 在本题搜索时在匹配到 target 之后依旧向边缘走当做没匹配到，注意 2 个边界条件
// O(logN) // ok
func edgeBinSearch(nums []int, leftest bool, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		switch {
		case target < nums[mid]:
			r = mid - 1
		case target > nums[mid]:
			l = mid + 1
		default:
			if leftest {
				if mid == 0 || nums[mid] > nums[mid-1] { // 不再继续向左走的 2 个边界条件
					return mid
				}
				r = mid - 1 // 继续在左侧找
			} else {
				if mid == n-1 || nums[mid] < nums[mid+1] {
					return mid
				}
				l = mid + 1 // 继续在右侧找
			}
		}
	}
	return -1
}

// 类二分搜索
// 最左边数 < 中间数则左侧有序，最右边数 > 中间数则右侧有序
// 在缩小搜索区域时，一直只在确定的有序区域内查找
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		switch {
		case nums[mid] == target: // bingo
			return mid
		case nums[l] <= nums[mid]:
			if nums[l] <= target && target < nums[mid] { // 保证 target 一定在有序的左侧内
				r = mid - 1
			} else {
				l = mid + 1
			}
		case nums[r] >= nums[mid]:
			if nums[mid] <= target && target < nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

// 数组规律题
// 从后往前找第一个下降点 i，再从后往前找它的 ceil 值，交换
// 再将 [i+1:] 之后的数据从降序反转为升序，为最小序列
func nextPermutation(nums []int) {
	// 处理降序的 case
	desc := true
	n := len(nums)
	for i := range nums[:n-1] {
		if nums[i] < nums[i+1] {
			desc = false
		}
	}
	if desc {
		reverse(nums)
		return
	}
	// 从后向前找第一个下降的点
	//[1, 2, 7, 3, 4, 1]
	var i int
	for i = n - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			i-- // 找到 2
			break
		}
	}
	fmt.Println(i)
	// [1, 2, 7, 4, 3, 1]
	// 从后向前，找向上最接近的值
	for j := n - 1; j > i; j-- {
		if nums[j] > nums[i] {
			nums[j], nums[i] = nums[i], nums[j] // 交换 2 和 3     // [1 3 7 4 2 1]
			break
		}
	}

	reverse(nums[i+1:]) // 反转7 4 2 1    // [1 3 1 2 4 7]
}

func reverse(nums []int) []int {
	len := len(nums) - 1
	for i := range nums {
		if i < len-i {
			//fmt.Println(nums[i], nums[len-i])
			nums[i], nums[len-i] = nums[len-i], nums[i]
		}
	}
	return nums
}

func removeElement1(nums []int, val int) (int, []int) {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow, nums[:slow]
}

func removeElement2(nums []int, val int) (int, []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			if slow != fast {
				nums[slow] = nums[fast]
			}
			slow++
			fast++
			continue
		}
		fast++
	}
	return slow, nums[:slow]
}

// 针对有序数组，双指针法是十分常见且有用的
func removeDuplicates1(nums []int) (int, []int) {
	slow, fast := 0, 0
	for fast < len(nums)-1 {
		if nums[fast] != nums[fast+1] { //相邻不相等
			slow++
			fast++
			nums[slow] = nums[fast]
			continue
		}
		fast++
	}
	return slow + 1, nums[:slow+1]
}

// 充分利用数组有序的已知条件
func removeDuplicates2(nums []int) (int, []int) {
	n := len(nums)
	l, r := 0, 1
	for r < n {
		if nums[l] < nums[r] { // 比我大就放到我的下一个
			l++
			nums[l], nums[r] = nums[r], nums[l]
		}
		r++
	}
	return l + 1, nums[:l+1]
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	for i := 0; i < n-1; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 去重1
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] { // 去重2    // 注意条件：j>i+1 与 i>0 相同都是为了排除第一个相同数
				continue
			}
			head, tail := j+1, n-1
			for head < tail {
				sum := nums[i] + nums[j] + nums[head] + nums[tail]
				switch {
				case sum < target:
					head++
				case sum > target:
					tail--
				case sum == target: // 向前向后走
					res = append(res, []int{nums[i], nums[j], nums[head], nums[tail]})
					// 去重3：注意 for 循环条件的判断，避开死循环
					for head < tail && nums[head] == nums[head+1] {
						head++
					}
					for head < tail && nums[tail] == nums[tail-1] {
						tail--
					}
					head++
					tail--
				}
			}
		}
	}
	return res
}

// twoSum 的思路，不好
func badTwoSum(nums []int) [][]int {
	// 避开全是 0 的 case     // ugly
	if len(nums) >= 3 {
		allZero := true
		for _, num := range nums {
			if num != 0 {
				allZero = false
			}
		}
		if allZero {
			return [][]int{{0, 0, 0}}
		}
	}

	n := len(nums)
	num2index := make(map[int]int, n)
	for i, num := range nums {
		num2index[num] = i
	}

	// 获取三元组
	var res [][]int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			remain := 0 - (nums[i] + nums[j])
			if k, ok := num2index[remain]; ok && j != k && i != k {
				res = append(res, []int{nums[i], nums[j], remain})
			}
		}
	}

	// 剔除重复的三元组
	m := make(map[string][]int)
	for i := range res {
		sort.Ints(res[i])
		m[intStr(res[i])] = res[i]
	}

	var arrs [][]int
	for _, arr := range m {
		arrs = append(arrs, arr)
	}
	return arrs
}

// trick    // 使得整数数组能做 map 的 key
func intStr(nums []int) string {
	str := ""
	for _, num := range nums {
		str += fmt.Sprintf("%d_", num)
	}
	return str
}

// 同样也是指针法
// 排序预处理能知道双指针移动的方向，记录最小 abs
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	minAbs := 1<<31 - 1
	minSum := 0
	for i, num := range nums {
		// 第一层遍历数向前去重
		if i > 0 && nums[i] == nums[i-1] { // 因为双指针从 i 之后取，不能使用 nums[i] == nums[i+1] 向后去重
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := num + nums[l] + nums[r]
			if abs(target-sum) < minAbs {
				minAbs = abs(target - sum)
				minSum = sum
			}
			switch {
			case sum < target:
				l++
			case sum > target:
				r--
			default:
				return target
			}
		}
	}
	return minSum
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums) //[-4,-1,-1,0,1,2]
	n := len(nums)
	var res [][]int
	for i, num := range nums {
		if num > 0 {
			break // 优化，再往后三个正数和不可能为 0
		}
		// 第一层遍历数向前去重
		if i > 0 && nums[i] == nums[i-1] { // 因为双指针从 i 之后取，不能使用 nums[i] == nums[i+1] 向后去重
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := num + nums[l] + nums[r]
			switch {
			case sum > 0:
				r--
			case sum < 0:
				l++
			default:
				res = append(res, []int{num, nums[l], nums[r]})
				// 第二层候选数向后去重
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for r > l && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return res
}
func maxArea1(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	maxMulti := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			wide := j - i //宽
			hight := min(height[i], height[j])
			area := wide * hight
			fmt.Printf("wide = %d,hight = %d \n", wide, hight)
			maxMulti = max(area, maxMulti)
		}
	}
	return maxMulti
}

func maxArea2(height []int) int {
	maxMulti := 0
	left, right := 0, len(height)-1
	for left < right {
		wide := right - left
		hight := min(height[left], height[right])
		area := wide * hight
		if height[left] <= height[right] {
			left++ // 往右边走找更长的线
		} else {
			right-- // 往左边走
		}
		maxMulti = max(area, maxMulti)
	}
	return maxMulti
}

func twoSum2(nums []int, total int) []int {
	num2index := make(map[int]int, len(nums))
	//for i, num := range nums {
	//	num2index[num] = i
	//}
	for i, num := range nums {
		pair := total - num
		if j, ok := num2index[pair]; ok && j != i {
			return []int{j, i}
		}
		num2index[num] = i
	}
	return nil
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	tatalNum := len1 + len2
	if tatalNum == 0 {
		return -1
	}
	if tatalNum%2 == 0 {
		lNum := findKth(nums1, nums2, tatalNum/2)
		rNum := findKth(nums1, nums2, tatalNum/2+1) // 此处 +1 与上边 len(mergedNums)/2 同理
		return float64(lNum+rNum) / 2
	}
	return float64(findKth(nums1, nums2, tatalNum/2+1))
}

func findKth(nums1, nums2 []int, k int) int {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		n1, n2 = n2, n1
		nums1, nums2 = nums2, nums1 // 为避免数组长度的分类讨论，先做预处理
	}

	if n1 == 0 {
		return nums2[k-1] // bingo
	}

	if k == 1 {
		return min(nums1[0], nums2[0]) // bingo
	}

	k1 := min(k/2, n1) // 避免越界
	k2 := k - k1       // 不能理想的 k/2, k/2 划分

	fmt.Println(nums1, k1-1, nums2, k2-1)

	switch {
	case nums1[k1-1] < nums2[k2-1]:
		return findKth(nums1[k1:], nums2, k2) // 彻底舍弃区域 1
	case nums1[k1-1] > nums2[k2-1]:
		return findKth(nums1, nums2[k2:], k1) // 彻底舍弃区域 3
	default:
		return nums1[k1-1] // bingo
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
