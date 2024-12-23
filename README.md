# beetlejuice.go

Simple Go library to hide your secrects in a YAML configuration file.

## How to use it

    $ cat data/.bj
    foo: bar
    some-key: "Quoted!"

    
    package main
    
    import (
        "github.com/mackode/beetlejuice.go"
    )

    func main() {
        bj := beetlejuice.NewBeetleJuice()

        secret, err := bj.Loopup("secret-name")
        if err != nil {
            panic(err)
        }
    }

