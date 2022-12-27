package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Version struct {
	MAJOR int
	MINOR int
	PATCH int
}

var (
	currentVersion = Version{
		MAJOR: 0,
		MINOR: 0,
		PATCH: 0,
	}
)

var (
	BUMPVER string
)

// Commandline bools
var (
	shouldBumpMajor bool
	shouldBumpMinor bool
	shouldBumpPatch bool

	shouldDryRun bool
	shouldGitTag bool

	shouldMake bool
)

func bumpMajor() {
	currentVersion.MAJOR += 1
	currentVersion.MINOR = 0
	currentVersion.PATCH = 0
}

func bumpMinor() {
	currentVersion.MINOR += 1
	currentVersion.PATCH = 0
}

func bumpPatch() {
	currentVersion.PATCH += 1
}

func writeVersion() {

	versionJSON, err := json.Marshal(currentVersion)

	if err != nil {
		panic(err)
	}

	os.WriteFile("VERSION", []byte(versionJSON), 0666)
}

func main() {

	// Setup our flags
	flag.BoolVar(&shouldBumpMajor, "M", false, "Bumps the major version")
	flag.BoolVar(&shouldBumpMinor, "m", false, "Bumps the minor version")
	flag.BoolVar(&shouldBumpPatch, "p", false, "Bumps the patch version")

	// flag.BoolVar(&shouldDryRun, "dry-run", false, "Shows what the outcome of running the bump command would be without making changes")
	flag.BoolVar(&shouldDryRun, "d", false, "Shows what the outcome of running the bump command would be without making changes")

	flag.BoolVar(&shouldGitTag, "t", false, "Adds a tag with the current version number to the current commit")
	flag.BoolVar(&shouldMake, "make", false, "Runs Make with the variables MAJOR, MINOR, and PATCH")

	flag.Parse()

	fmt.Println("*******************")
	fmt.Printf("Bump Version %v\n", BUMPVER)
	fmt.Println("*******************")

	if shouldDryRun {
		fmt.Println("*******************")
		fmt.Println("*THIS IS A DRY RUN*")
		fmt.Println("*******************")
	}

	dat, err := os.ReadFile("./VERSION")

	if err != nil {

		if strings.Contains(err.Error(), "no such file or directory") {
			fmt.Println("No VERSION file found. Creating...")
			writeVersion()

		} else {
			panic(err.Error())
		}

		fmt.Println("VERSION file found...")

	}

	err = json.Unmarshal(dat, &currentVersion)

	if err != nil {

		if strings.Contains(err.Error(), "unexpected end of JSON input") {
			fmt.Println("Unexpected end of JSON")
		} else {
			panic(err.Error())

		}
	}

	fmt.Printf("Current Version: %d.%d.%d\n", currentVersion.MAJOR, currentVersion.MINOR, currentVersion.PATCH)

	// Run these in reveres order. It's a bad way of handling mutual exclusivity.
	if shouldBumpPatch {
		bumpPatch()
	}

	if shouldBumpMinor {
		bumpMinor()
	}

	if shouldBumpMajor {
		bumpMajor()
	}

	fmt.Printf("New Version: %d.%d.%d\n", currentVersion.MAJOR, currentVersion.MINOR, currentVersion.PATCH)

	if !shouldDryRun {
		writeVersion()
	}

	if shouldGitTag {

		gitTagVal := fmt.Sprintf("v%d.%d.%d", currentVersion.MAJOR, currentVersion.MINOR, currentVersion.PATCH)
		gitTag := exec.Command("git", "tag", gitTagVal)
		err = gitTag.Run()

		fmt.Printf("Git tag %s created", gitTagVal)

		if err != nil {
			panic(err)
		}
	}

	if shouldMake {

		// Make sure the makefile exists

		_, err := os.Stat("makefile")

		if err != nil {

			fmt.Println("makefile was not found...")

		} else {

			makeCmdVer := fmt.Sprintf("SEMVER=%d.%d.%d", currentVersion.MAJOR, currentVersion.MINOR, currentVersion.PATCH)
			makeCmd := exec.Command("make", "buildall", makeCmdVer)

			// Set the makefile output to standard out se we cal see whats going on
			makeCmd.Stdout = os.Stdout
			err = makeCmd.Run()

			if err != nil {
				panic(err.Error())
			}
		}

	}

	os.Exit(0)
}
