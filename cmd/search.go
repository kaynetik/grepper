package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a term in given file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Get values of flag arguments
		filename, _ := cmd.Flags().GetString("file")
		sTerm, _ := cmd.Flags().GetString("sterm")

		res, _ := searchFile(filename, sTerm)
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringP("file", "f", "", "Filename | Path to a file")
	searchCmd.Flags().StringP("sterm", "s", "", "Search Term")
}

func searchFile(path, sTerm string) (string, error) {
	scanner, err := openFile(path)
	if err != nil {
		return "", err
	}

	line := 1
	var lines int
	res := make([]string, lines)
	for scanner.Scan() {
		// if the search term is found on the current line, append it to the resulting slice
		if strings.Contains(scanner.Text(), sTerm) {
			res = append(res, scanner.Text())
		}

		line++
	}

	if err := scanner.Err(); err != nil {
		return "", errors.New("an unexpected error occurred")
	}

	if len(res) < 1 {
		return "", errors.New("nothing found by that search term")
	}

	return buildStrFromSlice(res), nil
}

func openFile(path string) (*bufio.Scanner, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(f), nil
}

// Build response as a single string from a slice of strings
func buildStrFromSlice(ss []string) string {
	var sb strings.Builder
	for _, str := range ss {
		sb.WriteString(str)
		sb.WriteString("\n")
	}
	return sb.String()
}

func getPath(filename string) string {
	// TODO: Sanitize
	// check if its a path/filename, and if that given file exists
	return ""
}
