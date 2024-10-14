package controllers

import (
	"bufio"
	"fmt"
	"microwave/constants"
	"microwave/interfaces"
	"os"
	"strconv"
	"strings"
)

type CommandController struct {
	MicrowaveService interfaces.MicrowaveInterface
	commands         map[string]func(args []string)
}

func NewCommandController(service interfaces.MicrowaveInterface) *CommandController {
	cc := &CommandController{
		MicrowaveService: service,
		commands:         make(map[string]func(args []string)),
	}
	cc.initCommands()
	return cc
}

func (cc *CommandController) initCommands() {
	cc.commands = map[string]func(args []string){
		constants.CommandOn:       cc.commandOn,
		constants.CommandOff:      cc.commandOff,
		constants.CommandStart:    cc.commandStart,
		constants.CommandStop:     cc.commandStop,
		constants.CommandSetPower: cc.commandSetPower,
		constants.CommandSetTime:  cc.commandSetTime,
		constants.CommandShowTime: cc.commandShowTime,
		constants.CommandOpen:     cc.commandOpen,
		constants.CommandClose:    cc.commandClose,
		constants.CommandInsert:   cc.commandInsert,
		constants.CommandGet:      cc.commandGet,
		constants.CommandExit:     cc.commandExit,
	}
}

func (cc *CommandController) printAvailableCommands() {
	fmt.Println("Доступные команды:")
	fmt.Printf(" %s, %s,\n", constants.CommandOn, constants.CommandOff)
	fmt.Printf(" %s,\n", constants.CommandStart)
	fmt.Printf(" %s,\n", constants.CommandStop)
	fmt.Printf(" %s <1-10>,\n", constants.CommandSetPower)
	fmt.Printf(" %s <секунды>,\n", constants.CommandSetTime)
	fmt.Printf(" %s,\n", constants.CommandOpen)
	fmt.Printf(" %s,\n", constants.CommandClose)
	fmt.Printf(" %s,\n", constants.CommandInsert)
	fmt.Printf(" %s,\n", constants.CommandGet)
	fmt.Printf(" %s\n", constants.CommandExit)
}

func (cc *CommandController) Run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Добро пожаловать в симулятор микроволновки!")
	cc.printAvailableCommands()

	for {
		fmt.Println("> ")
		input, _ := reader.ReadString('\n')
		args := strings.Split(input, " ")

		if len(args) == 0 || args[0] == "" {
			continue
		}

		command := strings.TrimSpace(args[0])

		command = strings.ToLower(command)

		cmdTestFunc := cc.commands[command]

		fmt.Println(cmdTestFunc)

		if cmdFunc, exists := cc.commands[command]; exists {
			cmdFunc(args)
		} else {
			fmt.Println("Неизвестная команда, попробуйте еще раз.")
		}
	}

}

func (cc *CommandController) commandOn(args []string) {
	err, mw := cc.MicrowaveService.TurnOn()
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println(mw)
}

func (cc *CommandController) commandOff(args []string) {
	err, mw := cc.MicrowaveService.TurnOff()
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println(mw)
}

func (cc *CommandController) commandStart(args []string) {
	err, mw := cc.MicrowaveService.Start()
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println(mw)
}

func (cc *CommandController) commandStop(args []string) {
	err, mw := cc.MicrowaveService.Stop()
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println(mw)
}

func (cc *CommandController) commandSetPower(args []string) {
	if len(args) < 2 {
		fmt.Println("Использование: setpower <1-10>")
		return
	}
	level, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Ошибка: уровень мощности должен быть числом")
		return
	}
	var correctLevel = false
	for _, powerLevelItem := range constants.ValidPowerLevels {
		if level == int(powerLevelItem) {
			correctLevel = true
		}
	}
	if !correctLevel {
		fmt.Println("Ошибка: мощность не должна быть из 600, 800, 1000") //TODO replace with values
		return
	}
	err, mw := cc.MicrowaveService.SetPowerLevel(constants.PowerLevel(level))
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println(mw)
}

func (cc *CommandController) commandSetTime(args []string) {
	if len(args) < 2 {
		fmt.Println("Использование: settime <секунды>")
		return
	}
	seconds, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Ошибка: время должно быть числом")
		return
	}
	err, mw := cc.MicrowaveService.SetTimer(seconds)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println(mw)
}

func (cc *CommandController) commandShowTime(args []string) {}

func (cc *CommandController) commandOpen(args []string) {
	err, mw := cc.MicrowaveService.OpenDoor()
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println(mw)
}

func (cc *CommandController) commandClose(args []string) {
	err, mw := cc.MicrowaveService.CloseDoor()
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println(mw)
}

func (cc *CommandController) commandInsert(args []string) {
	err, mw := cc.MicrowaveService.InsertFood()
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println(mw)
}

func (cc *CommandController) commandGet(args []string) {
	err, mw := cc.MicrowaveService.GetFood()
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println(mw)
}

func (cc *CommandController) commandExit(args []string) {
	fmt.Println("Выход из программы.")
	os.Exit(0)
}
