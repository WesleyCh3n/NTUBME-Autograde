FROM debian:stable-slim

RUN apt-get update && \
    apt install jq curl -y
ADD ./yq_linux_amd64 /usr/bin/yq
ADD ./ga /usr/bin/ga
RUN chmod +x /usr/bin/yq
RUN chmod +x /usr/bin/ga

WORKDIR /workdir/
ENTRYPOINT ["ga"]
