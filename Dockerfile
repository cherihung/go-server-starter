FROM golang:alpine AS builder
LABEL stage=builder

RUN apk add --no-cache gcc libc-dev

WORKDIR /workspace

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix -o ${MODULE_NAME}

FROM alpine AS final

ARG ENV
RUN echo $ENV

WORKDIR /src

COPY --from=builder /workspace/go-server-starter .
COPY --from=builder /workspace/_envs/env_$ENV.yaml ./_envs/
COPY --from=builder /workspace/_certs/server.crt ./_certs/
COPY --from=builder /workspace/_certs/server.key ./_certs/

CMD [ "./go-server-starter" ]


