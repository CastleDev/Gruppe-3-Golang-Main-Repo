package main

import "fmt"
// Kunne ikke henvise til import av "log" pakke som ligger i samme mappe
// ./log henviser til search for pakker i mappen log
// Se mappestruktur for klarifisering eller ask Tor
import "./log"
import "bufio"
import "os"

func main()	{


	fmt.Println("Initiating")

	// Defines a value that cna be changed to send to Log2
	var logVal float64 = 10
	//Prints out the value sent to log2 and the answer
	//fmt.Println("The answer from log2 from value:", logVal, "is", log.Log2(logVal))

	reader := bufio.NewReader(os.Stdin)
	logVal2, _ := reader.ReadBytes('\n')
	fmt.Println(logVal2)


	if logVal == 2 {
				fmt.Println("The answer from log2 from value:", logVal, "is", log.Log2(logVal))
	}
	if logVal == 10 {
			fmt.Println("The answer from log10 from value:", logVal, "is", log.Log10(logVal))
	}
	fmt.Println(allDone())



}

func allDone() string{
	//Just for teh keks
	return "Calculated and planned"

}
