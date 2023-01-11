package common

import (
	"io"
	"os"
	"strconv"
)

func Input(name string) (io.ReadCloser, error) {
	input, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func GoldenInt(name string) (int, error) {
	golden, err := os.Open(name)
	if err != nil {
		return 0, err
	}
	defer golden.Close()

	b, err := io.ReadAll(golden)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(string(b))
	if err != nil {
		return 0, err
	}

	return i, nil
}
