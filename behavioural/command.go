// https://golangbyexample.com/command-design-pattern-in-golang/
package behavioural

import "fmt"

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type command interface {
	execute()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type device interface {
	on()
	off()
}

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
