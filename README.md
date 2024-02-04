# Gerr

An error handling library that makes code simpler<br>
catches errors explicitly<br>
Goodbye err != nil

# Install

```go
go get -u github.com/huangjc7/gerr@latest
```

# Use

1. If your program is executed once and then terminated, use the following method:

```go

func main() {
	// When shouldWait is true, waitGroup will be used to facilitate the scenario where the goroutine has not completed execution when the function exits.
	// When shouldWait is false, waitGroup will not be used to continue receiving errors from the error channel.
	// Note: When shouldWait is false, there is no need to call the Close method

	g := gerr.New(true)

	// Start goroutine to receive error messages
	g.Receive()

	//Register an error callback function to perform business logic processing when catching errors.
	fmtProcessFunc := func(err error) {
		fmt.Println("Error", err)
		//Business processing logic
	}

	logProcessFunc := func(err error) {
		log.Printf("Error:[%s]\n", err)
		return
	}

	//New test error
	err1 := errors.New("this is an error 1")
	err2 := errors.New("this is an error 2")

	// Use LogProcessFunc to handle errors
	g.CatchError(fmtProcessFunc, err1)
	g.CatchError(logProcessFunc, err2)

	g.Close()

}
```

2. If your program runs continuously, use the following method:

```go

func main() {
	// When shouldWait is true, waitGroup will be used to facilitate the scenario where the goroutine has not completed execution when the function exits.
	// When shouldWait is false, waitGroup will not be used to continue receiving errors from the error channel.
	// Note: When shouldWait is false, there is no need to call the Close method

	g := gerr.New(false)

	// Start goroutine to receive error messages
	g.Receive()

	//Register an error callback function to perform business logic processing when catching errors.
	fmtProcessFunc := func(err error) {
		fmt.Println("Error", err)
		//Business processing logic
	}

	logProcessFunc := func(err error) {
		log.Printf("Error:[%s]\n", err)
		return
	}

	//New test error
	err1 := errors.New("this is an error 1")
	err2 := errors.New("this is an error 2")

	for {
		// Use LogProcessFunc to handle errors
		g.CatchError(fmtProcessFunc, err1)
		g.CatchError(logProcessFunc, err2)
	}
}
```

3. Recommended example<br>

   Very simple to use, no if != nil in the code

```go

// BusinessLogic struct contains business logic
type BusinessLogic struct {
    errorHandler *gerr.Error
}

// Callback function for error handling logic
func (b *BusinessLogic) FmtProcessFunc(err error) {
    fmt.Println("Error in business logic:", err)
}

// Callback function for error handling logic
func (b *BusinessLogic) LogProcessFunc(err error) {
    log.Printf("Error in business logic: [%s]\n", err)
}

// NewBusinessLogic creates a BusinessLogic instance
func NewBusinessLogic(errorHandler *gerr.Error) *BusinessLogic {
    return &BusinessLogic{
        errorHandler: errorHandler,
    }
}

// PerformOperation performs an operation, which may produce an error
func (b *BusinessLogic) PerformOperation() {
    // Here we simulate some operations that may produce errors
    err := errors.New("operation error")
    // Catch errors and if there are any, use the specified callback function for handling
    b.errorHandler.CatchError(b.LogProcessFunc, err)
    b.errorHandler.CatchError(b.FmtProcessFunc, err)
}

func main() {
    g := gerr.New(false)

    // Start a goroutine to receive error messages
    g.Receive()

    // Create a BusinessLogic instance
    bl := NewBusinessLogic(g)

    // Simulate the execution of business logic
    for {
        bl.PerformOperation()
    }
}
```
