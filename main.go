package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/immesys/spawnpoint/spawnable"
	bw2 "gopkg.in/immesys/bw2bind.v5"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [message] <num_repetitions>\n", os.Args[0])
		os.Exit(1)
	}

	var repetitions int
	var message string
	var err error
	parameters := spawnable.GetParamsOrExit()

	if len(os.Args) >= 3 {
		message = os.Args[1]
		repetitions, err = strconv.Atoi(os.Args[2])
	} else {
		repetitions, err = strconv.Atoi(os.Args[1])
		message = parameters.MustString("msg")
	}

	if err != nil {
		fmt.Println("Invalid repetitions argument:", err)
		os.Exit(1)
	}

	bwClient := bw2.ConnectOrExit("")
	bwClient.SetEntityFromEnvironOrExit()

	for i := 0; i < repetitions; i++ {
		output := fmt.Sprintf("%v: %s", i, message)
		po := bw2.CreateStringPayloadObject(output)
		bwClient.PublishOrExit(&bw2.PublishParams{
			URI:            parameters.MustString("to"),
			AutoChain:      true,
			PayloadObjects: []bw2.PayloadObject{po},
		})
		fmt.Printf("Publishing %d\n", i)

		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Sent %v messages. Terminating.\n", repetitions)
}
