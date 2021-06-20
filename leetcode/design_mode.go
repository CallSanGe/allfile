package main

import (
	"fmt"
	"sync"
)

// 单例模式

type SingleObj struct {
	Name string `json:"name"`
}

var instance *SingleObj

func NewSingleObj() *SingleObj {
	if instance == nil {
		instance = &SingleObj{}
	}
	return instance
}

// 工厂模式
// 简单工厂
type Factory1 struct{}

func (f Factory1) Generate(name string) Product {

	switch name {
	case "product1":
		return Product1{}
	case "project2":
		return Product2{}
	default:
		return nil
	}
}

type Factory2 struct{}

func (f Factory2) Generate(name string) Product {

	switch name {
	case "product3":
		return Product3{}
	case "project4":
		return Product4{}
	default:
		return nil
	}
}

type Factory interface {
	Generate(name string) Product
}

type Product interface {
	Create()
}

// 产品
type Product1 struct {
	Name string `json:"name"`
}

func (p Product1) Create() {
	fmt.Println("这是产品1")
}

type Product2 struct {
	Name string `json:"name"`
}

func (p Product2) Create() {
	fmt.Println("这是产品2")
}

// 产品
type Product3 struct {
	Name string `json:"name"`
}

func (p Product3) Create() {
	fmt.Println("这是产品3")
}

type Product4 struct {
	Name string `json:"name"`
}

func (p Product4) Create() {
	fmt.Println("这是产品4")
}

// 抽象工厂
// 多个工厂可以生产多个产品
// 不通的工厂生产的产品还都不一样
type FactoryInter interface {
	Create1() ProductInter
	Create2() ProductInter
}

type ProductInter interface {
	Introduce()
}

type FactoryA struct {
}

func (f FactoryA) Create1() ProductInter {
	return &ProductA1{}
}

func (f FactoryA) Create2() ProductInter {
	return &ProductA2{}
}

type ProductA1 struct {
}

func (p ProductA1) Introduce() {

}

type ProductA2 struct {
}

func (p ProductA2) Introduce() {

}

// 加锁模式 这个和懒汉模式啥区别 双重检测，主要用于多线程的使用吗？

var mu *sync.Mutex

func NewSingleObjMutex() *SingleObj {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &SingleObj{}
	}
	return instance
}

// 代理模式

// 建造者模式 和选项模式啥区别
type Header struct {
	SrcAddr  string
	SrcPort  uint64
	DestAddr string
	DestPort uint64
	Items    map[string]string
}
type Body struct {
	Items []string
}

type Message struct {
	Header *Header
	Body   *Body
}

type builder struct {
	once *sync.Once
	msg  *Message
}

func Builder() *builder {
	return &builder{
		once: &sync.Once{},
		msg:  &Message{&Header{}, &Body{}},
	}
}

func (b *builder) WithSrcAddr(srcAddr string) *builder {
	b.msg.Header.SrcAddr = srcAddr
	return b
}

func (b *builder) WithSrcPort(srcPort uint64) *builder {
	b.msg.Header.SrcPort = srcPort
	return b
}

func (b *builder) WithDestAddr(destAddr string) *builder {
	b.msg.Header.DestAddr = destAddr
	return b
}

func (b *builder) WithDestPort(destPort uint64) *builder {
	b.msg.Header.DestPort = destPort
	return b
}

func (b *builder) WithHeaderItem(key, value string) *builder {
	// 确保初始化过后不在初始化了
	b.once.Do(func() {
		b.msg.Header.Items = make(map[string]string)
	})
	b.msg.Header.Items[key] = value
	return b
}

func (b *builder) WithBodyItem(record string) *builder {
	b.msg.Body.Items = append(b.msg.Body.Items, record)
	return b
}

// 创建Message对象，在最后一步调用
func (b *builder) Build() *Message {
	return b.msg
}

func cccc() {
	message := Builder().
		WithSrcAddr("192.168.0.1").
		WithSrcPort(1234).
		WithDestAddr("192.168.0.2").
		WithDestPort(8080).
		WithHeaderItem("contents", "application/json").
		WithBodyItem("record1").
		WithBodyItem("record2").
		Build()
	fmt.Println(message)
}
