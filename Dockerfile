FROM golang:1.9

ADD . /go/src/github.com/kalimatas/slack-cowbot

RUN go install github.com/kalimatas/slack-cowbot/cmd/cowbot

ENTRYPOINT /go/bin/cowbot

EXPOSE 80
