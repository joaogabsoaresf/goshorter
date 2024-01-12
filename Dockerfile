FROM golang:alpine AS build

WORKDIR /opt/app

COPY /. /opt/app/

RUN go mod download

RUN go build main.go

FROM golang:alpine as prod

WORKDIR /opt/app

COPY --from=build  /opt/app/dotenv-files ./dotenv-files

COPY --from=build  /opt/app/main ./

CMD ["./main"]