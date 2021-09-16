Here is my submission for the Golang excercise. If running the .exe file, enter HELP for instructions.

Please note that this project was written, tested, and compiled on Windows 10.

Included in this folder:
lynxdb.go
 - The uncompiled source code.
 - This compiles to a binary executable that includes a rudimentary command line interface to access all functions.
lynxdb_test.go
 - The test code for test cases. *Testing can be executed from command line with the command 'go test lynxdb.go lynxdb_test.go'*!
proof_of_concept.csv
proof_2.csv
 - These two files are part of the test cases. They are supposed to be empty, and are also supposed to be empty upon the conclusion of testing..
 - I used .csv files as a kind of poor man's persistency layer in this implementation.
 - The .exe can write new csv files in the working directory via user input, but for the purposes of testing I assumed the files already existed, as the creating of these files was not a part of the interface parameters.