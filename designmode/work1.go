/*
练习：

	设计一个电脑主板架构，电脑包括（显卡，内存，CPU）3个固定的插口，
	显卡具有显示功能（display，功能实现只要打印出意义即可），
	内存具有存储功能（storage），
	cpu具有计算功能（calculate）。
	现有Intel厂商，nvidia厂商，Kingston厂商，均会生产以上三种硬件。
	要求组装两台电脑，
			    1台（Intel的CPU，Intel的显卡，Intel的内存）
			    1台（Intel的CPU， nvidia的显卡，Kingston的内存）
	用抽象工厂模式实现。
*/
package main

import (
	"fmt"
)

// ===============抽象层================

type AbstractCPU interface {
	Calculate()
}
type AbstractGPU interface {
	Display()
}
type AbstractMemory interface {
	Storage()
}
type AbstractFactory interface {
	CreateCPU() AbstractCPU
	CreateGPU() AbstractGPU
	CreateMemory() AbstractMemory
}

//===============实现层================

// Intel产品族
type IntelCPU struct{}

func (ic *IntelCPU) Calculate() {
	fmt.Println("Intel CPU Calculate...")
}

type IntelGPU struct{}

func (ig *IntelGPU) Display() {
	fmt.Println("Intel GPU Display...")
}

type IntelMemory struct{}

func (im *IntelMemory) Storage() {
	fmt.Println("Intel Memory Storage...")
}

type IntelFactory struct{}

func (ifa *IntelFactory) CreateCPU() (cpu AbstractCPU) {
	cpu = new(IntelCPU)
	return
}

func (ifa *IntelFactory) CreateGPU() (gpu AbstractGPU) {
	gpu = new(IntelGPU)
	return
}

func (ifa *IntelFactory) CreateMemory() (memory AbstractMemory) {
	memory = new(IntelMemory)
	return
}

// Nvidia产品族
type NvidiaCPU struct{}

func (nc *NvidiaCPU) Calculate() {
	fmt.Println("Nvidia CPU Calculate...")
}

type NvidiaGPU struct{}

func (ng *NvidiaGPU) Display() {
	fmt.Println("Nvidia GPU Display...")
}

type NvidiaMemory struct{}

func (nm *NvidiaMemory) Storage() {
	fmt.Println("Nvidia Memory Storage...")
}

type NvidiaFactory struct{}

func (nf *NvidiaFactory) CreateCPU() (cpu AbstractCPU) {
	cpu = new(NvidiaCPU)
	return
}

func (nf *NvidiaFactory) CreateGPU() (gpu AbstractGPU) {
	gpu = new(NvidiaGPU)
	return
}

func (nf *NvidiaFactory) CreateMemory() (memory AbstractMemory) {
	memory = new(NvidiaMemory)
	return
}

// Kingston产品族
type KingstonCPU struct{}

func (kc *KingstonCPU) Calculate() {
	fmt.Println("Kingston CPU Calculate...")
}

type KingstonGPU struct{}

func (kg *KingstonGPU) Display() {
	fmt.Println("Kingston GPU Display...")
}

type KingstonMemory struct{}

func (km *KingstonMemory) Storage() {
	fmt.Println("Kingston Memory Storage...")
}

type KingstonFactory struct{}

func (kf *KingstonFactory) CreateCPU() (cpu AbstractCPU) {
	cpu = new(KingstonCPU)
	return
}

func (kf *KingstonFactory) CreateGPU() (gpu AbstractGPU) {
	gpu = new(KingstonGPU)
	return
}

func (kf *KingstonFactory) CreateMemory() (memory AbstractMemory) {
	memory = new(KingstonMemory)
	return
}

//===============业务层================

type Computer struct {
	Cpu    AbstractCPU
	Gpu    AbstractGPU
	Memory AbstractMemory
}

func (c *Computer) Run() {
	c.Cpu.Calculate()
	c.Gpu.Display()
	c.Memory.Storage()
}

func main() {

	var iFac AbstractFactory
	iFac = new(IntelFactory)
	var nFac AbstractFactory
	nFac = new(NvidiaFactory)
	var kFac AbstractFactory
	kFac = new(KingstonFactory)

	//1台（Intel的CPU，Intel的显卡，Intel的内存）
	var c1 Computer
	c1.Cpu = iFac.CreateCPU()
	c1.Gpu = iFac.CreateGPU()
	c1.Memory = iFac.CreateMemory()
	c1.Run()

	// 1台（Intel的CPU， nvidia的显卡，Kingston的内存）
	var c2 Computer
	c2.Cpu = iFac.CreateCPU()
	c2.Gpu = nFac.CreateGPU()
	c2.Memory = kFac.CreateMemory()
	c2.Run()
}
