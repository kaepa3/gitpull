package main

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	fmt.Println("start")

	r, err := git.PlainOpen("./")
	CheckIfError(err)

	// Get the working directory for the repository
	w, err := r.Worktree()
	CheckIfError(err)

	// Pull the latest changes from the origin remote and merge into the current branch
	fmt.Println("git pull origin")
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	CheckIfError(err)
}

func CheckIfError(err error) {
	if err != nil {
		panic(err)
	}
}
