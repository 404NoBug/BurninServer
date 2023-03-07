package main

import (
	"fmt"
	"time"
)

//state interface
type LightState interface {
	// 亮起当前状态的交通灯
	Light()
	// 转换的新状态的时候，调用的方法
	EnterState()
	// 设置一个状态要转变的状态
	NextLight(light *TrafficLight)
	//检测车速
	CarPassingSpeed(*TrafficLight, int, string)
}

//Context
type TrafficLight struct {
	State      LightState
	SpeedLimit int
}

func NewSimpleTrafficLight(speedLimit int) *TrafficLight {
	return &TrafficLight{
		SpeedLimit: speedLimit,
		State:      NewRedState(),
	}
}

type DefaultLightState struct {
	StateName string
}

func (state *DefaultLightState) CarPassingSpeed(road *TrafficLight, speed int, licensePlate string) {
	if speed > road.SpeedLimit {
		fmt.Printf("Car with license %s was speeding\n", licensePlate)
	}
}

func (state *DefaultLightState) EnterState() {
	fmt.Println("changed state to:", state.StateName)
}

func (tl *TrafficLight) TransitionState(newState LightState) {
	tl.State = newState
	tl.State.EnterState()
}

//红灯状态
type redState struct {
	DefaultLightState
}

func NewRedState() *redState {
	state := &redState{}
	state.StateName = "RED"
	return state
}

func (state *redState) Light() {
	fmt.Println("红灯亮起，不可行驶")
}

func (state *redState) CarPassingSpeed(light *TrafficLight, speed int, licensePlate string) {
	//红灯不能行驶，所以这里要重写覆盖 DefaultLightState 里定义的方法
	if speed > 0 {
		fmt.Printf("Car with license \"%s\" ran a red light!\n", licensePlate)
	}
}

func (state *redState) NextLight(light *TrafficLight) {
	light.TransitionState(NewGreedState())
}

//绿灯状态
type greedState struct {
	DefaultLightState
}

func NewGreedState() *greedState {
	state := &greedState{}
	state.StateName = "GREED"
	return state
}

func (state *greedState) Light() {
	fmt.Println("绿灯亮起，请行驶")
}

func (state *greedState) NextLight(light *TrafficLight) {
	light.TransitionState(NewAmberState())
}

//黄灯状态
type amberState struct {
	DefaultLightState
}

func NewAmberState() *amberState {
	state := &amberState{}
	state.StateName = "AMBER"
	return state
}

func (state *amberState) Light() {
	fmt.Println("黄灯亮起，请注意")
}

func (state *amberState) NextLight(light *TrafficLight) {
	light.TransitionState(NewRedState())
}

func main() {
	trafficLight := NewSimpleTrafficLight(500)
	interval := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-interval.C:
			trafficLight.State.Light()
			trafficLight.State.CarPassingSpeed(trafficLight, 25, "CN1024")
			trafficLight.State.NextLight(trafficLight)
		default:
		}
	}
}
