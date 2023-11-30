package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
)

var (
	printVersion                  bool
	initializeConventionalCommits bool
	removeConventionalCommits     bool
	version                       string
)

func init() {
	flag.BoolVar(&printVersion, "v", false, "Prints version number")
	flag.BoolVar(&initializeConventionalCommits, "i", false, "Sets up git with conventional commits.")
	flag.BoolVar(&removeConventionalCommits, "r", false, "Removes conventional commit validation")
}

func main() {
	flag.Parse()

	run()
}

func run() {

	if printVersion {
		fmt.Printf("version: %s\n", version)
		return
	}

	if initializeConventionalCommits {
		fmt.Println("Added conventional commits to Git repository")
		err := setupGitHook()
		if err != nil {
			fmt.Println("Failed to add conventional commits to Git repository")
		}
		return
	}

	if removeConventionalCommits {
		fmt.Println("Removed conventional commits from Git repository")
		err := removeGitHook()
		if err != nil {
			fmt.Println("Failed to remove conventional commits from Git repository")
		}
		return
	}

	verifyCommitMessage()

}

func setupGitHook() error {

	// check if .git folder exists, to verify that it's a git repo
	if !checkFileExists(".git") {
		return errors.New("can not find .git folder")
	}

	// check if pre-commit file exists in .git/hooks/ folder
	if checkFileExists(".git/hooks/commit-msg") {
		return errors.New("pre-commit file exists")
	}

	var commitContent string

	commitContent = "#!/bin/bash"

	// creates a commit-msg file and adds content
	// sets correct access rights to files in commit-msg
	commitContent = commitContent + fmt.Sprintf("\nkoche \"$1\"")

	s := []byte(commitContent)
	err := os.WriteFile(".git/hooks/commit-msg", s, 0777)

	if err != nil {
		return err
	}

	return nil
}

func removeGitHook() error {

	// check if .git folder exists, to verify that it's a git repo
	if !checkFileExists(".git") {
		return errors.New("can not find .git folder")
	}

	// check if pre-commit file exists in .git/hooks/ folder
	if checkFileExists(".git/hooks/commit-msg") {
		return errors.New("pre-commit file exists")
	}

	err := os.Remove(".git/hooks/commit-msg")

	if err != nil {
		return err
	}

	return nil
}

func verifyCommitMessage() {

	// Verifies that there is only one parameter
	if len(os.Args) != 2 {
		fmt.Println("This command must be run from within a Git commit hook directory\nTo setup Koche, run 'koche -i' in a Git repository")
		return
	}

	// Opening the commit_editmsg file
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Could not open ./git/COMMIT_EDITMSG file:", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		errFile := file.Close()
		if errFile != nil {
			fmt.Println("Could not close ./git/COMMIT_EDITMSG file:", err)
			fmt.Println("If this problem continues, remove the commit-msg hook and run koche -i again.")
			os.Exit(1)
		}
	}(file)

	// reads the commit message from file
	scanner := bufio.NewScanner(file)
	var commitMessage string
	for scanner.Scan() {
		commitMessage += scanner.Text() + "\n"
	}

	// defines a REGEX for conventional commits format
	conventionalCommitRegex := regexp.MustCompile(`^(feat|fix|docs|style|refactor|perf|test|chore)(\(.+\))?!?: .+`)

	// verifies that the commit message follows conventional commits format
	if !conventionalCommitRegex.MatchString(commitMessage) {
		fmt.Println("\nInvalid commit-message. The commit message needs to follow this format:\n\n    <type>[optional scope]: <description>\n\n    [optional body]\n\n    [optional footer(s)]\n\n")
		fmt.Println("Valid types: feat, fix, docs, style, refactor, perf, test, chore\n")
		os.Exit(1)
	}
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	// check if error is "file not exists"
	if os.IsNotExist(err) {
		fmt.Printf("%v file does not exist\n", filePath)
		return false
	} else {
		fmt.Printf("%v file exist\n", filePath)
		return true
	}
}
