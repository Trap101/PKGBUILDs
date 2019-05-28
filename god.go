package main

import (
	"io"
    "fmt"
	"os"
	"flag"
	"bytes"
	"os/exec"
	"strings"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/common-nighthawk/go-figure"
	"github.com/carmark/pseudo-terminal-go/terminal"
)

var version string = "1.0"
var outb, errb bytes.Buffer // Buffer for command output
var cmdSlice, cmdList []string // Arrays for the shortened git commands
var whiteColor (*color.Color) = color.New(color.FgWhite, color.Bold)

// Executes the terminal command and returns output.
// stdout parameter determines the output stream.
func execCmd(input string, stdout bool) string {
    // Remove the newline character.
    input = strings.TrimSuffix(input, "\n")
	// Prepare the command to execute.
	// sh used for handling the command parameters.
	// Otherwise, exec library takes the parameters
	// as argument which is something that we don't
	// want due to the complexity of git commands.
	cmd := exec.Command("sh", "-c", input)
	// Set the correct output device.
	if stdout{
    	cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
	}else{
		cmd.Stdout = &outb
		cmd.Stderr = &errb
	}
	// Execute the command and return
	// error or output depending on the
	// stdout parameter.
	err := cmd.Run()
	if err != nil{
		return err.Error()
	}
	if !stdout{
		return outb.String()
	} 
	return ""
}

// Search the given query in slice.
// Returns true if the element exists.
func searchInSlice(slice []string, query string) bool{
	set := make(map[string]bool)
	for _, v := range slice {
		set[v] = true
	}
	return set[query]
}

// Prepare (shorten) the git commands.
func prepareCmds(){
	// Show status if repository exists in directory.
	execCmd("git status", true)
	// Trimming the string using sed.
	// 's/^\s*//' -> Substitute the found expression (' ' with '').
	// 's/ *[A-Z].*//' -> Remove the git command description
	// && -> For the next command.
	removeSpaces := "sed -e 's/^\\s*//' -e 's/ *[A-Z].*//' && "
	// Parsing the git commands.
	// grep '^  *[a-z]' -> Select the lines starting with indent.
	// tr -d '*' -> Remove the '*' character.
	parseGitCmd := 
		"git help | grep '^  *[a-z]' | " + removeSpaces +
		"git branch | tr -d '*' | " + removeSpaces +
		"git remote"
	cmdStr := execCmd(parseGitCmd, false)
	cmdList = strings.Split(cmdStr, "\n")
	for _, cmd := range cmdList {
		if (len(cmd) > 0){
			// Use the first character of git command
			// for the new command if not exists in the
			// commands slice. (cmdSlice)
			// If first character is in the list, compose 
			// a two character abbreviation for it and
			// add it to slice.
			firstChar := string([]rune(cmd)[0])
			if (!searchInSlice(cmdSlice, firstChar)){
				cmdSlice = append(cmdSlice, firstChar)
			}else{
				cmdSlice = append(cmdSlice, firstChar + 
					string([]rune(cmd)[len(cmd)/2])) 
			}
		}
	}
}

// Create a git command from the given string.
// Returns changed/new command.
func buildCmd(line string) string {
	// Support the normal usage.
	line = strings.Replace(line, " git ", " ", -1)
	// Replace the shortened command with its original.
	for index, cmd := range cmdSlice {
		cmd = " " + cmd + " "
		if (strings.Contains(line, cmd)) {
			line = strings.Replace(line, cmd, 
				" " + cmdList[index] + " ", -1)
		}
	}
	return "git" + line
}

// Start the interactive shell.
func startTerm() {
	term, err := terminal.NewWithStdInOut()
	if err != nil {
		panic(err)
	}
	defer term.ReleaseFromStdInOut()
	whiteColor.Println("Type '?' for help or 'git' for list of commands.")
	term.SetPrompt("[god ~]$ ")
	cmdLoop:
	for {
		// Read the keyboard input.
		line, err := term.ReadLine()
		// Exit on Ctrl-D and Ctrl-C.
		if err == io.EOF ||  line == "^C" {
			fmt.Println()
			return
		}
		// Built-in commands.
		switch line{
		case "", " ": 
			break
		case "clear":
			execCmd("clear", true)
		case "exit":
			break cmdLoop
		case "?", "help":
			showHelp()
		case "version":
			showVersion()
		case "git":
			showCommands()
		default:
			// Build the git command.
			gitCmd := buildCmd(" " + line + " ")
			// Handle the execution of the input.
			if retval := execCmd(gitCmd, true); len(retval) > 0 {
				fmt.Fprintln(os.Stderr, retval)
			}
		}
	}
}

// Takes 'table' parameter and returns colored.
func setTableColors(table (*tablewriter.Table)) (*tablewriter.Table) {
	whiteTable := tablewriter.Colors{
		tablewriter.Bold, 
		tablewriter.FgHiWhiteColor}
	blackTable := tablewriter.Colors{
		tablewriter.Bold, 
		tablewriter.FgHiBlackColor}
	table.SetHeaderColor(whiteTable, whiteTable)
	table.SetColumnColor(whiteTable, blackTable)
	return table
}

// Display help message.
func showHelp(){
	cliCmds := map[string]string{
		"git": "List available simplified git commands",
		"help": "Show this help message",
		"version": "Show version information",
		"clear": "Clear the terminal", 
		"exit": "Exit shell"}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Command", "Description"})
	table = setTableColors(table)
	for k, v := range cliCmds {
		table.Append([]string{k, v})
	}
	table.Render()
}

// Show git commands in table.
func showCommands(){
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Command", "git"})
	table = setTableColors(table)
	for index, cmd := range cmdSlice {
		table.Append([]string{cmd, cmdList[index]})
	}
	table.Render()
}

// Show project information including version.
func showVersion(){
	fmt.Println()
	asciiFigure := figure.NewFigure("god", "cosmic", true)
	asciiFigure.Print()
	whiteColor.Println("\n ~ god:v" + version)
	fmt.Println(" ~ utility for simplifying the git usage" +
		"\n ~ github.com/keylo99/god\n")
}

func main() {
	// If -v argument is given, show project information and exit.
	// If not, start the god terminal.
	versionFlag := flag.Bool("v", false, "Show version information")
	flag.Parse()
	if(*versionFlag){
		showVersion()	
	}else{
		prepareCmds()
		startTerm()
	}
}
