#version of go and linux to use (from docker hub)
FROM golang:1.14-alpine

#make a directory to hold app
RUN mkdir /app

#copy files from current directory to newly created app directory
COPY . /app

#set working directory to new app directory 
WORKDIR /app

#get dependecies 
RUN go get -d -v ./...

#install dependencies
RUN go install -v ./...

#builds binary exe
RUN go build -o server .

#set app entry point to executable
CMD ["/app/server"]