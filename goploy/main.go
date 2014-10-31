package main

import (
	"bytes"
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/edsrzf/go-git"
	"github.com/ingenieux/goploy"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func oopse(e error) {
	if nil != e {
		panic(e)
	}
}

func oops(format string) {
	fmt.Fprintf(os.Stderr, format+"\n")

	os.Exit(1)
}

func main() {
	showUsage := func() {
		fmt.Println(usage)
		os.Exit(1)
	}

	if 1 == len(os.Args) {
		showUsage()
	}

	params, err := docopt.Parse(usage, nil, false, "", false)

	if nil != err {
		fmt.Fprintf(os.Stderr, "Error: %q\n\n%s\n", err, usage)

		os.Exit(1)
	} else if params["--help"].(bool) {
		showUsage()
	}
	oopse(err)

	environment := ""
	commitId := ""

	if vCommit, ok := params["--commitId"].(string); ok {
		commitId = vCommit
	}

	if vEnvironment, ok := params["ENVIRONMENT"].(string); ok {
		environment = vEnvironment
	}

	Context{
		directory:   params["--directory"].(string),
		branch:      params["--branch"].(string),
		commitId:    commitId,
		application: params["APPLICATION"].(string),
		environment: environment,
		genurl:      params["genurl"].(bool),
		push:        params["push"].(bool),
	}.runWith()
}

type Context struct {
	directory   string
	branch      string
	commitId    string
	application string
	environment string
	genurl      bool
	push        bool
}

func (c Context) runWith() {
	gitPath := filepath.Join(c.directory, ".git")

	{ // Validates git directory
		d, err := os.Stat(gitPath)

		oopse(err)

		if !d.IsDir() {
			oops("Not a directory: " + gitPath)
		}
	}

	parentDir := path.Dir(gitPath)

	commitId := c.commitId

	branch := c.branch

	// Must lookup commit Id
	if "" == c.commitId {
		r := git.NewRepo(gitPath)

		refs := r.Refs()

		if ref, ok := refs[branch]; ok {
			commitId = ref.String()
		} else {
			branchList := make([]string, 0)

			for k, _ := range refs {
				branchList = append(branchList, k)
			}

			oops(fmt.Sprintf("Unable to find branch '%s'. Available branches are: %s", branch, branchList))
		}
	}

	deploy, err := goploy.NewDeployment()

	oopse(err)

	err = deploy.CommitId(commitId)

	oopse(err)

	deploy.ApplicationName(c.application)

	deploy.EnvironmentName(c.environment)

	if c.genurl {
		fmt.Println(deploy.GetPushURL())
	} else if c.push {
		cmd := exec.Command("git", "push", "--force", deploy.GetPushURL(), branch)

		cmd.Dir = parentDir

		fmt.Printf("command: %s\n", cmd.Args)

		out := bytes.Buffer{}
		cmd.Stdout = &out

		err = cmd.Run()

		fmt.Printf("Output:\n%s\n", out.String())

		oopse(err)
	}
}
