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


/*
you can add constant and variable declarations outside of any functions. And when you do that, that constant or variable is scoped to the entire code file instead of a specific function and is therefore available in all the functions of that file.
*/
const inflationRate = 2.5
var p1 float64
var p2 float64


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

    var years float64 = 10  // this variable is scoped to the main func

    futureValue := investmentAmount * math.Pow ( 1 + expectedReturnRate / 100 , years )

    // const inflationRate = 2.5 
    futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

    // outputs info: 
    // 1- using println:
    // fmt.Println("Future Value: ", futureValue)
    // fmt.Println("Future Real Value: ",futureRealValue)
    // or
    // 2- using printf: 
    fmt.Printf("Future Value: %v\nFuture Real Value: %v\n", futureValue, futureRealValue)
    // %v : a placeholder to output the value
    // \n : a new line

    fmt.Printf("Future Value: %.2f\nFuture Real Value: %.1f\n", futureValue, futureRealValue)
    // %.2f : how many decimal places you want output 4.78

    fmt.Print("////////////////////////////////////////////////\n")

    // Created Formatted Strings:
    formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
    formattedRFV := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)
    fmt.Print(formattedFV, formattedRFV)

    fmt.Print("////////////////////////////////////////////////\n")

    // Building Multiline Strings using backtick: 
    fmt.Printf(`Future Vale: %v\n
    Future Real Value: %v\n`, futureValue, futureRealValue) 

    fmt.Print("////////////////////////////////////////////////\n")

    // Called funcTest:
    printfuncText := funcTest(2,3)
    fmt.Println(printfuncText)


   // If Statements: 
   var xif int 
   fmt.Print("input xif value: ")
   fmt.Scan(&xif)
   if xif > 5 {
		fmt.Println("xif is gt 5")
	} else if xif > 10 {
		fmt.Println("xif is gt 10")
	} else {
		fmt.Println("else case")
	}

    // Nested If statement
    age := 26
    if age > 18 {
        if age >= 21 {
            fmt.Println("You are an adult and can legally drink.")
        } else {
            fmt.Println("You are an adult but cannot legally drink.")
        }
    } else {
        fmt.Println("You are a minor.")
    }


}

// Functions:

// func funcTest(text string)  inflationRate float64 {    

    // text is a parameter for this func
    // we need to specify the type for return variable: inflationRate float64
    // we need to specify the type for the parameter: text string

// func funcTest(text string) ( inflationRate float64, parameter valuetype, .... ) { 
func funcTest(p1 float64 , p2 float64) (p3 float64) {
    // Body Func:
    p3 = p1 + p2
    return p3
}