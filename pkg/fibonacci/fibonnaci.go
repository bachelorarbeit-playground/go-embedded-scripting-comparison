package fibonacci

import (
	"github.com/rs/zerolog/log"
	golua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func RunFibonacciLua(script []byte, fibonacci_number int) int {
	L := golua.NewState()
	defer L.Close()

	if err := L.DoString(string(script)); err != nil {
		log.Error().Msgf("loading Lua script failed: %s", err.Error())
	}

	if err := L.CallByParam(golua.P{
		Fn:      L.GetGlobal("fib"),
		NRet:    1,
		Protect: true,
	}, luar.New(L, fibonacci_number)); err != nil {

		log.Error().Msgf("executing Lua function '%s' failed: %s", "fib", err.Error())
	}

	result := L.Get(-1)
	L.Pop(1)

	// fmt.Printf()
	number := result.(golua.LNumber)
	return int(number)
}
