// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

//     // CSV data already read into memory
//     filtered := [][]string{
//         {"id", "name", "website", "email"},      // header
//         {"1", "Alice", "google.com", "alice@gmail.com"},
//         {"2", "Bob", "amazon.com", "bob@yahoo.com"},
//         {"3", "Carol", "openai.com", "carol@openai.com"},
//     }

// 	// fmt.Println(filtered[1:])
//     // Loop through rows, skipping header
//     // for _, row := range filtered[1:] {
// 	// 	fmt.Println(row[3])
//     // }

// 	var result []any

// 	for i := 1; i < len(filtered); i++ {
// 		index := fmt.Sprintf("%d", i)
// 		value := strings.Join(filtered[i][1:], " ") // all columns except first
// 		sp := strings.Split(value, " ")
// 		last := sp[len(sp)-1] // last column

// 		lastObj := map[string]string{
// 			"Index": index,
// 			"last":  last,
// 		}
// 		result = append(result, lastObj)

// 	}
// 	fmt.Printf("%+v\n", result)

// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// IST location
// 	ist, _ := time.LoadLocation("Asia/Kolkata")

// 	// IST time
// 	istTime := time.Date(2026, 1, 21, 10, 0, 0, 0, ist)

// 	// Convert to UTC
// 	utcTime := istTime.UTC()

// 	fmt.Println("IST:", istTime)
// 	fmt.Println("UTC:", utcTime)
// 	fmt.Println(time.Now())
// }

package main

import (
	"fmt"
	// "time"
	// "sync"
)

// func main() {
// 	var numbers  = [3]int{10, 20, 30}
// 	fmt.Println("Array:", numbers)
// 	sum:=0
// 	// for i:=0;i<len(numbers);i++
// 	for index,value:=range numbers {
// 		// sum+=numbers[i]
// 		// fmt.Println(numbers[i])
// 		fmt.Println(index,value)
// 	}
// 	fmt.Println("sum :",sum)
// }

// func main() {
// 	num := [] int {1,2,3,4,5,6,7,8}
// 	for i,j := range num {
// 		fmt.Println(i,j)
// 	}
// 	fmt.Println(num)
// }

// func main() {
// 	marks := map[string]int{
// 		"maths":98,
// 		"science":96,
// 		"english":94,
// 	}
// 	fmt.Println(marks)
// 	for i,j := range marks{
// 		fmt.Println(i,j)
// 	}
// }

// func main() {
// 	a := 10
// 	p := &a

// 	fmt.Println("Value of a:", a)
// 	fmt.Println("Address of a:", p)
// 	fmt.Println("Value using pointer:", *p)
// }

// type Book struct{
// 	ID uint
// 	Name string
// 	Price int
// }

// func main() {
// var b Book
// b.ID = 209
// b.Name = "Golang"
// b.Price = 250
// fmt.Println(b)

// e := Book{ID : 202 , Name : "Python", Price : 260}
// fmt.Println(e)
// }

// type Student struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

// func main() {
// 	s := Student{1, "Asha", "asha@gmail.com"}
// 	fmt.Println(s)
// }

// type Vehicle interface {
// 	Start()
// }

// type Car struct{}
// type Bike struct{}

// func (c Car) Start() {
// 	fmt.Println("Car started")
// }

// func (b Bike) Start() {
// 	fmt.Println("Bike started")
// }

// func main() {
// 	vehicles := []Vehicle{Car{}, Bike{}}

// 	for _, v := range vehicles {
// 		v.Start()
// 	}
// }

// type Animal interface {
// 	Sound()
// 	Legs() int
// }

// type Dog struct{}

// func (d Dog) Sound() {
// 	fmt.Println("Bark")
// }

// func (d Dog) Legs() int {
// 	return 4
// }

// func main() {
// 	var a Animal = Dog{}
// 	a.Sound()
// 	fmt.Println(a.Legs())
// }

// type Student struct {
// 	name string
// }

// func greet()string{
// 	return "helo"
// }
// func main(){
// 	var s Student
// 	s.name = "prathi"
// 	fmt.Println(greet(),s)
// }

// func make_sound () {
// 	fmt.Println("bow-bow")
// }
// func main() {
// 	go make_sound()
// 	time.Sleep(1*time.Second)
// 	fmt.Println("i am dog")
// }

// func print_num(n int){
// 	fmt.Println(n)
// }
// func num (n int){
// 	// time.Sleep(2*time.Second)
// 	fmt.Println(n)
// }
// func main(){
// 	go print_num(67)
// 	go print_num(0)
// 	go print_num(-78)
// 	go num(6)
// 	time.Sleep(1*time.Second)
// }

// func printNum(n int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println(n)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(3)

// 	go printNum(67, &wg)
// 	go printNum(0, &wg)
// 	go printNum(-78, &wg)

// 	wg.Wait()
// 	fmt.Println("All done")
// }


// func task(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Task", id, "done")
// }

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(3)

// 	go task(1, &wg)
// 	go task(2, &wg)
// 	go task(3, &wg)

// 	wg.Wait()
// 	fmt.Println("All tasks completed")
// }


// func worker(n int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Worker", n, "finished")
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("All workers done")
// }

// func num(n int,wg*sync.WaitGroup){
// 	defer wg.Done()
// 	fmt.Println(n)
// }
// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(4)
// 	go num(1,&wg)
// 	go num(2,&wg)
// 	go num(3,&wg)
// 	go num(4,&wg)

// 	wg.Wait()
// 	fmt.Println("Done")
// }

// func main(){
// 	x:=10
// 	p:=&x

// 	fmt.Print(x,p,*p)
// }

// func addOne(n int) {
//     n = n + 1
// 	fmt.Println(n)
// }

// func main() {
//     x := 5
//     addOne(x)
//     fmt.Println(x)
// }

func addone(n *int){
	*n=*n+10
	fmt.Println(*n)
}
func main(){
	x:=10
	addone(&x)
	fmt.Println(x)
}
