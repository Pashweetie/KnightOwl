# The first image contains Go and is used only to compile the program.
FROM golang:1.26.5-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY apps/server ./apps/server
RUN CGO_ENABLED=0 go build -trimpath -ldflags='-s -w' -o /out/knightowl ./apps/server/cmd/knightowl

# The final image contains only the compiled program, not the Go compiler.
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/knightowl /knightowl
COPY config/development.json /app/config/development.json
WORKDIR /app
# Run as an unprivileged user so the program has less power if it is compromised.
USER nonroot:nonroot
EXPOSE 8787
ENTRYPOINT ["/knightowl"]
