 #### Stage 1
 FROM golang:1.21.4-alpine AS build
 RUN mkdir /app
 COPY . /app
 WORKDIR /app
 RUN go build -o main .
 RUN chmod +x /app/main

 #### Stage 2
 FROM scratch
 COPY --from=build /app/main /main
 COPY --from=build /app/.env /.env
 COPY --from=build /app/static /static
 ENTRYPOINT [ "/main" ]