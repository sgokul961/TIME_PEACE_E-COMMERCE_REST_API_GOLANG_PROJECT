FROM golang:1.21.0-alpine AS build_stage
WORKDIR /home/project/
COPY ./ /home/project/
RUN mkdir -p /home/project
RUN go mod download
RUN go build -v -o /home/build/api ./cmd/api

FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /home/build/api /api
COPY --from=build-stage /home/project/templates /templates
COPY --from=build-stage /home/project/.env /
EXPOSE 3000
CMD ["/api"]
