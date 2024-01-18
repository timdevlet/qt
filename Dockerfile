# Build stage
FROM golang:1.21 AS BuildStage

RUN mkdir /app

WORKDIR /app

COPY go.mod /app

RUN go mod download

COPY . /app

RUN APP="cli" ./hacks/build.sh

#---------------------------------------------------------------------------------------------

# Deploy stage
FROM golang:1.21-alpine as Deploy

EXPOSE 8080

# [copy binaries] 
COPY --from=BuildStage ./app/bin/cli /usr/local/bin/

WORKDIR /usr/local/bin/
