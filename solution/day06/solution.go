package day06

import (
	"container/list"
	"io"

	"github.com/pvlrmnnk/advent-of-code-2022/solution"
)

type slidingWindow struct {
	uniqueCnt int
	queue     *list.List
}

func (w *slidingWindow) add(b byte) {
	w.queue.PushFront(b)
	if w.queue.Len() > w.uniqueCnt {
		w.queue.Remove(w.queue.Back())
	}
}

func (w *slidingWindow) hasUniqueVals() bool {
	if w.queue.Len() < w.uniqueCnt {
		return false
	}

	seen := make(map[byte]struct{}, w.queue.Len())
	el := w.queue.Front()
	for el != nil {
		seen[el.Value.(byte)] = struct{}{}
		el = el.Next()
	}

	return len(seen) == w.uniqueCnt
}

func newWindow(uniqueCnt int) *slidingWindow {
	return &slidingWindow{
		uniqueCnt: uniqueCnt,
		queue:     list.New(),
	}
}

func SolutionPt1(input solution.Input) solution.Result {
	data, _ := io.ReadAll(input)

	window := newWindow(4)
	var idx int
	for idx = 0; idx < len(data); idx++ {
		window.add(data[idx])
		if window.hasUniqueVals() {
			break
		}
	}

	return solution.IntResult(idx + 1)
}

func SolutionPt2(input solution.Input) solution.Result {
	data, _ := io.ReadAll(input)

	window := newWindow(14)
	var idx int
	for idx = 0; idx < len(data); idx++ {
		window.add(data[idx])
		if window.hasUniqueVals() {
			break
		}
	}

	return solution.IntResult(idx + 1)
}
