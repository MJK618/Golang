package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "math"
)

func GenDisplaceFn(a, vo, so float64) func (float64) float64 {

    fn := func (t float64) float64 {
        return 0.5 * a * math.Pow(t, 2) + (vo * t) + so
    }
    return fn
}

func read_fp(var_type string) float64{

    var temp string
    var valid int
    var fp float64

    read_input := bufio.NewScanner(os.Stdin)

    temp = var_type
    valid = 0
    for valid == 0 {
        fmt.Printf("\nPlease enter %s as a floating point: ", var_type)
        read_input.Scan()
        temp = read_input.Text()
        fp, err := strconv.ParseFloat(temp, 64)
        if err == nil{
            valid = 1
            return fp
        } else {
            fmt.Println("Invalid input. Floating point number required.")
        }
    }
    return fp
}

func main(){

    var a float64
    var vo float64
    var so float64
    var t float64

    a = read_fp("accelleration")
    vo = read_fp("initial velocity")
    so = read_fp("initial displacement")
    t = read_fp("time")

    fn := GenDisplaceFn(a, vo, so)
    fmt.Println("\nThe displacement is", fn(t))

}