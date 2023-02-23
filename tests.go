package main

import (
	"errors"
	"os"
	"testing"
)

func TestInputCommands(t *testing.T) {
	tests := []struct {
		tName         string
		inputGiven    string
		dirExpected   string // for testing directory changes
		errorExpected error
	}{
		{
			tName:         "Incorrect cd",
			inputGiven:    "cd",
			errorExpected: errors.New("path required"),
		},
		{
			tName:         "Correct cd",
			inputGiven:    "cd /",
			errorExpected: nil,
			dirExpected:   "/",
		},
		{
			tName:         "Correct dir",
			inputGiven:    "dir",
			errorExpected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.tName, func(t *testing.T) {
			if e := execInput(tt.inputGiven); e != tt.errorExpected {
				t.Errorf("execInput() error = %v, wantErr %v", e, tt.errorExpected)
			}

			// test for cd
			if tt.dirExpected != "" {
				curDir, e := os.Getwd()
				if e != nil {
					t.Errorf("Failed to get new directory: %v", e)
				}
				if tt.dirExpected != curDir {
					t.Errorf("Failed to change to desired directory.")
				}
			}
		})
	}
}
