FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY . ./

COPY vendor ./
RUN ls

#  -mod=vendor to not re-download vendor package (for now)
# build binary is output to server
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -v -o server ./cmd/server/

FROM alpine:3.15.0 AS bin
# RUN apk add --no-cache ca-certificate=20211220-r0
RUN apk add --no-cache tzdata

ENV TZ=Asia/Bangkok
ENV HOST=0.0.0.0

# Copy binary to image from builder stage
COPY --from=builder /app/server /server

CMD ["/server"]
