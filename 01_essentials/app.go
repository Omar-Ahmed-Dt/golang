package main

/* 
main is a special Package name that tells GO that this Package will be the main entry point 
of the application we're building (tell go where execution should start). 

- How run Go: 
$ go mod init <module name>    >> this will turn this project into a module
one moulde consists of multiple packages. a module is simply a Go project. 
$ go run <file name>
or  
$ go build   >> will create executalbe file that could be executed on systems that don't have Go installed.

Go programs are organized into packages. The main package is a special package that is used to create an executable program. 
Any Go program that you want to compile and run as a standalone application must use the main package
*/


import(
    "fmt"   //standard library
    "math"
)

func main() {
/* 
main func, this also must be named main so that Go knows which code to execute when the app starts,
and which code should be executed first.

The main() function is the entry point of every Go program. When you run a Go program, 
the Go runtime looks for the main() function in the main package to start the execution

- package main: Declares that this is the main package of the application.
- func main(): The starting point of the program. When you run this program, it prints "Hello, World!" to the console. one main func in the file.
*/

    // To Print text: 
    fmt.Print("Investment Amount: ") 

    /*
    declaring a variable: 
        1- declaration without initialization: var foo string
        2- Declaration with initialization: var foo string = "Go is awesome"
        3- Multiple declarations:
            var foo, bar string = "Hello", "World"

            // OR

            var (
            	foo string = "Hello"
            	bar string  = "World"
            )
        4- Type is omitted but will be inferred: var foo = "What's my type?"
        5- Shorthand declaration , we omit var keyword and type is always implicit: foo := "Shorthand!"
        Note: Shorthand only works inside function bodies.
    */

    var investmentAmount float64

    // To allow user to enter the value of variable: 
    fmt.Scan(&investmentAmount)

    fmt.Print("Expected Retrun Rate: ")
    var expectedReturnRate float64
    fmt.Scan(&expectedReturnRate)

    var years float64 = 10

    futureValue := investmentAmount * math.Pow ( 1 + expectedReturnRate / 100 , years )

    const inflationRate = 2.5 
    futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

    fmt.Println("Future Value: ", futureValue)
    fmt.Println("Future Real Value: ",futureRealValue)

    fmt.Printf("Future Vale: %v\nFuture Real Value: %v", futureValue, futureRealValue)

}
