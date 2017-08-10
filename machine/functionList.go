package machine

import "errors"

type multiply struct{}
type divide struct{}
type plus struct{}
type minus struct{}

func (f multiply) Perform(left int64, right int64) (int64, error) {
    return left * right, nil
}
func (f divide) Perform(left int64, right int64) (int64, error) {
    if right != 0 {
        return left / right, nil
    } else {
        return 0, errors.New("[Error]: Divide by zero")
    }

}
func (f plus) Perform(left int64, right int64) (int64, error) {
    return left + right, nil
}
func (f minus) Perform(left int64, right int64) (int64, error) {
    return left - right, nil
}
