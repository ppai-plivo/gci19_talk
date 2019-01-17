package main

// START OMIT
type numGoroutine struct{}

// String ensures that numGoroutine implements expvar.Var interface
// which expects String() to return a valid JSON value.
func (g *numGoroutine) String() string {
	return strconv.Itoa(runtime.NumGoroutine())
}

func foo() {
	// EITHER use an expvar map if you have multiple metrics.
	metrics := expvar.NewMap("runtime")
	metrics.Set("NumGoroutine", new(numGoroutine))

	// OR without expvar map if it's just one metric.
	expvar.Publish("NumGoroutine", new(numGoroutine))
}

// END OMIT
