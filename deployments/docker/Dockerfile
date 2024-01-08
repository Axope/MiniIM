FROM centos:7.9.2009

RUN [ "mkdir", "/app" ]
WORKDIR /app

COPY ../../bin/im /app
COPY ../../config.yaml /app
RUN [ "mkdir", "/app/logs" ]

EXPOSE 9876

CMD [ "/app/im" ]
