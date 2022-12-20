package main

import (
	"fmt"
)

func main() {
	var factory AbstractFactory
	var tv ITelevision
	var air IAirConditioner

	factory = &HuaWeiFactory{}
	tv = factory.CreateTelevision()
	air = factory.CreateAirConditioner()
	tv.Watch()
	air.SetTemperature(25)

	factory = &MiFactory{}
	tv = factory.CreateTelevision()
	air = factory.CreateAirConditioner()
	tv.Watch()
	air.SetTemperature(26)
}

// AbstractFactory 抽象工厂类
// 目前有两个实际工厂类一个是华为的工厂，一个是小米的工厂
// 他们用来实际生产自家的产品设备
type AbstractFactory interface {
	CreateTelevision() ITelevision
	CreateAirConditioner() IAirConditioner
}

type ITelevision interface {
	Watch()
}

type IAirConditioner interface {
	SetTemperature(int)
}

// HuaWeiTV 华为电视
type HuaWeiTV struct{}

func (ht *HuaWeiTV) Watch() {
	fmt.Println("Watch HuaWei TV")
}

// HuaWeiAirConditioner 华为空调
type HuaWeiAirConditioner struct{}

func (ha *HuaWeiAirConditioner) SetTemperature(temp int) {
	fmt.Printf("HuaWei AirConditioner set temperature to %d ℃\n", temp)
}

// MiTV 小米电视
type MiTV struct{}

func (mt *MiTV) Watch() {
	fmt.Println("Watch HuaWei TV")
}

// MiAirConditioner 小米空调
type MiAirConditioner struct{}

func (ma *MiAirConditioner) SetTemperature(temp int) {
	fmt.Printf("Mi AirConditioner set temperature to %d ℃\n", temp)
}

// HuaweiFactory 华为工厂，实现 AbstractFactory
type HuaWeiFactory struct{}

func (hf *HuaWeiFactory) CreateTelevision() ITelevision {
	return &HuaWeiTV{}
}
func (hf *HuaWeiFactory) CreateAirConditioner() IAirConditioner {
	return &HuaWeiAirConditioner{}
}

// MiFactory 小米工厂，实现 AbstractFactory
type MiFactory struct{}

func (mf *MiFactory) CreateTelevision() ITelevision {
	return &MiTV{}
}
func (mf *MiFactory) CreateAirConditioner() IAirConditioner {
	return &MiAirConditioner{}
}
