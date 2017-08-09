package machine

import (
    "calculator/structs"
    "calculator/machine/utils"
    "log"
)

var stackOperators = structs.NewStack()
var stackRPN = structs.NewStack()
var operators = []string{"+-", "*/"}

func translate(expr string) *structs.Stack {
    var bufStackOperators = structs.NewStack()
    for _, e := range expr {
        if isOperator, precedenceCurrentOperator :=
            utils.Operator(string(e)).IsOperatorAndLevel(operators); isOperator {
            for stackOperators.Size() > 0 {
                var operatorInStack, _ = stackOperators.PopHead()
                if _, precedenceOperatorInStack :=
                    utils.Operator(string(operatorInStack)).IsOperatorAndLevel(operators);
                    precedenceOperatorInStack >= precedenceCurrentOperator {
                    stackRPN.Push(string(operatorInStack))
                } else {
                    bufStackOperators.Push(operatorInStack)
                }
            }
            for bufStackOperators.Size() > 0 {
                var operator, _ = bufStackOperators.PopHead()
                stackOperators.Push(operator)
            }
            stackOperators.Push(string(e))
        } else {
            stackRPN.Push(string(e))
        }
    }

    for stackOperators.Size() > 0 {
        var operator, _ = stackOperators.PopHead()
        stackRPN.Push(operator)
    }
    log.Print("[Info]: RPN : ", stackRPN)
    return stackRPN
}
