# 1.14-alpine
# pull from the exact digest for security purpose to make sure it is the exact image you want
#
FROM golang@sha256:e484434a085a28801e81089cc8bcec65bc990dd25a070e3dd6e04b19ceafaced AS goBuilder

# Install git for fetching go modules
#
RUN apk update && \
    apk add --no-cache git

WORKDIR $GOPATH/src/github.com/mikejianzhang/jenkins-operator
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/jenkins-operator ./cmd/manager

# run stage
FROM alpine:3.10

USER nobody

COPY --from=goBuilder /go/bin/jenkins-operator /usr/local/bin/jenkins-operator

CMD [ "/usr/local/bin/jenkins-operator" ]