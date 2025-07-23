FROM golang:1.24-alpine AS router

WORKDIR /app/router
COPY router /app/router
RUN go mod download 
RUN GOOS=linux go build -o ./server .


FROM node:latest AS frontend
WORKDIR /app/src
COPY frontend_src /app/src

RUN npm install
RUN npm run build


FROM scratch
WORKDIR /app
COPY --from=router /app/router/server /app/router/server
COPY --from=frontend /app/public /app/public
EXPOSE 8080
CMD ["/app/router/server"]

