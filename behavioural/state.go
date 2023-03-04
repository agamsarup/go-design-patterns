// https://golangbyexample.com/state-design-pattern-go/
package behavioural

import "fmt"

type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state

	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}
	noItemState := &noItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

type noItemState struct {
	vendingMachine *vendingMachine
}

func (i *noItemState) requestItem() error {
	return fmt.Errorf("item out of stock")
}

func (i *noItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *noItemState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *noItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
}

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (i *hasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *hasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (i *hasItemState) dispenseItem() error {
	return fmt.Errorf("please select item first")
}

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("item Dispense in progress")
}

func (i *itemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("please insert money first")
}

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (i *hasMoneyState) requestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}
