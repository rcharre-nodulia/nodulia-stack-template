FROM node:23-alpine AS build-css
WORKDIR /app
COPY . .

RUN npm install
RUN npm run build 
FROM golang:1.23.4-alpine AS build-go

WORKDIR /app
COPY --from=build-css /app .

RUN go mod download && \
  go test ./... && \
  go build -o build/bin

FROM scratch
COPY --from=build-go app/build/bin /app/build/bin

CMD ["/app/bin"]
