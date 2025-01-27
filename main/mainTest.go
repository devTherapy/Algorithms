// package main

// import "fmt"

// type ServerState int

// const (
//     StateIdle ServerState = iota
//     StateConnected
//     StateError
//     StateRetrying
// )

// var stateName = map[ServerState]string{
//     StateIdle:      "idle",
//     StateConnected: "connected",
//     StateError:     "error",
//     StateRetrying:  "retrying",
// }

// func (ss ServerState) String() string {
//     return stateName[ss]
// }

// type MyIntSlice []int

// func main() {
//     slice := MyIntSlice{10, 20, 30, 40}
//     value := 30
//     index := SlicesIndex(slice, value)
//     fmt.Println(index)
// }

// func transition(s ServerState) ServerState {
//     switch s {
//     case StateIdle:
//         return StateConnected
//     case StateConnected, StateRetrying:

//         return StateIdle
//     case StateError:
//         return StateError
//     default:
//         panic(fmt.Errorf("unknown state: %s", s))
//     }

	
// }

// func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
//     for i := range s {
//         if v == s[i] {
//             return i
//         }
//     }
//     return -1
// }