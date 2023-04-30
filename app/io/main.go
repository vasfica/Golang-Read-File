package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	notesExec()
	commandInput()

}

func notesExec() {
	fmt.Println("            Hi, Welcome!                ")
	fmt.Println("Put files on folder location: files")
	fmt.Println("The program only execute CSV, JSON, and SQL Files", "\n")
	fmt.Println("These are commands used in various situations:", "\n")
	fmt.Println("la", "            List all files from directory files")
	fmt.Println("close", "         Close program")
	fmt.Println("----------------------------------------", "\n")
}

/*Function for user can input text on command line (CMD)*/
func commandInput() {
	/*Initialization buffer io with os.Stdin input*/
	scanner := bufio.NewScanner(os.Stdin)
	for {
		/*Print Enter text: */
		fmt.Print("Enter file name or command -> ")
		/*Execute bufio indeed os.Stdin*/
		scanner.Scan()
		/*Hold file name in fileName variable*/
		txt := scanner.Text()

		if txt != "" { //Check if text input not empty
			commandFilter(txt) //Go to function commandFilter(txt) parameter string
			// return             //return to exit function
		}

	}
}

/*commandFilter() is function used for filter command*/
/*If user typing command ex: close or la function direct to runCommand()*/
/*Else function direct to csv, sql, or json to execute file*/
func commandFilter(s string) {
	if strings.Contains(s, ".") { //Check if string contain dot "."
		x := strings.Split(s, ".")[1]                //Pull string after "."
		if x != "csv" && x != "sql" && x != "json" { //Check if string not contain format csv, sql, and json
			fmt.Println("There is no such file format") //Print error message
			return                                      //return use for exit function
		} else {
			files, err := ioutil.ReadDir("../files/")

			if err != nil {
				fmt.Println("Directory Not Found")
			}

			for _, j := range files { //Looping to check how much files in directory
				if j.Name() == s {
					if x == "csv" { //Check if input contain "csv" format string
						csvFormat(fmt.Sprintf("../files/%s", s))
					} else if x == "sql" { //Check if input contain "sql" format string
						sqlFormat(fmt.Sprintf("../files/%s", s))
					} else if x == "json" { //Check if input contain "json" format string
						jsonFormat(fmt.Sprintf("../files/%s", s))
					}
				} else {
					fmt.Println("File Not Found")
				}
			}

		}
	} else { //Else if string not contain dot "."
		runCommand(s) //Run function runCommand() with parameter string
	}

}

/*Function to check user input contain command list (ex: la or close)*/
func runCommand(s string) {
	if s != "la" && s != "close" { //Check if string not contain command la or close
		fmt.Println("Command Not Found") //Print error message

	} else { //Else
		if s == "la" { //If string contain "la" command
			commandListFiles() //Go to commandListFiles() function

		} else if s == "close" { //Else if string contain "close" command
			commandCloseProgram() //Go to commandCloseProgram()

		}
	}

}

/*Command for list all files in directory name "files"*/
func commandListFiles() {
	files, err := ioutil.ReadDir("../files/") //declaration ioutil for read directory on "../files/"

	if err != nil { //Check if there are no directory name: files
		fmt.Println("Directory Not Found") //Print error message
	}

	if len(files) == 0 { //Check if there are no file in directory files
		fmt.Println("There are no files in directory files") //Print error message
	}

	fmt.Println("----------") //Print boundaries for list files
	for _, j := range files { //Looping to check how much files in directory
		fmt.Println(j.Name()) //Print file name on directory
	}
	fmt.Println("----------") //Print boundaries for list files

}

/*Kill the process*/
func commandCloseProgram() {
	os.Exit(0) //Exit program
}

/*Read Excel File*/
func csvFormat(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Unable to open file: ", f.Name)
	}

	defer f.Close()

	reader := csv.NewReader(f)

	elements, err := reader.Read()

	if err != nil {
		log.Fatal(err)
	}

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for _, j := range elements {
		fmt.Print(j, "   ")
	}
	fmt.Print("\n")
	for _, j := range records {
		fmt.Println(j)
	}
}

/*Read SQL File*/
func sqlFormat(s string) {

}

/*Read JSON File*/
func jsonFormat(s string) {

}
