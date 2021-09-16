Please note that this project was written, tested, and compiled on Windows 10.

Included in this folder:
lynxdb.go
 - The uncompiled source code.
 - This compiles to a binary executable that includes a rudimentary command line interface to access all functions.
lynxdb_test.go
 - The test code for test cases. *Testing can be executed from command line with the command 'go test lynxdb.go lynxdb_test.go'*!
 
To run this properly, you're going to want to run it from a folder with 1+ empty csv files. Not exactly the world's most sophisticated persistency layer, but that's not the point.
