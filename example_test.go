package gitcmd

import (
	"fmt"
	"log"
)

func Example() {
	git := New(nil)

	out, err := git.Exec("config", "--get", "remote.origin.url")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
	// Output:
	// git@github.com:tsuyoshiwada/go-gitcmd.git
}
