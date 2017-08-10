package machine

var operators = []string{"+-", "*/"}

func Process(expression string) (float64, error) {
    var rpnStack = translate(expression)
    var answer, errorAnswer = solve(rpnStack)

    if errorAnswer != nil {
        return 0, errorAnswer
    } else {
        return answer, nil
    }
}
