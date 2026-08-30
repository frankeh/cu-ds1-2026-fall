package tutorial

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/*
 * COMS 4113 Go Tutorial
 * Originally from Mingen Pan
 * Adapted by Jay Karp, Xindi Xu, Pavan Kalyan Damalapati
 */

 // Comments can either be written like this
 /* 
  * Or like this
  */


  /*
   * To run any of these test cases, open up your terminal
   * in this folder and run `go test -run Test<func>`
   * any function that starts with Test can be substituted
   */

  /********************/
  /* Basic Data Types */
  /********************/

  func TestVariables(t *testing.T) {
	  // Define variable
	  // name and type, note that it's different from C.

	  var a int

	  // var (
	  // 	a int
	  // 	b string
	  // )
	  // var a, b int


	  // Type inference: Go compiler will figure out 
	  // the type for you based on the value

	  // var b int = 2
	  // var b = 2
	  b := 2

	  // Once a variable is defined, it cannot be redefined, only assigned
	  // b := 3 // Error: no new variables on left side of :=
	  b = 3 // valid

	  a = b + 1
	  fmt.Printf("a: %d, b: %d\n", a, b);
  }

  /********************************/
  /* Basic Flow Control */
  /********************************/

  func TestControl(t *testing.T) {
	  a := 1

	  // If-else statement
	  // No need to use parentheses
	  if a > 0 {
		  fmt.Println("a is > 0")
	  } else {
		  fmt.Println("a is <= 0")
	  }

	  b := 0

	  for i := 0; i < 10; i++ {
		  b++
	  }
	  fmt.Printf("b is now %d\n", b)

	  i := 0
	  for i < 10 {
		  b++
		  i++
	  }
	  fmt.Printf("b is now %d\n", b)

	  // Infinite loop, use carefully
	  // for {
		 //  // do something
		 //  break
	  // }
  }

  /*****************************/
  /* Important Data Structures */
  /*****************************/

  func TestDS(t *testing.T) {
	  // static array [size]type{elements,...}
	  arr := [3]int{1, 2, 3}

	  for i, num := range arr {
		  fmt.Printf("arr[%v]= %v\n", i, num)
	  }


	  // Slice (dynamic array) - think of it as a data structure with 3 fields:
	  //   1. pointer to an el in the array (any el)
	  //   2. length
	  //   3. capacity
	  // Creates an array (continues segment of memory) and a slice
	  slice := make([]int, 5, 10) // len = 5 (initialized), cap = 10 (allocated slots)
	  fmt.Printf("len(slice) = %d, but capacity is %d\n", len(slice), cap(slice))
	  
	  // How to increase the capacity of a slice?
	  // Use append
	  slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	  fmt.Printf("len(slice) = %d, but capacity is %d\n", len(slice), cap(slice))

	  // map -> extremely important
	  myMap := make(map[string]string)
	  myMap["a"] = "apple"

	  myMap = map[string]string{
		  "a": "apple",
		  "b": "banana",
		  "c": "cat",
	  }

	  for key, val := range myMap {
		  fmt.Printf("myMap[%s] = %s\n", key, val)
	  }
	  // Note: map doesn't guarantee the order of iteration
  }

  /*****************************/
  /* Go Routines */
  /*****************************/

  func TestRoutine(t *testing.T) {
	  arr := []int{5, 3, 2, 1, 4}

	  // Notice the lambda function
	  // []int is a reference, so no need to use a pointer
	  sleepSort := func(arr []int) {
		  for _, num := range arr {
			  go sleepAndPrint(num)
		  }
	  }

	  sleepSort(arr);
	  // Comment out this line and observe the result
	  time.Sleep(1 * time.Second)
	  fmt.Println()
  }

  func sleepAndPrint(num int) {
	  time.Sleep(time.Duration(num) * time.Millisecond)
	  fmt.Printf("%d ", num)
  }

  /*****************************/
  /* Channels */
  /*****************************/

  func TestChannel(t *testing.T) {
	  arr := []int{5,3,2, 1, 4}

	  sleepSort := func(arr []int) {
		  ch := make(chan int, len(arr))

		  for _, num := range arr {
			  go sleepAndReturn(num, ch)
		  }

		  for i:=0; i< len(arr); i++ {
			  num := <- ch
			  fmt.Printf("%d ", num)
		  }
	  }
	  sleepSort(arr)
	  fmt.Println()
  }

  func sleepAndReturn(num int, ch chan int) {
	  time.Sleep(time.Duration(num) * time.Millisecond)
	  ch <- num
  }

  /*****************************/
  /* Structs */
  /*****************************/

  type Entry struct {
	  // Capitalized field names are exported (public)
	  Index int
	  // Lowercase field names are not exported (private), only accessible within the package
	  val string
  }

  // Go will secretly pass in the struct as an argument
  func (entry Entry) incrementIndexOfCopy() {
	  entry.Index++
  }

  // Go will secretly pass in the pointer to the struct as the argument
  func (entry *Entry) incrementIndexOfRef() {
	  entry.Index++
  }

  // More Traditional way of defining methods
  func incrementIndexOfRef(entry *Entry) {
	  entry.Index++
	  // Notice that we use the dot operator for accessing fields for both structs and pointers. 
	  // This might be confusing for those with C backgrounds
	  // actually, the compiler will automatically convert the pointer to the struct
	  // (*entry).Index++ is how the compiler sees it.
  }

  func TestPointer(t *testing.T) {
	  entry := Entry {
		  Index: 0, 
		  val: "val",
	  }

	  // Initial value
	  fmt.Printf("Initial entry.Index = %d\n", entry.Index)

	  // Update copy of entry
	  entry.incrementIndexOfCopy()
	  fmt.Printf("After updating copy entry.Index = %d\n", entry.Index)

	  // Update the actual struct
	  entry.incrementIndexOfRef()
	  fmt.Printf("After updating the actual entry.Index = %d\n", entry.Index)

	  // Update the actual struct again but using a normal function
	  incrementIndexOfRef(&entry)
	  fmt.Printf("Updating the actual entry.Index = %d\n", entry.Index)
  }

  /*****************************/
  /* Type Casting */
  /*****************************/

  func TestTypeCasting(t *testing.T) {
	  var i interface{} = "hello"

	  s := i.(string)
	  fmt.Println(s)

	  s, ok := i.(string)
	  fmt.Println(s, ok)

	  f, ok := i.(float64)
	  fmt.Println(f, ok)
	  
	  // f = i.(float64) // panic
	  // fmt.Println(f)

	  printType(10)
	  printType("abc")
	  printType(nil)
	  printType(10.2)
  }

  func printType(i interface{}) {
	  switch i.(type) {
	  case int:
		  fmt.Println("int")
	  case string:
		  fmt.Println("string")
	  default:
		  fmt.Println("unknown")
	  }
  }

  /*****************************/
  /* Locking */
  /*****************************/

  func TestLock(t *testing.T) {
	  lock := sync.Mutex{}

	  for i := 0; i < 10; i++ {
		  // There will be no dead lock don't worry
		  goodLockHabit(&lock)
	  }

	  fmt.Println("Correct locking habits work")

	  for i :=0; i< 10; i++ {
		  badLockHabit(&lock)
	  }
	  fmt.Println("This will never be printed because of deadlock")
  }

  func goodLockHabit(lock *sync.Mutex) int {
	  lock.Lock()
	  defer lock.Unlock()

	  // Early stop
	  // Don't have to unlock before every return because it's handled by `defer`

	  if rand.Int() % 2 ==0 {
		  return 1
	  }

	  // Do some other stuff
	  time.Sleep(500 * time.Millisecond)

	  return 0;
  }

  func badLockHabit(lock *sync.Mutex) int {
	  lock.Lock()

	  // Early stop
	  if rand.Int() % 2 ==0 {
		  return 1
	  }

	  time.Sleep(500 * time.Millisecond)

	  lock.Unlock()
	  return 0
  }


  /*****************************/
  /* Recommended IDEs */
  /*****************************/

  /*
  - Visual Studio Code
  	- Good remote editing support, good go support
  - GoLand
  	- Good testing support, similar to Intellij
  - (Neo)Vim
  	- very convenient on remote machines
  - Other Editors + CLI
  */

  /*****************************/
  /* Recommended Resources */
  /*****************************/
  /*
  - Go Tour: https://go.dev/tour/welcome/1
  - Go by Example: https://gobyexample.com/
  - Effective Go: https://golang.org/doc/effective_go.html
  */


