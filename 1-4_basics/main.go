/*
----------------- 16/01/26 ----------------

----------------  Lec 2 -- 5 reasons to choose Golang  ----------------

1. Fast build time -- program/ app has to be build before deployment in CI/CD, fast means it builds quickly and consumes less CI/CD resources, fast test and ship

2. Fast startup -- how much time it takes for the application from deployment to get up and running to accept requests, it matters not only for the first time but also for auto scaling it needs quick restart. bcz go program compiles in machine code and thus runs fast.

3. performance and efficiency -- Go uses CPU and other resources of server very less. bcz of goroutines (light weight threads)

4. concurrency model -- goroutines uses all threads of cpu

5. static typing and compilation -- this is good as compared to py, error will be caught in compilation not in production, also size of compilation file (.exe) is less.

Go is case sensitive

*/

//  ----------------  Lec 4 -- First Program  ----------------

package main // package -- logical group of files of code, every go program has main package

import "fmt"

func main() { // main fn is entry pt. in go program
	fmt.Println("hello world") // fmt is package of std. lib. used for formatting which has Println as method
}

/*
1. programs can be either saved directly or inside folder and the file name can be given as main.go

2. for building/ compiling (.exe) -- go build <path_of_file> (main.go)

3. builds binary/ machine code (.exe) of small size

4. to run the machine code(main) in windows -- write "main.exe" in cmd, in mac/ linux "./main" but cmd have to be open from that same dir. else have to write complete path

5. direct build + run -- go run main.go --> internally written in go it doesn't create binary file. just executes internally might be creating temp .exe then releases

6. you should generally build it and run the binary directly when using it in production.
*/


/* 

Naming Convention 

In Go, the file naming convention is to use short, all-lowercase names, with underscores (_) to separate multiple words when necessary. Can start with num but check more detail later


*/