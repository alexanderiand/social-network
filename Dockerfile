FROM golang
# base image for creatign our a new image

LABEL platform social_network 
LABEL author Loid Forger  

WORKDIR /

COPY . .

# update packages and install postgresql-client for working with pg
RUN apt-get update
RUN apt-get install -y postgresql-client



# install third party libs
RUN go mod download

# 
RUN go build -o social_network ./cmd/social-network/main.go

CMD ["./social_network"]