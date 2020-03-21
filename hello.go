package main

import "fmt"

// variables are statically typed and type cant change once its been defined
var age int = 40;
var favNum float64 = 1.618;
randNum := 1;
const pi float64 = 3.14;
var (
	varA = 2
	varB = 3
)
var myName string = "philip loeffler"
//lens returns the length of the string
fmt.Println(len(myName))
//\n creates a new line
fmt.Println("i like new lines \n \n")

var isOverForty bool = false;
var isUnderForty bool = true;

// using the percent sign with a letter will print out specific things
// based on the information you want, you can look these up in the future,
// but just remember they are there to grab info about the string or
// integer that you 
fmt.Println(%d %b %c %x %e);
//logical operators
&& // and
|| // or 
! // not symbol

// for loops

i := 1; 
//for loops, will print i on the screen check itself then print again up to ten
for (i <= 10) {
	fmt.Println(i);
	i ++
}
//different way to write the for loop
for j := 0; j < 5; j ++ {
	fmt.println(j);
}
// if statements
// one thing to remember is if the first condition is met, you wont see any of 
// the other conditions inside the if statement
yourAge := 30;
 
if yourAge  >= 38 {
	fmt.Println('hooray youre old')
} else  if{
	fmt.Println('hooray youre young')
} else {
	fmt.Println('hooray just hooray')
}

func main() {
	fmt.Printf("hello, world\n")




}