FROM golang:1.17-alpine as builder
RUN mkdir -p /go/src/mutant
WORKDIR /go/src/mutant
COPY . /go/src/mutant

#Build binary program
RUN GOPATH=/go/src \
    GOROOT=/usr/local/go/ \
    PATH=/go/bin:$PATH \
    CGO_ENABLED=0 \
    GOODS=linux \
    GOARCH=amd64 \
    go build -a -o /app.linux cmd/main.go 

#generate clean, final image for end users
FROM alpine:3.11.3
EXPOSE 8080
WORKDIR /home/user/app/
COPY --from=builder /app.linux .
#Run the binary program 
ENV PS1="\[\e[1;32m\]mutant$ \[\e[1;31m\][`pwd`] # \[\e[0m\]"

CMD ["./app.linux"]