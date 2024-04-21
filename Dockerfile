FROM cgr.dev/chainguard/go:latest as build

WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=./booking/go.sum,target=go.sum \
    --mount=type=bind,source=./booking/go.mod,target=go.mod \
    go mod download

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,rw,source=./booking,target=. \
    go build \
    -ldflags "-linkmode external -extldflags -static" \
    -o api

FROM cgr.dev/chainguard/static:latest

LABEL version="X.Y.Z"

# USER nonroot

# ENV TINI_VERSION v0.19.0
# ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static ./tini
# # RUN chmod +x ./tini

COPY --from=build /app/api .

EXPOSE ${BOOKING_PORT}

# ENTRYPOINT ["./tini", "--"]

CMD ["/api"]
