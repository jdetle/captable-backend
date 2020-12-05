FROM golang:latest
ENV PORT 8000
COPY . /go/src/github.com/jdetle/captable-backend
WORKDIR /go/src/github.com/jdetle/captable-backend
COPY main.go .
RUN go install -mod=vendor
EXPOSE ${PORT}
CMD ["captable-backend"]
