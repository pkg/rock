// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/blang/semver"
	"github.com/gogits/git"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var Major bool
var Minor bool
var Patch bool

// tagCmd represents the tag command
var tagCmd = &cobra.Command{
	Use:     "tags",
	Short:   "View or add semantic version tags to a project",
	Long:    `The tag command allows you to view or create tags in your version control system.  Tags are based on Semantic Version numbering, and are comprised of Major, Minor, and Patch version identifiers.`,
	Aliases: []string{"tag"},
	Run: func(cmd *cobra.Command, args []string) {
		var maxMajor, maxMinor, maxPatch uint64

		repo, err := git.OpenRepository("./.git")
		if err != nil {
			fmt.Println("The current directory doesn't appear to be a git repository.")
			return
		}
		tags, err := repo.GetTags()
		if err != nil {
			fmt.Println("Error Getting Tags: ", err)
			return
		}
		fmt.Println("Existing Tags:")
		for _, tag := range tags {
			if string(tag[0]) == "v" {
				tag = tag[1:len(tag)]
			}
			sv, err := semver.Make(tag)
			if err != nil {
				fmt.Println("error: ", err)
				continue
			}
			if sv.Major > maxMajor {
				maxMajor = sv.Major
				maxMinor = 0
				maxPatch = 0
			}

			if sv.Minor > maxMinor {
				maxMinor = sv.Minor
				maxPatch = 0
			}

			if sv.Patch > maxPatch {
				maxPatch = sv.Patch
			}
			fmt.Printf("v%d.%d.%d\n", sv.Major, sv.Minor, sv.Patch)
		}
		var tagged bool
		if Major {
			maxMajor = maxMajor + 1
			maxMinor = 0
			maxPatch = 0
			if err := tag(repo, maxMajor, maxMinor, maxPatch); err != nil {
				fmt.Println(err)
				return
			}
			tagged = true
		}

		if Minor {
			maxMinor = maxMinor + 1
			maxPatch = 0
			if err := tag(repo, maxMajor, maxMinor, maxPatch); err != nil {
				fmt.Println(err)
				return
			}
			tagged = true
		}

		if Patch {
			maxPatch = maxPatch + 1
			if err := tag(repo, maxMajor, maxMinor, maxPatch); err != nil {
				fmt.Println(err)
				return
			}
			tagged = true
		}
		if tagged {
			fmt.Printf("Tagged and pushed release %s\n", tagString(maxMajor, maxMinor, maxPatch))
		}
	},
}

func init() {
	RootCmd.AddCommand(tagCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagCmd.PersistentFlags().String("foo", "", "A help for foo")
	tagCmd.PersistentFlags().BoolVarP(&Major, "major", "m", false, "increment major version")
	tagCmd.PersistentFlags().BoolVarP(&Minor, "minor", "n", false, "increment minor version")
	tagCmd.PersistentFlags().BoolVarP(&Patch, "patch", "p", false, "increment patch version")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}

func tag(repo *git.Repository, major, minor, patch uint64) error {

	fmt.Printf("Creating v%d.%d.%d\n", major, minor, patch)
	id, err := repo.GetCommitIdOfBranch("master")
	if err != nil {
		return errors.Wrap(err, "Unable to get commit ID of branch ")
	}
	fmt.Println("Current Commit ID: ", id)
	tag := tagString(major, minor, patch)
	err = repo.CreateTag(tag, id)
	if err != nil {
		return errors.Wrap(err, "Unable to create tag")
	}
	cmd := exec.Command("git", "push", "origin", tag)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, "Unable to push to origin")
	}
	return errors.Wrap(err, "Unable to create tag")
}

func tagString(major, minor, patch uint64) string {

	return fmt.Sprintf("v%d.%d.%d", major, minor, patch)
}
