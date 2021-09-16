package main

import (
	"fmt"
	"io"
	"os"
    "sync"
    "errors"
    "encoding/csv"
    "path/filepath"
    "strconv"
)
/*
 testdb interface
 - assume a key/value DB
*/

type BackendDatabase interface {
	//DB_New() error // make a new database *I chose to use a 
	DB_Open() error // opens existing database
	DB_Put([]byte, []byte) error // put things in database 
	DB_Get([]byte) ([]byte, error) // get things in database 
	DB_Delete([]byte) (error) // delete things from database 
	DB_Close() // close database and saves changes
	DB_Stats() string // list traits of database
	DB_Flush() error // flush all things from database
}

// declare datastore medium
type Datastore struct {
    lock *sync.RWMutex
    name string
    datas map[string]string
}

// This constructor was necessary to make the Datastore struct  
func (s *Datastore) DB_Init() {
    s.datas = make(map[string]string)
    s.lock = new(sync.RWMutex)
}

// make new database file in working directory. Note this is not part of the interface, and is instead a function intended solely for cmd line use.
func New(filename string) error {
	//create file
	var err error
    var _, check = os.Stat(filename)
    if os.IsNotExist(check) {
        file, err := os.Create(filename)
        if err != nil {
            fmt.Println(err)
            return err
        }
        defer file.Close()
    } else {
        fmt.Println(err, filename)
        return err
    }
    return err
	
}

// open database from working directory and make data available within datastore struct
func (s *Datastore) DB_Open() error {
    // open file
    csvfile, err := os.Open(s.name)
    if err != nil {
        fmt.Println(err)
        return err
    }
    defer csvfile.Close()

    // read thre file
    r := csv.NewReader(csvfile)
    for {   
        line, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println(err)
        }
        holder := line[0]
        // if key (holder) is not in the database, enter key/value pair
        if _, check := s.datas[holder]; check == false {
            s.datas[holder] = line[1]
        }
    }
    return err
}


// put things in the slices of the currently open database
func (s *Datastore) DB_Put(key []byte, val []byte) error {
	key_string := string(key)
    val_string := string(val)
    // check if key exists
    _, check := s.datas[key_string]
    if !check {
        s.datas[key_string] = val_string
    } else {
        return errors.New("Key already exists")
    }
    return nil
}


//get things from discrete location within database
func (s *Datastore) DB_Get(key []byte) ([]byte, error) {
    key_string := string(key)
    val, check := s.datas[key_string]
    var val_slice []byte
    if check {
        val_slice = []byte(val)
    } else {
        return nil, errors.New("No key or associated value")
    }
    return val_slice, nil
}

// delete things from discrete location within database
func (s *Datastore) DB_Delete(key []byte) error {
    key_string := string(key)
    _, check := s.datas[key_string]
    if check {
        delete(s.datas, key_string)
    } else {
        return errors.New("No key or associated value to delete")
    }
    return nil
}

// close database and save all changes to file in directory
func (s *Datastore) DB_Close() {
    filename := s.name
    destination, err := os.Create(filename)
    if err != nil {
        fmt.Println("os.Create:", err)
        return
    }
    defer destination.Close()

    // write new data to database csv
    for key, val := range s.datas {
        fmt.Fprintf(destination, "%v,", string(key))
        fmt.Fprintf(destination, "%v\n", string(val))
    }
}

// list length of database
func (s *Datastore) DB_Stats() string {
    output_int := len(s.datas)
    output := strconv.Itoa(output_int)
    return output
}

// clear all things from open database
func (s *Datastore) DB_Flush() error {
    s.DB_Init()
    return nil
}

//Find string in slice
func Find(body []string, entry string) bool {
    //iterate thru slice until entry is found
    for _, value := range body {
        if value == entry {
            return true
        }
    }
    return false
}

func main() {
    // Get all csv files in working directory at start. 
    var files []string
    wdir, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }
    err = filepath.Walk(wdir, func(path string, info os.FileInfo, err error) error {
        fileExt := filepath.Ext(path)
        if fileExt == ".csv" {
            fmt.Println(path)
            files = append(files, info.Name())
        }
        return nil
    })
    // if for some reason you cannot access wd
    if err != nil {
        panic(err)
    }
    // This assembled list of files are the available databases.

    //get user input from command line
    var input string
    var key string
    var value string
    halter := 0;
    halter2 := 0;
    var db_name string
    // let user enter commands until end condition (halter) is met
    for halter < 1 {
        fmt.Println("Enter DB manipulation command. Type HELP for available commands:")
        fmt.Scanln(&input)
        //check input against commands
        switch input {
        case "NEW":
            fmt.Println("Enter the name of your new database. *Please include a .csv extension*")
            fmt.Scanln(&input)
            err := New(input)
            fmt.Println(err)
        case "OPEN":    
            fmt.Println("Enter the name of the database you wish to open. \n *Please include the .csv extenstion in the name*")
            fmt.Scanln(&db_name)
            // make sure files exists before attempting to open
            checker := Find(files, db_name)
            // proceed
            if checker {
                operating_db := Datastore{name: db_name}
                operating_db.DB_Init()
                operating_db.DB_Open()
                halter2 = 0;
                // let user enter second layer of commands until end condition (halter) is met
                for halter2 < 1 {
                    fmt.Println("Enter DB manipulation command for open DB. Type HELP for available commands:")
                    fmt.Scanln(&input)
                    //check input against commands
                    switch input {
                    case "CLOSE":
                        operating_db.DB_Close()
                        halter2 = 1
                    case "PUT": 
                        fmt.Println("Enter key string:")
                        fmt.Scanln(&key)
                        fmt.Println("Enter value value:")
                        fmt.Scanln(&value)
                        operating_db.DB_Put([]byte(key),[]byte(value))
                    case "GET":
                        fmt.Println("Enter key string")
                        fmt.Scanln(&key)
                        output, _ := operating_db.DB_Get([]byte(key))
                        fmt.Printf("The value assoicated with key '%v' is ", key)
                        fmt.Printf("'%v\n'", string(output))
                    case "DELETE":
                        fmt.Println("Enter key string")
                        fmt.Scanln(&key)
                        operating_db.DB_Delete([]byte(key))
                    case "STATS":
                        results := operating_db.DB_Stats()
                        fmt.Println(results)
                    case "FLUSH":
                        operating_db.DB_Flush()
                    case "EXIT":
                        fmt.Println("Exiting before closing the database will result in all changes being lost. Enter Y to confirm. Press any other key to cancel.")
                        var answer string
                        fmt.Scanln(&answer)
                        switch answer {
                        case "Y":
                            halter2 = 1
                            halter = 1
                        default:
                            fmt.Println("cancelled")
                        }
                    case "HELP":
                        fmt.Println("The available database functions are CLOSE, PUT, GET, DELETE, STATS, AND FLUSH. Enter EXIT to close the program.")
                    default: 
                        fmt.Println("Command not recognized.")
                    }
                }
            } else {
                fmt.Println("The given database does not exist.")
            }
        // back to first-layer options
        case "LIST":
            fmt.Println("Here are the available DB source files:")
            for _, entry := range files {
                fmt.Println(entry)
            }
        case "HELP":
            fmt.Println("The available database functions are NEW, OPEN, and LIST. Enter EXIT to close the program.")
        case "EXIT":
            // breaks the for loop
            halter = 1
        default: 
            fmt.Println("Command not recognized.")
        }
    }
}