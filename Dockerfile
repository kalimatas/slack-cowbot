FROM golang:1.9

RUN apt-get update && apt-get install -y cowsay

ADD . /go/src/github.com/kalimatas/slack-cowbot

RUN go install github.com/kalimatas/slack-cowbot/cmd/cowbot

CMD ["/go/bin/cowbot"]
