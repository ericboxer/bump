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

// Commandline bools
var (
	shouldBumpMajor bool
	shouldBumpMinor bool
	shouldBumpPatch bool

	shouldDryRun bool
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
	flag.BoolVar(&shouldBumpPatch, "p", true, "Bumps the patch version")

	// flag.BoolVar(&shouldDryRun, "dry-run", false, "Shows what the outcome of running the bump command would be without making changes")
	flag.BoolVar(&shouldDryRun, "d", false, "Shows what the outcome of running the bump command would be without making changes")

	flag.Parse()

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
			panic(err)
		}

	}

	err = json.Unmarshal(dat, &currentVersion)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Current Version: %d.%d.%d\n", currentVersion.MAJOR, currentVersion.MINOR, currentVersion.PATCH)

	// Run these in reveres order. Extra work, but its so small who cares.

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

	gitTag := exec.Command("git", "tag", fmt.Sprintf("v%d.%d.%d", currentVersion.MAJOR, currentVersion.MINOR, currentVersion.PATCH))

	err = gitTag.Run()

	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
