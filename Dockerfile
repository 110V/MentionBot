FROM golang
WORKDIR /go/src/github.com/110V/MentionBot
ADD . .
RUN go get ./... && go install
WORKDIR /home
CMD ["MentionBot"]
