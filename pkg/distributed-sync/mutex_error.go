package distributed_sync

const (
	errorLockingKeyMessage   = "Some problem ocurred while acquiring processes"
	errorReleasingKeyMessage = "Some problem ocurred while releasing processes"
)

type ErrorMutex struct {
	message    string
	identifier string
	previous   error
}

func (i ErrorMutex) Error() string {
	return i.message
}

func NewErrorLockMutexKey(identifier string, previous error) *ErrorMutex {
	return &ErrorMutex{message: errorLockingKeyMessage, identifier: identifier, previous: previous}
}

func NewErrorReleaseLockMutexKey(key string, previous error) *ErrorMutex {
	return &ErrorMutex{message: errorReleasingKeyMessage, identifier: key, previous: previous}
}
