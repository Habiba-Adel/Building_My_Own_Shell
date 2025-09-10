package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
  
)

//i will seperate the handling the built in commands in alone function cause both can use it and rather of duplication i will do that
//i pass the io.writer object to enable both functions call it without repeating 
func HandlingBuiltInCommands (command string, here io.Writer) bool{
  Args := strings.Split(command, " ")

	if Args[0] == "exit" {
    os.Exit(0)
		return true
	}

	//now lets start implementing the second built in command which is the cd
	if Args[0] == "cd" {
		//there is a function in os that chaning the current directory
		err := os.Chdir(Args[1])
		if err != nil {
			fmt.Fprintln(here,"This directory is not existed ^^ ")
		}
		return true
	}

	//now lets handle the environment variables
	if Args[0] == "export" {
		//we need here to read the name of the variable and the value and then passing it to save them in the environmet values
		//and there is just export command it just appear to us all the defines variables in this environment

		//we have 2 options here the first option if the user just write export that means he want to see all the defines variables
		if len(Args) == 1 {
			all := os.Environ()
			//there is a for each in go and that by using the range word
			for _, item := range all {
				fmt.Fprintln(here,item)
			}
		} else {
			//so that means the user enter name and value and we will need to split the second argument based on 0 and then send it to os to save it
			info := strings.Split(Args[1], "=")
			//now we will need to send the first and the second to the os to save them
			if len(info) != 2 { //that means the user entered invlaid things and that will lead to exception
				fmt.Println("Invalid formation for export command ^^ ")
				return true
			}
			os.Setenv(info[0], info[1])

		}
		return true
	}
	//now the second thing is when the user enter echo command so we will need to print to it the value of this key
	if Args[0] == "echo" {
		//there is 2 things here if the first character in the key name sis $ so we will print to it the val else so we will print to it the key again
		//if just echo word we will print empty line like real shells
		if len(Args) == 1 {
			fmt.Fprintln(here)
			return true
		}
		if Args[1][0] == '$' {
			fmt.Fprintln(here,os.Getenv(Args[1][1:]))
		} else {
			//we will just print the same key again
			fmt.Fprintln(here,Args[1])
		}
		return true

	}
	//now the third thing is to implement the unset and it will delete the mentioned key from the variables
	if Args[0] == "unset" { //its all cases in the reall shell it going us to the prompt again
		os.Unsetenv(Args[1])
		return true 
	}
  //otherwise is there is no any valid condtion for it that means it is an external one so
  return false

}



// first the pipe commands handling function
func HandlingPipe(line string) {
	//now we have the whole entered line without any whote spcae
	//now our first step is to splite it in array of strings based on |
// 	commands := strings.Split(line, "|")
// 	//and now each command from it will need to define its stdin and its std out and then split its command things to can pass it
// //first thing we need to define the pbject that will store the inital dat to it 
// previoushehe:=io.Reader(nil)
//   //we have 3 special cases here the firs tone cause it will not having previous inptu and the last one cause it will write in the os and the normal cases 
// 	for index, command := range commands {


// }
}

// second is the non pipe commands handling function
func HandlingNonPipe(line string) {
	//we will pass it first cause if is it built in so handle it and return 
  //cause it here in the non pipe version our result must appear directly in the terminal so that why we pass the os
  if HandlingBuiltInCommands(line,os.Stdout) == true {//that means it is a built in one so just return 
   return 
  }

	//using the exec is better very much rather than using the start process one 
  //now the second option is that the command is external one
  Args := strings.Split(line, " ")
	cmd := exec.Command(Args[0], Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("This command is not existed "+"😝 ")
	}
}

func main() {

	//lets implement REPL read eval print loop
	//it will be infinit loop
	for {
		//this is called prompot of the shell  and with the current working directory
		cwd, _ := os.Getwd()
    fmt.Print("🚀\033[1;31m<Habiba_Shell>\033[0m:" + cwd + " \033[1;31m$\033[0m ")//fprint here used to print what you write but in another extra thing like a file or something
		line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		line = strings.TrimSpace(line) //this is for removing the white spaces and inline characters from the command

		//now the first thing we will need to know if the entered command is pipe or no to know which func we will pass it to ypu
		//if we find | on the command that means pipe
		if strings.Contains(line, "|") {
			HandlingPipe(line)
		} else {
			HandlingNonPipe(line)
		}

	}

}
