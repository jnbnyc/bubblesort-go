jnbnyc/bubblesort-go
==

    $ curl -isS machine:8080/api/bubblesort -d '{ "list": [ 1, 7, 4] }'
    HTTP/1.1 200 OK
    Content-Type: application/json; charset=UTF-8
    Date: Mon, 21 Dec 2015 04:06:28 GMT
    Content-Length: 96

    {"original":[1,7,4],"sorted":[1,4,7],"reps":2,"comparisons":3,"swaps":1,"duration":"616.548µs"}


