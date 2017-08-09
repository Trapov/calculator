package machine

func Process(expression string) (float64, error) {
	var answer, errorAnswer = solve(translate(expression))
	if errorAnswer != nil {
		return 0, errorAnswer
	} else {
		return answer, nil
	}
}
