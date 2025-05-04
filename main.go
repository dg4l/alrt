package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"flag"
)

func recompile(command string, filename string){
	cmd := exec.Command(command, filename)
	err := cmd.Run()
	if err != nil{ 
		fmt.Printf("ERROR: ")
		fmt.Println(err)
	}
}

func main(){
	filename := flag.String("f", "", "file which will be recompiled") 
	command := flag.String("c", "", "compiler to use") 
	flag.Parse()
	if *filename == ""{
		fmt.Printf("ERROR: NO FILE GIVEN AS INPUT\n")
		return
	}
	if *command == ""{
		*command = "pdflatex"
	}
	info, err := os.Stat(*filename)
	if err != nil{
		fmt.Println("ERROR: FILE INVALID")
		return
	}
	modtime := info.ModTime()
	for{
		info, err = os.Stat(*filename)
		if err != nil{
			fmt.Println(err)
			return
		}
		if info.ModTime() != modtime{
			recompile(*command, *filename)
			modtime = info.ModTime()
		}
		time.Sleep(1 * (time.Second / 4))
	}
}
