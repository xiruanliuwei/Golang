
package main

import "fmt"
import "io"    // package io provides basic interfaces to I/O primitives
import "net"   // package net provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution,
			   // and Unix domain sockets.
import "log"
import "os"

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {  // As with for, parentheses are never used around the condition in an if statement,
												  // but braces are required for the body
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	// to print the type of done
	fmt.Printf("type of done is: %T\n", done)

	// fmt.Printf("type of struct{} is: %T\n", struct{})  // It is incorrect, because "type struct {} is not an expression"
	fmt.Printf("type of struct{}{} is: %T\n", struct{}{})
	
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	} ()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}