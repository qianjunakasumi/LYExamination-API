FROM alpine:3.14.0

COPY api /bin/lyexamination-api

RUN chmod 777 /bin/lyexamination-api

CMD lyexamination-api
