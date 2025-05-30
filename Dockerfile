FROM golang:1.24-alpine AS build

WORKDIR /usr/local/ledsproj

COPY . .

# RUN go mod download
# RUN go mod tidy
RUN CGO=0 go build -mod=vendor -ldflags "-s -w" -o ledsproj .

FROM alpine:3.21

COPY --from=build /usr/local/ledsproj/ledsproj /usr/local/bin

ENTRYPOINT [ "ledsproj" ]
