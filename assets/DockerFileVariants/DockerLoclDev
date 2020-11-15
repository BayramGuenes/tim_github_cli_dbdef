FROM scratch as builder
COPY cmd/logger/main /cli/db/logger/main
COPY cmd/receiver/main /cli/db/receiver/main
COPY cmd/repo/main /cli/db/repo/main
#COPY assets/timFileSys/ /timFileSys
COPY assets/timFileSys/settings /cli/config


#===============================================================
# develop stage
#===============================================================
FROM alpine as develop

RUN apk update && \
    apk add \
    bash \
    curl \
    wget && \
  rm -rf  /var/cache/apk/*  


COPY --from=builder / .

CMD exec /bin/bash -c "trap : TERM INT; sleep infinity & wait"