#!/bin/ash

FROM alpine:3.16.2

COPY ./thin-slack-bot /
RUN chmod +x /thin-slack-bot

CMD [ "/thin-slack-bot" ]