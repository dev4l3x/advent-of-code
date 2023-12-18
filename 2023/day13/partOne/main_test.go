package main

import (
	"testing"
)

func TestSummarizingNotesIsCorrect(t *testing.T) {

	notes := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

	const expectedSummary = 405

	summary := SummarizeNotes(notes)

	if expectedSummary != summary {
		t.Fatalf("Expected %v but got %v", expectedSummary, summary)	
	}
}
