FROM golang:alpine as final

RUN addgroup -g 1000 -S guardrails \
    && adduser  guardrails -u 1000 -D -S -s /bin/bash -G guardrails \
    && mkdir -p /data \
    && chown -R guardrails:guardrails /data

USER 1000

ENV TZ=Asia/Bangkok

ENV CONFIG_PATH=$APP_DIR/configs/

COPY dist/grchallengeapi-linux-x64 /data/app/
COPY configs /data/app/configs

WORKDIR /data/app
CMD ["./grchallengeapi-linux-x64"]