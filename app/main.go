package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//lets implement REPL read eval print loop
	//it will be infinit loop
	for {
		//this is called prompot of the shell
		fmt.Print("<Habiba_Shell>$ ") //fprint here used to print what you write but in another extra thing like a file or something
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command) //this is for removing the white spaces and inline characters from the command

//now we want to split the command to can get its arguments 
Args:=strings.Split(command," ")



		if command == "exit" {
			break
		}

		//now here we will start new process and this function is built in os package
		//func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)

		//first thing we will need to make an proceattr instance and tell him that this child process can treat with the stdin and std out and so on
		child := &os.ProcAttr{ //we need to put & cause it is pointer
			Dir:   "",                                         //this is related to the directory that the child will work on it and i want it to work from the parent directory
			Env:   nil,                                        //it is the environment variables and i will make it take the same of the parent too
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}, //this is to enable the child to can deal with the input and output from inside it and be careful the order here is very important
		}

    //it is must to give the ardg the command cause the first argument 0 is the commad name 
		// Args :=[]string{command}//to make the syntax easy to you you say args is arr of string from command 

     //it will throw an error if the command is not found in the external commands 
		proc, err := os.StartProcess(Args[0], Args, child) //now i start it

    if err != nil {//that means this command not an external command so we will print the error message
			fmt.Println(err) //print ln here its role is to print in the shell
      continue
		}

    //and this can throw error if there is anything make the child terminate in just bad way 
		 proc.Wait() //now my curr shell will wait until the child finish cause i am single threaded

	

	}

}




//remember to test clear command in your shell 