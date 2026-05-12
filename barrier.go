package writebarrier

// IsEnabled returns whether the write barrier is currently enabled.
// Note that result might be stale by the time the function returns.
func IsEnabled() bool
