FROM --platform=linux/amd64 golang:1.22-alpine3.19 AS app_builder

WORKDIR /code

ENV PORT=8080

ENV CGO_ENABLED=1
ENV GOPATH=/code
ENV GOCACHE=/go-build

RUN <<EOF
apk update
apk add gcc libc-dev
EOF

COPY ./go.mod /code/

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

COPY *.go .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o bin/mc-hints *.go

FROM --platform=linux/amd64 node:20.4-alpine3.17 AS styles_builder

WORKDIR /styles

COPY package.json yarn.lock tailwind.config.js /styles/
COPY templates/ /styles/templates/
COPY styles/ /styles/styles/

RUN yarn install --dev

RUN yarn build-css-prod
# RUN pwd

FROM --platform=linux/amd64 alpine:3.17 AS final

EXPOSE 8080

WORKDIR /app

COPY --from=app_builder /code/bin/mc-hints /app/
COPY static/ ./static
COPY templates/ ./templates
COPY images/ ./images
COPY hints/ ./hints
COPY additional_tasks/ ./additional_tasks
COPY --from=styles_builder /styles/static/styles.css /app/static/

CMD ["/app/mc-hints"]