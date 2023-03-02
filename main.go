package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")

    //string
    var nameOne string = "testName1"
    var nameTwo = "testName2"
    var nameThree string

    nameThree = "testName3"

    fmt.Println(nameOne, nameTwo,nameThree)

    nameOne = "hi"
    nameTwo = "there"

    fmt.Println(nameOne, nameTwo)

    nameFour := "testName4"

    fmt.Println(nameFour)

    //Numbers

    var ageOne int = 20
    var ageTwo = 30
    ageThree := 40

    fmt.Println(ageOne, ageTwo, ageThree)

    //bits and memory

    // var numOne int8 = 25
    // var numTwo int8 =  -128
    // var numThree uint8 =  256            //Error-Number positve and less than equal to 355 are allowed
    // var numThree uint8 =  255

    // var scoreOne float32 = -1.2
    // var scoreTwo float64 = -1.2
    // scoreThree  := -1.2

}
