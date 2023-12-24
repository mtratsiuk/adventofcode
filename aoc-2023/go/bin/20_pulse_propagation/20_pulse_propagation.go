package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("20_pulse_propagation")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	config := NewConfig(in)
	pulses := make([]Pulse, 0)
	low := 0
	high := 0

	for i := 0; i < 1000; i += 1 {
		low += 1

		for _, o := range config.modules[string(ModuleTypeBroad)].outs {
			pulses = append(pulses, Pulse{PulseTypeLow, string(ModuleTypeBroad), o})
		}

		for len(pulses) > 0 {
			var pulse Pulse

			pulse, pulses = pulses[0], pulses[1:]

			if pulse.value == PulseTypeLow && slices.Contains([]string{"vz", "bq", "qh", "lt"}, pulse.to) {
				// solving part 2
				fmt.Printf("Low for %v at push %v\n", pulse.to, i)
			}

			if pulse.value == PulseTypeLow {
				low += 1
			} else {
				high += 1
			}

			mod := config.modules[pulse.to]

			if mod.variant == ModuleTypeFlip && pulse.value == PulseTypeLow {
				for _, o := range mod.outs {
					pulses = append(pulses, Pulse{FlipStateToPulseType[mod.flipState], pulse.to, o})
				}
				mod.flipState = !mod.flipState
			}

			if mod.variant == ModuleTypeConj {
				mod.conjState[pulse.from] = pulse.value

				allHigh := true

				for _, v := range mod.conjState {
					if v == PulseTypeLow {
						allHigh = false
						break
					}
				}

				for _, o := range mod.outs {
					pulses = append(pulses, Pulse{FlipStateToPulseType[allHigh], pulse.to, o})
				}
			}

			config.modules[pulse.to] = mod
		}
	}

	return low * high
}

func solve2(in string) int {
	// lol
	return gotils.LcmAll([]int{4476 - 655, 5862 - 1769, 6665 - 2776, 6526 - 2787})
}

var FlipStateToPulseType = map[bool]PulseType{
	true:  PulseTypeLow,
	false: PulseTypeHigh,
}

type PulseType int

const (
	PulseTypeLow  PulseType = 0
	PulseTypeHigh PulseType = 1
)

type Pulse struct {
	value PulseType
	from  string
	to    string
}

type ModuleType string

const (
	ModuleTypeFlip  ModuleType = "%"
	ModuleTypeConj  ModuleType = "&"
	ModuleTypeBroad ModuleType = "broadcaster"
)

type Module struct {
	variant   ModuleType
	flipState bool
	conjState map[string]PulseType
	ins       []string
	outs      []string
}

type Config struct {
	modules map[string]Module
}

func NewConfig(in string) Config {
	modules := make(map[string]Module, 0)

	for _, line := range strings.Split(in, "\n") {
		line = strings.ReplaceAll(line, " ", "")
		moduleStr, outsStr, _ := strings.Cut(line, "->")
		outs := strings.Split(outsStr, ",")

		module := Module{}
		module.outs = outs
		module.flipState = false
		module.conjState = make(map[string]PulseType, 0)

		var moduleName string

		if moduleStr == string(ModuleTypeBroad) {
			module.variant = ModuleTypeBroad
			moduleName = moduleStr
		} else if strings.HasPrefix(moduleStr, string(ModuleTypeFlip)) {
			module.variant = ModuleTypeFlip
			moduleName = moduleStr[1:]
		} else if strings.HasPrefix(moduleStr, string(ModuleTypeConj)) {
			module.variant = ModuleTypeConj
			moduleName = moduleStr[1:]
		}

		modules[moduleName] = module
	}

	for n := range modules {
		for _, o := range modules[n].outs {
			mod := modules[o]
			mod.ins = append(mod.ins, n)

			if mod.variant == ModuleTypeConj {
				mod.conjState[n] = PulseTypeLow
			}

			modules[o] = mod
		}
	}

	return Config{modules}
}
