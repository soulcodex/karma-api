package utils

func RetryFunc(fn func() (interface{}, error), timesToRetry int) (interface{}, error) {
	var (
		err    error
		result interface{}
	)

	for i := 1; i <= timesToRetry; i++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}
	}

	return nil, err
}
