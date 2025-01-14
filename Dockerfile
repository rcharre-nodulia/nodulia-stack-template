FROM node:23-alpine AS build-css
WORKDIR /app
COPY . .

RUN npm install && \
  cp -r static build/static && \
  npm run build:css

FROM golang:1.23.4-alpine AS build-go

WORKDIR /app
COPY --from=build-css /app .

RUN go mod download && \
  go test ./... && \
  go build -o build/bin

FROM scratch
COPY --from=build-go app/build/bin /app/build/bin

CMD ["/app/bin"]
