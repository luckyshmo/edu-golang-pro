package main

// #include <stdio.h>
// void callC() {
// printf("Calling C code!\n");
// }
import "C"

import "fmt"

func goToC() {
	fmt.Println("I'm Go")
	// C.callC()
	fmt.Println("Go again")

	//much more on pages 86-91
}
