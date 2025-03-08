package main

import (
	"fmt"
	"sync"
)

// Go: Parity Distribution 

// Implement a function server which runs a server with the following 

// It accepts values through a channel.
// Each value is an int32. 
// You should describe a type 'in' and initialize a global variable serverChan 

// Depending upon whether the value is odd or even, it directs them to oddChan and evenChan 

// Ex:
// arr = [1,2,3,4,5]

// oddChan receives [1,3,5] 
// evenChan receives [2,4]
// The values are printed to STDOUT 

// func Server() {
//     // complete     
// }

type in int32 
var serverChan = make(chan in)

func Server() {
    oddChan := make(chan in)
    evenChan := make(chan in)
    
    go func() {
        for val := range serverChan {
            if val % 2 == 0 {
                evenChan <- val 
            } else {
                oddChan <- val 
            }
        }            
        close(oddChan)
        close(evenChan)
    }()

    // go func() {
    //     for val := range evenChan {
    //         fmt.Println("Even:",val)
    //     }
    // }()
    
    // go func() {
    //     for val := range oddChan {
    //         fmt.Println("Odd:",val)
    //     }
    // }()

}
