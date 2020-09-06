package main

import (
	"chatbot.com/pkg/grpcserver"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cleanUpChan := make(chan int,1)
   go  grpcserver.StartServer(cleanUpChan);
   stopChan := make(chan os.Signal, 2)  
   signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
   {
	   fmt.Println(<-stopChan)
	   fmt.Println("exiting ")
	grpcserver.CleanUpServer(cleanUpChan)
	   fmt.Println(<-cleanUpChan)
      fmt.Println("cleanup/saving users data completed")
	   

   }

}
