package machine

import "calculator/structs"

var stackOut = structs.NewStack()

func Process(expression string) float64 {
    return solve(translate(expression))
}
