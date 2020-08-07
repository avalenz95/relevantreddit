#version of go and linux to use (from docker hub)
FROM golang:1.14-alpine

#make a directory to hold api
RUN mkdir /api

#copy files from current directory to newly created app directory
COPY . /api

#set working directory to new app directory 
WORKDIR /api

#get dependecies 
RUN go get -d -v ./...

#install dependencies
RUN go install -v ./...

#builds binary exe
RUN go build -o server .

EXPOSE 8080

#set app entry point to executable
CMD ["/api/server"]