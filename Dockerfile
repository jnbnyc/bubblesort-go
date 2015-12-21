FROM golang

ADD ./*.go ./*.sh /go/
ENTRYPOINT ["go", "run", "/go/bubblesort.go"]
#ENTRYPOINT /go/sleepd.sh
