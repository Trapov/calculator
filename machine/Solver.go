package machine

import (
    "strconv"
    "log"
    "calculator/machine/utils"
    "calculator/structs"
)

func solve(rpnStack *structs.Stack) float64 {
    for rpnStack.Size() > 0 {
        var token, _ = rpnStack.PopBack()
        if isOperator, _ := utils.Operator(token).IsOperatorAndLevel(operators); isOperator {
            if stackOut.Size() < 2 {
                log.Fatal("[Error]: the given RPN expression is not correct! = > ", rpnStack);
            } else {
                var rightOperand, _ = stackOut.PopHead()
                var leftOperand, _ = stackOut.PopHead()
                log.Println("[Info]: Trying to solve : " + leftOperand + " " + string(token) + " " + rightOperand)
                var calAnswer = strconv.FormatInt(calculate(leftOperand, rightOperand, string(token)), 10)
                stackOut.Push(string(calAnswer))
                log.Print("[Info]: Answer [" + string(calAnswer) + "] is pushed onto the stack : ")
                log.Println(stackOut)
            }
        } else {
            stackOut.Push(string(token))
        }

    }
    if stackOut.Size() == 1 {
        var answer, _ = stackOut.PopHead()
        var floatAnswer, _ = strconv.ParseFloat(answer, 64)
        return floatAnswer
    } else {
        log.Fatal("[Error]: Can't get the final answer!")
        return 0
    }

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
