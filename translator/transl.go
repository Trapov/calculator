package translator

import (
    "calculator/structs"
    "strconv"
    "fmt"
    "strings"
    "log"
)

var stackOperators = structs.NewStack()
var stackOut = structs.NewStack()
var stringOut = ""

var operators = [4]string{"+", "-", "*", "/"}

func makeSliceOfTwoNumbers(stack *structs.Stack) (string, string) {
    var bufStack = structs.NewStack()

    var lOperand, rOperand string

    for stack.Size() > 0 {
        if rOperand == "" {
            rOperand, _ = stack.Pop()
        }
        if lOperand == "" {
            lOperand, _ = stack.Pop() //TODO: изменить на рядом проверку
        }
        fmt.Println("[lOperand] : " + lOperand + " | [rOperand] : " + rOperand)
        var lCheck, _ = isOperatorAndLevel(lOperand)
        var rCheck, _ = isOperatorAndLevel(rOperand)

        if !(lCheck && rCheck) {
            if !(lCheck || rCheck) {
                break
            } else {
                if !lCheck {
                    bufStack.Push(rOperand)
                    rOperand = lOperand
                    lOperand = ""
                } else {
                    bufStack.Push(lOperand)
                    lOperand = ""
                }
            }
        }
    }

    for bufStack.Size() > 0 {
        var propValue, _ = bufStack.Pop()
        stack.Push(propValue)
    }
    return lOperand, rOperand

}

func translate(expr string) string {
    var bufStackOperators = structs.NewStack()
    for _, e := range expr {
        if isOperator, precedenceCurrentOperator := isOperatorAndLevel(string(e)); isOperator {
            for stackOperators.Size() > 0 {
                var operatorInStack, _ = stackOperators.Pop()
                if _, precedenceOperatorInStack := isOperatorAndLevel(string(operatorInStack));
                    precedenceOperatorInStack >= precedenceCurrentOperator {
                    stringOut += string(operatorInStack)
                } else {
                    bufStackOperators.Push(operatorInStack)
                }
            }
            for bufStackOperators.Size() > 0 {
                var e, _ = bufStackOperators.Pop()
                stackOperators.Push(e)
            }
            stackOperators.Push(string(e))
        } else {
            stringOut += string(e)
        }
    }

    for stackOperators.Size() > 0 {
        var operator, _ = stackOperators.Pop()
        stringOut += operator
    }
    return stringOut
}

func Process(expression string) float64 {
    return solve(translate(expression))
}
func solve(rpnExp string) float64 {
    for _, token := range rpnExp {
        if isOperator, _ := isOperatorAndLevel(string(token)); isOperator {
            if stackOut.Size() < 2 {
                log.Fatal("[Error]: the given RPN expression is not correct! = > " + rpnExp);
            } else {
                var leftOperand,_ = stackOut.Pop()
                var rightOperand,_ = stackOut.Pop()
                log.Println("[Info]: Trying to solve : " + leftOperand + " " + string(token) + " " + rightOperand)
                var calAnswer = strconv.FormatInt(calculate(leftOperand, rightOperand, string(token)), 16)
                stackOut.Push(string(calAnswer))
                log.Print("[Info]: Answer ["+string(calAnswer)+"] is pushed onto the stack : ")
                log.Println(stackOut)
            }
        } else {
            stackOut.Push(string(token))
        }

    }
    var answer,_ = stackOut.Pop()
    var floatAnswer,_ = strconv.ParseFloat(answer, 64)
    return floatAnswer
}


func calculate(lOperand string, rOperand string, operator string) int64 {
    var leftConverted, _ = strconv.ParseInt(lOperand, 16, 64)
    var rightConverted, _ = strconv.ParseInt(rOperand, 16, 64)
    switch operator {
    case "*":
        return leftConverted * rightConverted
    case "/":
        return leftConverted / rightConverted
    case "+":
        return leftConverted + rightConverted
    case "-":
        return leftConverted - rightConverted
    default:
        return 0
    }
}

func isOperatorAndLevel(symbol string) (bool, int) {
    for l, e := range operators {
        if strings.ContainsAny(symbol, string(e)) {
            return true, l
        }
    }
    return false, -1
}
