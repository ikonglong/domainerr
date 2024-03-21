package domainerr

// RetryAdvice is the advice on retry for a Status.
type RetryAdvice string

var (
	// JustRetryFailingCall is for StatusUnavailable, the client can retry just the failing
	// call with exponential backoff. The minimum delay should be 1s unless it is documented
	// otherwise.
	JustRetryFailingCall = RetryAdvice("just_retry_failing_call")

	// RetryAtHigherLevel means that for StatusAborted, the client should retry at a higher level
	// (e.g., when a client-specified test-and-set fails, indicating the client should restart a
	// read-modify-write sequence). For StatusResourceExhausted, the client may retry at the higher
	// level with a delay determined by a sophisticated method.
	RetryAtHigherLevel = RetryAdvice("retry_at_higher_level")

	// NotRetryUntilStateFixed means that for StatusFailedPrecondition, the client should not retry
	// until the system state has been explicitly fixed. E.g., if a "rmdir" fails because the
	// directory is non-empty, `FAILED_PRECONDITION` should be returned since the client should not
	// retry unless the files are deleted from the directory.
	NotRetryUntilStateFixed = RetryAdvice("not_retry_until_state_fixed")

	// NoAdvice means that for all other status, retry may not be applicable - first ensure your request is idempotent.
	NoAdvice = RetryAdvice("no_advice")
)
