FROM debian:10-slim

ADD bin/parkinglistservice /app/bin/parkinglistservice

WORKDIR /app/bin

EXPOSE 8000

CMD [ "/app/bin/parkinglistservice" ]