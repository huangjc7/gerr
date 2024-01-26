# Gerr
An error handling library that makes code simpler<br>
catches errors explicitly<br>
Goodbye `if err != nil`

# Install
```go
go get -u github.com/huangjc7/gerr@latest
```

# Use
1. If your program is executed once and then terminated, use the following method:
```go
func main() {
	// Custom error handling logic
	customErrorHandler := func(err error) {
		log.Printf("Error:[%s]\n", err)
	}

	// When shouldWait is true, waitGroup will be used to facilitate the scenario where the goroutine has not completed execution when the function exits.
	// When shouldWait is false, waitGroup will not be used to continue receiving errors from the error channel.
	// Note: When shouldWait is false, there is no need to call the Close method

	g := gerr.New(customErrorHandler, true)
	// Start goroutine to receive error messages
	g.Receive()

	err := fmt.Errorf("this is an error")
        // You can capture error information anywhere in your code
	g.CatchError(err)
	g.Close()
}
```
2. If your program runs continuously, use the following method:
```go
func main() {
	// Custom error handling logic
	customErrorHandler := func(err error) {
		log.Printf("Error:[%s]\n", err)
	}

	// When shouldWait is true, waitGroup will be used to facilitate the scenario where the goroutine has not completed execution when the function exits.
	// When shouldWait is false, waitGroup will not be used to continue receiving errors from the error channel.
	// Note: When shouldWait is false, there is no need to call the Close method

	g := gerr.New(customErrorHandler, true)
	// Start goroutine to receive error messages
	g.Receive()

	err := fmt.Errorf("this is an error")
	for {
                // You can capture error information anywhere in your code
		g.CatchError(err)
	}
}
```

