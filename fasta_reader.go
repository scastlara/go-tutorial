package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "log"
)

func print_gc(sequence string, identifier string)()  {

    var gccount float32 = 0
    var length  float32 = 0
    for i := 0; i < len(sequence); i++ {
        if string(sequence[i]) == "G" || string(sequence[i]) == "C" {
            gccount++
        }
        if string(sequence[i]) != "\n" {
            length++
        }
    }
    /*
    Remember: integer division if 0.XX is always 0.
    that's why gccount and length must be floats!
    */
    gcperc := gccount / length
    fmt.Printf("%s\t%.3f\n", identifier, gcperc)

    return
}

func main()  {
    filename := os.Args[1]
    fh, err := os.Open(filename)
    if err != nil {
        panic("Can't read FASTA file.")
    }

    bf := bufio.NewReader(fh)
    var identifier string = ""
    var sequence   string = ""
    for {
        line, isPrefix, err := bf.ReadLine()

        if err == io.EOF {
            break
        }

        if isPrefix {
           log.Fatal("Error: Unexpected long line reading", fh.Name())
       }
       if string(line[0]) == ">" {
           if len(sequence) > 0 {
               print_gc(sequence, identifier)
               sequence = ""
            }
            identifier = string(line)[1:]
        } else {
            sequence += string(line)
       }
    }
    print_gc(sequence, identifier)
}
