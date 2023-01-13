package solution

import "strconv"

type Result string

func IntResult(i int) Result {
	return Result(strconv.Itoa(i))
}
