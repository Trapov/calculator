package machine

import (
	"calculator/machine/utils"
	"calculator/structs"
	"errors"
	"log"
	"strconv"
)

var stackOut = structs.NewStack()

func solve(rpnStack *structs.Stack) (float64, error) {
	for rpnStack.Size() > 0 {
		var token, _ = rpnStack.PopBack()
		if isOperator, _ := utils.Operator(token).IsOperatorAndLevel(operators); isOperator {
			if stackOut.Size() < 2 {
				return 0, errors.New("[Error]: Given RPN expression isn't correct!")
			} else {
				var rightOperand, _ = stackOut.PopHead()
				var leftOperand, _ = stackOut.PopHead()
				log.Println("[Info]: Trying to solve : " + leftOperand + " " + string(token) + " " + rightOperand)
				var stringAnswer, errAnswer = calculate(leftOperand, rightOperand, string(token))
				if errAnswer != nil {
					return 0, errAnswer
				}
				var answerFormatted = strconv.FormatInt(stringAnswer, 10)
				stackOut.Push(string(answerFormatted))
				log.Print("[Info]: Answer [" + string(answerFormatted) + "] is pushed onto the stack : ")
				log.Println(stackOut)
			}
		} else {
			stackOut.Push(string(token))
		}
	}

	if stackOut.Size() == 1 {
		var answer, _ = stackOut.PopHead()
		var floatAnswer, _ = strconv.ParseFloat(answer, 64)
		return floatAnswer, nil
	} else {
		return 0, errors.New("[Error]: Can't get the final answer!")
	}

}

func calculate(lOperand string, rOperand string, operator string) (int64, error) {
	var leftConverted, erLeft = strconv.ParseInt(lOperand, 16, 64)
	if erLeft != nil {
		return 0, errors.New("[Error]: Can't parse the left value")
	}
	var rightConverted, erRight = strconv.ParseInt(rOperand, 16, 64)
	if erRight != nil {
		return 0, errors.New("[Error]: Can't parse the right value")
	}
	switch operator {
	case "*":
		return leftConverted * rightConverted, nil
	case "/":
		return leftConverted / rightConverted, nil
	case "+":
		return leftConverted + rightConverted, nil
	case "-":
		return leftConverted - rightConverted, nil
	default:
		return 0, errors.New("[Error]: No operators has been matched to : " + operator)
	}
}
