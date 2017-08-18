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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const URL = "https://gopher.rocks"
const SearchPath = "/packages/search"

var (
	AuthorSearch  string
	EmailSearch   string
	KeywordSearch string
)

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Version struct {
	Tag        string    `json:"tag"`
	ReleasedOn time.Time `json:"released_on"`
}
type Package struct {
	Authors     []Author  `json:"authors"`
	ImportPath  string    `json:"import_path"`
	SourceURL   string    `json:"source_url"`
	License     string    `json:"license"`
	Keywords    []string  `json:"keywords"`
	Description string    `json:"description"`
	Versions    []Version `json:"versions"`
}
type PackageSearchQuery struct {
	Authors  []string `json:"authors"`
	Keywords []string `json:"keywords"`
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for packages",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		psq := PackageSearchQuery{
			Authors:  []string{AuthorSearch},
			Keywords: []string{KeywordSearch},
		}
		err := doSearch(psq)
		if err != nil {
			fmt.Println("Unable to search:", err)
		}
	},
}

func doSearch(p PackageSearchQuery) error {
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}
	url := URL + SearchPath
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var pkgs []Package
	err = json.Unmarshal(body, &pkgs)
	if err != nil {
		return errors.Wrap(err, "Unmarshaling response.")
	}
	for x, pac := range pkgs {
		fmt.Printf("%d. Name: %s \n", x+1, pac.ImportPath)
	}
	return nil
}

func init() {
	RootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	searchCmd.PersistentFlags().StringVarP(&AuthorSearch, "author", "a", "", "author name search term")
	searchCmd.PersistentFlags().StringVarP(&KeywordSearch, "keyword", "k", "", "keyword search term")

}
