FROM golang AS builder
# base image for creatign our a new image

# всё оке
LABEL platform social_network 
LABEL author Loid Forger  

WORKDIR /app

# install third party libs
COPY go.mod go.sum ./
RUN go mod download

# # update packages and install postgresql-client for working with pg
# RUN apt-get update
# RUN apt-get install -y postgresql-client

COPY . .

EXPOSE 8080:8080

RUN go build -o social_network ./cmd/social-network/main.go

FROM alpine

WORKDIR /app

# copy bin
COPY --from=builder /app/social_network .

RUN echo Social-Network is running...
CMD ["./social_network"]