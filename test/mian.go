package main

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"math"
)

type Hello struct{ Who string }

type Actor = actor.Actor

type HelloActor struct {
	actor.Actor
}

func NewHelloActor() Actor {
	return &HelloActor{}
}

func (state *HelloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

type WorldActor struct {
	actor.Actor
}

func NewWorldActor() Actor {
	return &WorldActor{}
}

func (state *WorldActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("World %v\n", msg.Who)
	}
}

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "love" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

//----------> 抽象层 <--------------
type Car interface {
	Run()
}
type Driver interface {
	Driver(car Car)
}

//----------> 实现层 <--------------
type Benz struct {
}

func (benz *Benz) Run() {
	fmt.Println("Benz Run")
}

type Bmw struct {
}

func (bmw *Bmw) Run() {
	fmt.Println("BMW Run")
}

type XiaoWang struct {
}

func (xiaowang *XiaoWang) Driver(car Car) {
	fmt.Println("xiaowang Driver")
	car.Run()
}

type XiaoZhang struct {
}

func (xiaoZhang *XiaoZhang) Driver(car Car) {
	fmt.Println("xiaoZhang Driver")
	car.Run()
}

//----------> 抽象层 <--------------
type Card interface {
	display()
}
type Memery interface {
	storge()
}
type CPU interface {
	calculate()
}
type Computer struct {
	card   Card
	memery Memery
	cPU    CPU
}

func NewComputer(cpu CPU, mem Memery, card Card) *Computer {
	return &Computer{
		card:   card,
		memery: mem,
		cPU:    cpu,
	}
}
func (computer *Computer) DoWork() {
	computer.cPU.calculate()
	computer.memery.storge()
	computer.card.display()
}

//----------> 实现层 <--------------
type IntelCPU struct {
	cPU CPU
}

func (intelCPU *IntelCPU) calculate() {
	fmt.Println("intelCPU calculate")
}

type IntelCard struct {
	card Card
}

func (intelCard *IntelCard) display() {
	fmt.Println("intelCard display")
}

type IntelMemery struct {
	memery Memery
}

func (intelMemery *IntelMemery) storge() {
	fmt.Println("intelMemery storge")
}

type Kingston struct {
	memery Memery
}

func (kingston *Kingston) storge() {
	fmt.Println("kingston storge")
}

type NVIDIA struct {
	card Card
}

func (nVIDIA *NVIDIA) display() {
	fmt.Println("nVIDIA display")
}

func main() {
	array := make([]int, 30)
	for i := 0; i < len(array); i++ {
		array[i] = i + 1
	}
	var result []int
	var resultNum int
	getReward(array, result, &resultNum, 2)
	fmt.Println("resultNum ", resultNum)
}

func getReward(array, resultArray []int, resultNum *int, num int) *int {
	if len(resultArray) == num {
		*resultNum++
	}
	for i := 0; i < len(array); i++ {
		newResult := make([]int, len(resultArray))
		copy(newResult, resultArray)
		newResult = append(newResult, array[i])
		newArray := make([]int, len(array))
		copy(newArray, array)
		newArray = append(newArray[i+1:])
		getReward(newArray, newResult, resultNum, num)
	}
	return resultNum
}

func getGroup(arrar, result []string, num int) {
	if len(result) == num {
		fmt.Printf("%v \n", result)
	}
	for i := 0; i < len(arrar); i++ {
		newResult := make([]string, len(result))
		copy(newResult, result)
		newResult = append(newResult, arrar[i])
		newArray := make([]string, len(arrar))
		copy(newArray, arrar)
		newArray = append(newArray[i+1:])
		getGroup(newArray, newResult, num)
	}
}

func testPermutate() {
	q_horses_time := map[string]float32{"q1": 1.0, "q2": 2.0, "q3": 3.0}
	t_horses_time := map[string]float32{"t1": 1.5, "t2": 2.5, "t3": 3.5}
	q_horses := []string{"q1", "q2", "q3"}
	t_horses := []string{"t1", "t2", "t3"}
	var result []string
	permutate(t_horses, result, q_horses, q_horses_time, t_horses_time)
}
func permutate(horse, result, qHorses []string, qHorsesTime, tHorsesTime map[string]float32) {
	if len(horse) == 0 {
		compare(result, qHorses, qHorsesTime, tHorsesTime)
		println()
		return
	}
	for i := 0; i < len(horse); i++ {
		//从剩下的未出战马匹中，选择一匹，加入结果
		newResult := make([]string, len(result))
		copy(newResult, result)
		newResult = append(newResult, horse[i])
		//将已选择的马匹从未出战的列表中移出
		restHorse := make([]string, len(horse))
		copy(restHorse, horse)
		restHorse = append(restHorse[:i], restHorse[i+1:]...)
		// 递归调用，对于剩余的马匹继续生成排列
		permutate(restHorse, newResult, qHorses, qHorsesTime, tHorsesTime)
	}
}
func compare(t, q []string, q_horses_time, t_horses_time map[string]float32) {
	t_won_cnt := 0
	fmt.Printf("%v \n", t)
	for i := 0; i < len(t); i++ {
		fmt.Printf("%v  %v", t_horses_time[t[i]], q_horses_time[q[i]])
		println()
		if t_horses_time[t[i]] < q_horses_time[q[i]] {
			t_won_cnt++
		}
	}
	if t_won_cnt > (len(t) / 2) {
		println("田忌获胜！")
	} else {
		println("齐王获胜！")
	}
}

func creatChan(array []int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range array {
			out <- v
		}
		close(out)
	}()
	return out
}
func merge_chan(chan1, chan2 chan int) []int {
	out := make(chan int)
	go func() {
		out1, ok1 := <-chan1
		out2, ok2 := <-chan2
		for {
			if ok1 || ok2 {
				if ok1 && ok2 {
					if out1 < out2 {
						out <- out1
						out1, ok1 = <-chan1
					} else {
						out <- out2
						out2, ok2 = <-chan2
					}
				} else if ok1 && !ok2 {
					out <- out1
					out1, ok1 = <-chan1
				} else {
					out <- out2
					out2, ok2 = <-chan2
				}
			} else {
				close(out)
				break
			}
		}
	}()
	var outArray []int
	for v := range out {
		outArray = append(outArray, v)
	}
	return outArray
}

func merge_sort(array []int) []int {
	if array == nil {
		return []int{0}
	}
	if len(array) == 1 {
		return array
	}
	mid := len(array) / 2
	left := array[0:mid]
	right := array[mid:]

	//嵌套调用
	left = merge_sort(left)
	right = merge_sort(right)

	//合并排序后的两半
	chan1 := creatChan(left)
	chan2 := creatChan(right)
	outCahan := merge_chan(chan1, chan2)
	return outCahan
}

func merge(left, right []int) []int {
	if left == nil {
		left = []int{0}
	}
	if right == nil {
		right = []int{0}
	}
	mergeOne := make([]int, len(left)+len(right))
	mi := 0
	li := 0
	ri := 0
	for li < len(left) && ri < len(right) {
		if left[li] <= right[ri] {
			mergeOne[mi] = left[li]
			li++
		} else {
			mergeOne[mi] = right[ri]
			ri++
		}
		mi++
	}
	//将某个数据剩下的数据放到合并后的数组中
	if li < len(left) {
		for i := li; i < len(left); i++ {
			mergeOne[mi] = left[i]
			mi++
		}
	} else {
		for i := ri; i < len(right); i++ {
			mergeOne[mi] = right[i]
			mi++
		}
	}
	return mergeOne
}

func get(total int, result []int) {
	reward := [4]int{1, 2, 5, 10}
	if total == 0 {
		fmt.Printf("%v", result)
		fmt.Println()
	} else if total < 0 {
		return
	} else {
		for i := 0; i < len(reward); i++ {
			newResult := make([]int, len(result))
			copy(newResult, result)
			newResult = append(newResult, reward[i])
			get(total-reward[i], newResult)
		}
	}
}

func getDivision(total int, result []int) {
	if total == 1 {
		if !member(1, result) {
			result := append(result, 1)
			fmt.Printf("%v \n", result)
			return
		}
	} else {
		for i := 1; i <= total; i++ {
			if i == 1 && member(1, result) {
				continue
			}
			newResult := make([]int, len(result))
			copy(newResult, result)
			newResult = append(newResult, i)
			if total%i != 0 {
				continue
			}
			get(total/i, newResult)
		}
	}
}

func member(element int, intslice []int) bool {
	for _, v := range intslice {
		if v == element {
			return true
		}
	}
	return false
}

func maiziNum(n int) int64 {
	var totalNum int64 //总麦子数
	var curNum int64   //当前麦子数
	curNum = 1
	totalNum += curNum
	for i := 2; i <= n; i++ {
		curNum *= 2
		totalNum += curNum
	}
	return totalNum
}

func MySqrt(n int, deltaThreshold float64, maxTry int) float64 {
	if n < 0 {
		return -1.0
	}
	min := 1.0
	max := float64(n)
	for i := 0; i < maxTry; i++ {
		middle := (min + max) / 2
		square := middle * middle
		delta := math.Abs((square / float64(n)) - 1)
		if delta <= deltaThreshold {
			return middle
		} else {
			if square > float64(n) {
				max = middle
			} else {
				min = middle
			}
		}
	}
	return -2.0
}
func returnButDefer() (t int) { //t初始化0， 并且作用域为该函数全域

	defer func() {
		t = t * 10
	}()

	return 1
}

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}
