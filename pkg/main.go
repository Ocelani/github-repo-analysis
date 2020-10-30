package main

import (
	"fmt"

	"github.com/Ocelani/github-repos-measure/pkg/analysis"
)

func main() {
	fmt.Println("Running...")
	java, python := analysis.ReadLanguagesCsv()
	fmt.Printf(`
    +------------+------------+------------+
    |   COUNT    |    JAVA    |   PYTHON   |
    +------------+------------+------------+
    |   STARS    |   %v   |   %v   |
    |  WATCHERS  |   %v    |   %v    |
    |   FORKS    |   %v   |   %v   |
    |  RELEASES  |    %v     |    %v    |
    |    AGE     |    %v     |    %v     |
    |   LINES    |  %v  |  %v  |
    |   CODE     |  %v  |  %v  |
    |  COMMENTS  |  %v   |   %v   |
    |   BLANKS   |  %v   |   %v   |
    | COMPLEXITY |  %v   |   %v   |
    +------------+------------+------------+
  `, java.StargazerCount, python.StargazerCount,
		java.Watchers, python.Watchers,
		java.ForkCount, python.ForkCount,
		java.Releases, python.Releases,
		java.Age, python.Age,
		java.Lines, python.Lines,
		java.Code, python.Code,
		java.Comments, python.Comments,
		java.Blanks, python.Blanks,
		java.Complexity, python.Complexity,
	)
	fmt.Printf(`
    +------------+------------+------------+
    |    MEAN    |    JAVA    |   PYTHON   |
    +------------+------------+------------+
    |   STARS    |    %v    |    %v    |
    |  WATCHERS  |    %v     |    %v     |
    |   FORKS    |    %v    |    %v    |
    |  RELEASES  |     %v      |     %v     |
    |    AGE     |     %v      |     %v      |
    |   LINES    |   %v   |   %v   |
    |   CODE     |   %v   |   %v   |
    |  COMMENTS  |   %v    |    %v    |
    |   BLANKS   |   %v    |    %v    |
    | COMPLEXITY |   %v    |    %v    |
    +------------+------------+------------+
  `, java.StargazerCountMean, python.StargazerCountMean,
		java.WatchersMean, python.WatchersMean,
		java.ForkCountMean, python.ForkCountMean,
		java.ReleasesMean, python.ReleasesMean,
		java.AgeMean, python.AgeMean,
		java.LinesMean, python.LinesMean,
		java.CodeMean, python.CodeMean,
		java.CommentsMean, python.CommentsMean,
		java.BlanksMean, python.BlanksMean,
		java.ComplexityMean, python.ComplexityMean,
	)
}
