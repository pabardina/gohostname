FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD gohostname /
ENTRYPOINT ["/gohostname"]


