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

//switch case, depending on the case it will print out the necessary one
switch yourAge {
case 16 : fmt.println('go outside')
case 18 : fmt.println('go inside')
default fmt.println('go have fun')
}

// array, holds a fixed number of values of the same data type
// defining favnumbers as an array with a length of five, yet those have not been defined
var favNumbers[5] float64
// here they are defined
// the zero is saying in the 0 index the first fave number is defined at 162
favNumbers[0] = 162
favNumbers[1] = 11
favNumbers[2] = 13
favNumbers[3] = 1989
favNumbers[4] = 52

// another way to define an array
// more fave is an array with a length of five numbers and the numbers are defined in the brackets
moreFaveNumbers := [5]int {11, 13, 1989, 4, 52}

//then to iterate index and value will tell you the value of that index.
// when using range, two value are returned for each iteration the first is the index and the second
// is a copy of the element of that index. 
for i, value := range moreFaveNumbers {
	fmt.println(value, i)
}

// and if you didnt want one of the things to be returned all you would need to do is this synta
for _, value := range moreFaveNumbers {
	fmt.println(value)
}

//slices are like arrays but whenevr you define them you leave out the size. 
numSlice := []int {5,4,3,2,1}

// here we are creating a new slice and adding values to our new slice variable
// the five will go to the fourth spot because as we know arrays start at 0
// and the new slice's zero index will be the first number from the old slice, ie: 2;
numSlice2 :=  numSlice[3:5]
fmt.println("numSlice2[0] =", numSlice2[0]);
// defining a slcie without a set of values 
// make is defining the slice of type you want, so its an array of ints
// the five is saying that the defualt value of the first five ints will be 0 
// and the ten is saying that the max length of the array is ten
numSlice3 := make([]int, 5, 10)
//if we want to add new values to this slice, we will use append
//the variable that you put right after the append is the variable that you want to 
//append new values to 
numSlice3 = append(numSlice3, 0, 1)
fmt.println(numSlice3[6]); // will print out the one
//so here we a creating an array of strings, and since the default value is being set to 0
// we are creating an empty array of strings
stringSlice := make([]string, 0)
//maps a collection of key value pairs. you're going to have a key associated with a value
// inside the [] will be the data type for the key and the int will be the value for the values
// we are going to store
presAge := make(map[string] int)

presAge["barack obama"] = 42
fmt.println(presAge["barack obama"]) // this will spit out 42
// so you put in the variable you want to use, then the key in which you want information from
// which is barack in this case, and equals 42
// we can then do something like this to delete a key
presAge['kennedy'] = 43
delete(presAge, 'kennedy')

//functions
listOfNums :=[]int{1,2,3,4,5}
//param of ints and the int outside of the bracket is int saying, we are expecting
// and int to be returned aka the return type
//range again is going to cycle through whatever the paramter is and get the values
func addThemUp(numbers []int) int {
	sum := 0;
	for _, val := range numbers {
		sum += val
	}
	return sum
}
//the for loop is providing you the value of the numbers. 
fmt.println(addThemUp(listOfNums));

//here we are creating a function that is going to return 2 values



func main() {
	fmt.Printf("henlo, world\n")




}