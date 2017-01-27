package main

import "fmt"
// Kunne ikke henvise til import av "log" pakke som ligger i samme mappe
// ./log henviser til search for pakker i mappen log
// Se mappestruktur for klarifisering eller ask Tor
import "./log"

func main()	{


	fmt.Println("Initiating")

	// Defines a value that cna be changed to send to Log2
	var logVal float64 = 2
	//Prints out the value sent to log2 and the answer
	fmt.Println("The answer from log2 from value:", logVal, "is", log.Log2(logVal))
	fmt.Println(allDone())



}

func allDone() string{
	//Just for teh keks
	return "Calculated and planned"

}
