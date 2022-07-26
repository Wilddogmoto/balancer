package main

import (
	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
	"log"
	"os"
	"time"
)

func main() {

	rep, err := runner.Run("Balancer.GetUrl", "localhost:44300",
		runner.WithProtoFile("balancer.proto", []string{"./server"}),
		runner.WithDataFromFile("data.json"),
		runner.WithInsecure(true),
		runner.WithTotalRequests(10000),
		runner.WithLoadDuration(time.Second*1),
	)

	if err != nil {
		log.Printf("error on run ghz: %v", err)
		return
	}

	p := printer.ReportPrinter{
		Out:    os.Stdout,
		Report: rep,
	}

	err = p.Print("pretty")
	if err != nil {
		log.Printf("error on print: %v", err)
		return
	}
}
