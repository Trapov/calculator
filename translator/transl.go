package translator

import (
    "calculator/structs"
    "strconv"
    "fmt"
)

var stackOperators = structs.NewStack()
var stackOut = structs.NewStack()

const operators = "+-*/"

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

func Process(expression string) float64 {
    for _, e := range expression {

        var isOperator, level = isOperatorAndLevel(string(e))

        if isOperator {
            if stackOut.Size() > 1 {
                var operatorHead, _ = stackOperators.Pop()
                var _, levelHead = isOperatorAndLevel(operatorHead)
                if level >= levelHead {
                    fmt.Println(
                        "[!level > levelHead!]  " +
                            "\n [push into Operators] : " + string(e) +
                            "\n [push into Operators] : " + operatorHead)
                    stackOperators.Push(string(e))
                    stackOperators.Push(operatorHead)
                } else { //TODO: пока новый оператор меньше приоритетом, то выталкиваем в ответ
                    stackOperators.Push(operatorHead)
                    for stackOperators.Size() > 0 {
                        var op, _ = stackOperators.Pop()
                        var _, levelOp = isOperatorAndLevel(op)
                        if levelOp > level {

                            fmt.Print(
                                "[!levelOp > level!]  " +
                                    "\n [push into Operators] : " + op +
                                    "\n [push into Operators] : " + string(e))

                            stackOperators.Push(op)
                            stackOperators.Push(string(e))
                        } else {
                            stackOut.Push(op)
                        }
                    }
                }
            } else {
                stackOperators.Push(string(e))
            }
        } else {
            stackOut.Push(string(e))
        }
    }
    return postProcess(stackOut)
}

func postProcess(stack *structs.Stack) float64 {
    for stackOperators.Size() > 0 {
        var e, _ = stackOperators.Pop()
        stack.Push(e)
    }
    var floatAnswer float64
    fmt.Print("[stackOut]: ")
    fmt.Println(stack)
    for stack.Size() > 1 {
        var valuePopped, _ = stack.Pop()
        var isOperator, _ = isOperatorAndLevel(valuePopped)
        fmt.Println("[valuePopped] : " + valuePopped)
        if isOperator {
            var lOperand, rOperand = makeSliceOfTwoNumbers(stack)
            fmt.Println("[leftOperand]: " + lOperand + "\n" +
                "[rightOperand]:" + rOperand)
            floatAnswer += calculate(lOperand, rOperand, valuePopped)
        }
    }
    return floatAnswer
}

func calculate(lOperand string, rOperand string, operator string) float64 {
    var leftConverted, _ = strconv.ParseFloat(lOperand, 64)
    var rightConverted, _ = strconv.ParseFloat(rOperand, 64)
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
        if symbol == string(e) {
            return true, l
        }
    }
    return false, -1
}
