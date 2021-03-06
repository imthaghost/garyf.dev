# Golang version
FROM golang:alpine3.12

# Enviornment variables
ENV APP_NAME garyf.dev
ENV PORT 8080

# Open system port
EXPOSE ${PORT}

# Working directory
WORKDIR /go/src/${APP_NAME}

COPY . /go/src/${APP_NAME}

# Install dependecies from mod file
RUN go mod download

# Build application
RUN go build -o ${APP_NAME}

# Run application
CMD ./${APP_NAME}
