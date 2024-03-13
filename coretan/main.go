package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Mutex
*/

var (
	x     = 0
	mutex sync.Mutex
	wg    sync.WaitGroup
)

func main() {
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
		wg.Wait()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter :", x)
}

/*
	Goroutine
*/

// var (
// 	ch1        = make(chan string)
// 	sayHelloTo = func(who string) {
// 		var data = fmt.Sprintf("hello %s", who)
// 		ch1 <- data
// 		close(ch1)
// 	}
// )

// func print() {

// 	for i := 0; i < 5; i++ {
// 		fmt.Println((i + 1), <-ch1)
// 	}
// }

// func main() {
// 	runtime.GOMAXPROCS(2)

// 	go sayHelloTo("john wick")
// 	go sayHelloTo("jason bourne")
// 	go print()
// 	time.Sleep(5 * time.Second)

// }
// var ch1 = make(chan string)

// func doPrintChannel(channel string) {
// 	var data = fmt.Sprintf("data %s", channel)
// 	ch1 <- data
// }

// func doPrint(wg *sync.WaitGroup, message string) {
// 	defer wg.Done()
// 	fmt.Println(message)
// }

// func main() {
// 	runtime.GOMAXPROCS(2)

// 	go doPrintChannel("channel")
// 	fmt.Println(<-ch1)
// 	// time.Sleep(5 * time.Second)

// 	var wg sync.WaitGroup

// 	for i := 0; i < 5; i++ {
// 		var data = fmt.Sprintf("data %d", i)
// 		var data2 = fmt.Sprintf("data %d", -i)
// 		wg.Add(2)
// 		go doPrint(&wg, data)
// 		go doPrint(&wg, data2)
// 	}

// 	wg.Wait()

// }

/*
	Arguments and flag
*/
// var (
// 	name string
// 	age  int
// )

// func main() {
// 	var argsRaw = os.Args
// 	fmt.Printf("%v\n", argsRaw)

// 	// Define flags
// 	flag.StringVar(&name, "name", "John", "User's name")
// 	flag.IntVar(&age, "age", 30, "User's age")

// 	// Parse command-line arguments
// 	flag.Parse()

// 	// Access values of flags
// 	fmt.Printf("Name: %s \n", name)
// 	fmt.Printf("Age: %d \n", age)
// }

// // Salting pake waktu. time.Now().UnixNano()
// func doHashUsingSalt(text string) (string, string) {
// 	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
// 	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
// 	fmt.Println(saltedText)
// 	var sha = sha1.New()
// 	sha.Write([]byte(saltedText))
// 	var encrypted = sha.Sum(nil)
// 	return fmt.Sprintf("%x", encrypted), salt
// }

// func main() {
// 	var text = "this is secret"
// 	fmt.Printf("original : %s\n\n", text)
// 	var hashed1, salt1 = doHashUsingSalt(text)
// 	fmt.Printf("hashed 1 : %s\n\n", hashed1)
// 	// 929fd8b1e58afca1ebbe30beac3b84e63882ee1a
// 	var hashed2, salt2 = doHashUsingSalt(text)
// 	fmt.Printf("hashed 2 : %s\n\n", hashed2)
// 	// cda603d95286f0aece4b3e1749abe7128a4eed78
// 	var hashed3, salt3 = doHashUsingSalt(text)
// 	fmt.Printf("hashed 3 : %s\n\n", hashed3)
// 	// 9e2b514bca911cb76f7630da50a99d4f4bb200b4
// 	_, _, _ = salt1, salt2, salt3
// }

// // Task struct to represent a task
// type task struct {
// 	id         int
// 	title      string
// 	desc       string
// 	isComplete bool
// }

// // ToDoList struct to manage tasks
// type toDoList struct {
// 	tasks []task
// }

// // Method to add a task to the ToDo List
// func (list *toDoList) addTask(title, desc string) {
// 	taskID := len(list.tasks) + 1

// 	newTask := task{
// 		id:         taskID,
// 		title:      title,
// 		desc:       desc,
// 		isComplete: false,
// 	}

// 	list.tasks = append(list.tasks, newTask)
// }

// // Method to list all tasks
// func (list toDoList) listTasks() {
// 	for _, task := range list.tasks {
// 		fmt.Printf("%v. %s (%s) %v \n", task.id, task.title, task.desc, task.isComplete)
// 	}
// }

// // // Method to mark a task as completed
// func (list *toDoList) markCompleted(taskID int) {

// }

// // // Method to delete a task from the list
// // func (list *toDoList) deleteTask(taskID int) {

// // }

// func main() {
// 	//todolist instance
// 	myToDoList := toDoList{}

// 	//Add tasks
// 	myToDoList.addTask("Buy groceries", "Milk, Egg, Banana")
// 	myToDoList.addTask("Practice guitar", "For 25 hours a day")

// 	//List tasks
// 	myToDoList.listTasks()

// 	// Mark Completed
// 	myToDoList.markCompleted(1)
// }

/*
	STRUCT
*/

// type student struct {
// 	name  string
// 	grade float32
// }

// func main() {
// 	// var s1 student
// 	// s1.name = "Gibran Rakabuming Raka"
// 	// s1.grade = 2.3
// 	var s1 = student{"Gibran", 2.3}

// 	fmt.Printf("%v memiliki IPK %v.", s1.name, s1.grade)
// }

/*
	FUNCTION
*/

// func main() {
// 	var names = []string{"John", "Doe"}

// 	for i := 0; i < len(names); i++ {
// 		fmt.Printf("element %d : %s \n", i, names[i])
// 	}
// 	fmt.Println(printMessage("hello", names))
// }

// func printMessage(message string, arr []string) (string, string) {
// 	var nameString = strings.Join(arr, " ")
// 	return message, nameString
// }

/*
	VARIADIC FUNCTION
*/

// func main() {
// 	var konsumsiPanadolHarian = []int{1, 5, 3, 4, 67, 8, 42, 2}
// 	var totalKonsumsi = summation(konsumsiPanadolHarian...)
// 	fmt.Println(totalKonsumsi)
// 	var result = strings.Contains("wow", "awow")
// 	fmt.Printf("Letak %v pada memori adalah %v", result, &result)
// }

// func summation(numbers ...int) int {
// 	var total int = 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }
