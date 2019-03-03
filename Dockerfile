# build stage
FROM golang:1.12.0-alpine3.9 AS build-env
WORKDIR /app
RUN apk add --no-cache git
ADD go.* /app/
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 go build -o /app/docker-credential-gitlab-login

# final stage
FROM gcr.io/kaniko-project/executor:debug
ENTRYPOINT ["sh", "-c"]
RUN ["ln", "-s", "/kaniko/executor", "/kaniko/build"]
COPY --from=build-env /app/docker-credential-gitlab-login /kaniko/
