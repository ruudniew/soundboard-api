FROM golang:1.12-alpine

# update and install dependency
RUN apk update && apk upgrade

# create destination directory
RUN mkdir -p /ruudniew-sbapi
WORKDIR /ruudniew-sbapi

# copy the app, note .dockerignore
COPY . /ruudniew-sbapi/

# expose 3300 on container
EXPOSE 3300:3300

# start the app

RUN go build -mod=vendor -o ./sbapi ./cmd/sbapi

RUN chmod a+x ./sbapi

CMD [ "./sbapi" ]
