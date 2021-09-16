package main

import (
	"testing"
	"fmt"
)

func TestTable(t *testing.T) {
	var table = []struct {
		db_name string
		input_key []byte
		input_val []byte
		expected_val []byte
		expected_err error
		expected_stats string
	}{
		{
			db_name: "proof_of_concept.csv",
			input_key: []byte("blah"),
			input_val: []byte("16"),
			expected_val: []byte("16"),
			expected_err: nil,
			expected_stats: "1",
		},
		{
			db_name: "proof_of_concept.csv",
			input_key: []byte("clock"),
			input_val: []byte("90"),
			expected_val: []byte("90"),
			expected_err: nil,
			expected_stats: "1",
		},
		{
			db_name: "proof_2.csv",
			input_key: []byte("tick"),
			input_val: []byte("8"),
			expected_val: []byte("8"),
			expected_err: nil,
			expected_stats: "1",
		},
		{
			db_name: "proof_2.csv",
			input_key: []byte("mort"),
			input_val: []byte("16"),
			expected_val: []byte("16"),
			expected_err: nil,
			expected_stats: "1",
		},
	}

	for _, out := range table {
		t.Run(out.db_name, func(t *testing.T) {
			tester := Datastore{name: out.db_name}
			tester.DB_Init()
			// check if open works
			given_err := tester.DB_Open()
			if given_err != out.expected_err {
				fmt.Println(given_err)
				t.Errorf("The file does not exist")
			}
			// check if put works
			given_err = tester.DB_Put(out.input_key, out.input_val)
			if given_err != out.expected_err {
				fmt.Println(given_err)
				t.Errorf("The data cannot be input")
			}
			// check if get works
			given_val, _ := tester.DB_Get(out.input_key)
			if string(given_val) != string(out.expected_val) {
				t.Errorf("The given key does not have the given associated value")
			}
			// check if stats works
			given_stats := tester.DB_Stats()
			if given_stats != out.expected_stats {
				t.Errorf("The given key does not have the expected stats")
			}
			// check if delete works
			given_err = tester.DB_Delete(out.input_key)
			if given_err != out.expected_err {
				fmt.Println(given_err)
				t.Errorf("The given key does not exist and cannot be deleted")
			}
			// check if delete reports the proper error when key does not exist
			given_err = tester.DB_Delete(out.input_key)
			if given_err == out.expected_err {
				fmt.Println(given_err)
				t.Errorf("The delete function was able to delete the same key twice in a row")
			}
			// check if flush works
			given_err = tester.DB_Flush()
			if given_err != out.expected_err {
				fmt.Println(given_err)
				t.Errorf("The flush function did not complete properly")
			}

			tester.DB_Close()
		}) 
	}
} 