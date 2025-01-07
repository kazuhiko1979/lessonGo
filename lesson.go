// Go:build main

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LogginSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogfile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogfile)

}

func main() {
	LogginSettings("test.log")
	_, err := os.Open("fafafa")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("logging!")
	log.Printf("%T %v\n", "test", "test")

	log.Fatalf("%T %v", "test", "test")

	log.Fatalln("error!")

	fmt.Println("ok")

}
