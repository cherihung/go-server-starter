FROM cherihung01/golang-builder:alpine AS builder
LABEL stage=builder

WORKDIR /workspace

COPY . .

RUN GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux go build -a -installsuffix -o ${MODULE_NAME}

FROM alpine AS final

ARG ENV
RUN echo $ENV

WORKDIR /src

COPY --from=builder /workspace/go-server-starter .
COPY --from=builder /workspace/_envs/env_$ENV.yaml ./_envs/
COPY --from=builder /workspace/_certs/server.crt ./_certs/
COPY --from=builder /workspace/_certs/server.key ./_certs/

CMD [ "./go-server-starter" ]


