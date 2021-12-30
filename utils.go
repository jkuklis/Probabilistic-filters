package my_main

func Check(e error) {
    if e != nil {
        panic(e)
    }
}