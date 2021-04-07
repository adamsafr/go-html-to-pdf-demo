FROM golang:1.15.7-buster

RUN apt-get update && apt-get install -y xvfb libfontconfig wkhtmltopdf

ENV GO111MODULE=on
ENV GOFLAGS="-mod=vendor"
ENV APP_USER app
ENV APP_HOME /go/src/html-to-pdf
ARG GROUP_ID
ARG USER_ID

RUN groupadd --gid $GROUP_ID app && useradd -m -l --uid $USER_ID --gid $GROUP_ID $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME

USER $APP_USER
WORKDIR $APP_HOME

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .