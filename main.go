package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	bw2 "gopkg.in/immesys/bw2bind.v5"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <num_repetitions>\n", os.Args[0])
		os.Exit(1)
	}

	repetitions, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid repetitions argument:", err)
		os.Exit(1)
	}

	bwClient := bw2.ConnectOrExit("127.0.0.1:28589")
	bwClient.SetEntityFileOrExit("entity.key")

	paramContents, err := ioutil.ReadFile("params.yml")
	if err != nil {
		fmt.Println("Failed to read parameters file:", err)
		os.Exit(1)
	}
	parameters := make(map[string]string)
	err = yaml.Unmarshal(paramContents, &parameters)

	for i := 0; i < repetitions; i++ {
		msg := fmt.Sprintf("%v: %s", i, parameters["msg"])
		po := bw2.CreateTextPayloadObject(bw2.PONumText, msg)
		bwClient.PublishOrExit(&bw2.PublishParams{
			URI:            parameters["to"],
			AutoChain:      true,
			PayloadObjects: []bw2.PayloadObject{po},
		})

		time.Sleep(1 * time.Second)
	}

    fmt.Printf("Sent %v messages. Terminating.\n", repetitions)
}
