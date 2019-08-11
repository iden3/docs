package main

import (
	"fmt"
	"time"

	git "gopkg.in/src-d/go-git.v4"
)

const PATH = "./"

func checkErr(step string, err error) {
	if err != nil {
		currTime := time.Now().Format("2006.01.02 15:04:05")
		fmt.Println(currTime+" [ERROR in step '"+step+"']:", err)
	}
}

func main() {
	r, err := git.PlainOpen(PATH)
	checkErr("open git", err)

	w, err := r.Worktree()
	checkErr("worktree", err)

	for {
		err = w.Pull(&git.PullOptions{RemoteName: "origin"})
		checkErr("pull", err)

		if err == nil {
			ref, err := r.Head()
			checkErr("head", err)

			commit, err := r.CommitObject(ref.Hash())
			checkErr("get commitObject", err)

			currTime := time.Now().Format("2006.01.02 15:04:05")
			fmt.Println(currTime+" Pulled new commit: ", commit)
		}
		time.Sleep(60 * 1000 * time.Millisecond)
	}
}
