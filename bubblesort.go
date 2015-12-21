package main

import (
    "encoding/json"
    "log"
    "net/http"
    "io/ioutil"
    "sort"
    "time"
)

type request_struct struct {
    List []int `json:"list"`
}

type resp_struct struct {
    Original []int `json:"original"`
    bubblesort_struct
}

type bubblesort_struct struct {
    Sorted []int `json:"sorted"`
    Reps int `json:"reps"`
    Comparisons int `json:"comparisons"`
    Swaps int `json:"swaps"`
    Duration string `json:"duration"`
}


func Bubblesort(i []int) (o bubblesort_struct) {
    repetitionCount := 0
    comparisonsCount := 0
    swappedCount := 0
    t0 := time.Now()

    slice := sort.IntSlice(i)
    for itemCount := slice.Len() - 1; ; itemCount-- {
        repetitionCount++
        hasChanged := false
        for index := 0; index < itemCount; index++ {
            log.Println(slice,slice[index],slice[index+1])
            comparisonsCount++
            if slice.Less(index+1, index) {
                slice.Swap(index, index+1)
                hasChanged = true
                swappedCount++
            }
        }
        if !hasChanged {
            break
        }
    }

    t := time.Since(t0).String()
    o = bubblesort_struct{ slice, repetitionCount, comparisonsCount, swappedCount, t }
    return
}


func server(rw http.ResponseWriter, req *http.Request) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }

    var r request_struct
    err = json.Unmarshal(body, &r)
    if err != nil {
        panic(err)
    }
    
    unsorted := make([]int, len(r.List), cap(r.List))
    copy(unsorted, r.List)

    response := resp_struct{r.List, Bubblesort(unsorted)}
    resp, err := json.Marshal(response)
    if err != nil {
        panic(err)
    }

    rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
    rw.Write(resp)
}


func main() {
    http.HandleFunc("/api/bubblesort", server)
    log.Fatal(http.ListenAndServe(":8080", nil))
}