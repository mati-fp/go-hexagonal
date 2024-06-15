FROM golang:1.19

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go mod init github.com/mati-fp/go-hexagonal

RUN go get -u github.com/spf13/cobra@latest
RUN go install github.com/golang/mock/mockgen@v1.5.0
RUN go install github.com/spf13/cobra-cli@latest

RUN apt-get update && apt-get install sqlite3 -y

RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache

# Garante que www-data tenha permissão de escrita no diretório de trabalho
RUN chown -R www-data:www-data /go/src

CMD ["tail", "-f", "/dev/null"]
