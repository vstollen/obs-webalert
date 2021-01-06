# Heavily inspired by https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#use-multi-stage-builds
FROM golang:alpine AS build

ENV APP_HOME /go/src/webalert

# Copy and build project
COPY . $APP_HOME

WORKDIR $APP_HOME
RUN go mod download
RUN go mod verify
RUN go build -o /bin/webalert

# Assemble the resulting image
FROM alpine

COPY --from=build /bin/webalert /bin/webalert
COPY --from=build /go/src/webalert/static/ /static/

EXPOSE 8080

ENTRYPOINT [ "/bin/webalert" ]
