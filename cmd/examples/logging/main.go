package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {
	log.Println("Standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	myLog := log.New(os.Stdout, "my:", log.LstdFlags)
	myLog.Println("from my Log")

	var buf bytes.Buffer
	bufLog := log.New(&buf, "buf:", log.LstdFlags)

	bufLog.Println("hello")

	fmt.Print("from buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there!")

	myslog.Info("hello there", "key", "val", "age", 25)

}
