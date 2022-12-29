package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// BankBusinessHandler 银行柜台办理业务各个流程
type BankBusinessHandler interface {
	TakeRowNumnber()
	WaitInHead()
	HandleBussiness()
	Commentate()
	checkVipIdentify() bool
}

type BankBusinessExecutor struct {
	handler BankBusinessHandler
}

// ExecuteBankBusiness 银行柜台办理业务全流程
func (b *BankBusinessExecutor) ExecuteBankBusiness() {
	b.handler.TakeRowNumnber()
	if !b.handler.checkVipIdentify() {
		b.handler.WaitInHead()
	}
	b.handler.HandleBussiness()
	b.handler.Commentate()
}

func NewBankBusinessExecutor(businessHandler BankBusinessHandler) *BankBusinessExecutor {
	return &BankBusinessExecutor{handler: businessHandler}
}

// DefaultBusinessHandler 默认实现
type DefaultBusinessHandler struct{}

func (d *DefaultBusinessHandler) TakeRowNumnber() {
	fmt.Println("请拿好您的取件码：" + strconv.Itoa(rand.Intn(100)) +
		" ，注意排队情况，过号后顺延三个安排")
}

func (d *DefaultBusinessHandler) WaitInHead() {
	fmt.Println("排队等号中...")
	time.Sleep(5 * time.Second)
	fmt.Println("请去窗口xxx...")
}

func (d *DefaultBusinessHandler) Commentate() {
	fmt.Println("请对我的服务作出评价，满意请按0，满意请按0，(～￣▽￣)～")
}

func (d *DefaultBusinessHandler) checkVipIdentify() bool {
	return false
}

// DepositBusinessHandler 取款
type DepositBusinessHandler struct {
	*DefaultBusinessHandler
	userVip bool
}

func (d *DepositBusinessHandler) HandleBussiness() {
	fmt.Println("处理业务中：取款....")
}

func main() {
	dh := &DepositBusinessHandler{
		DefaultBusinessHandler: &DefaultBusinessHandler{},
		userVip:                false,
	}
	bbe := NewBankBusinessExecutor(dh)
	bbe.ExecuteBankBusiness()
}
