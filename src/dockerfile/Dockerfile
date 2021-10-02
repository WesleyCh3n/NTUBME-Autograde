FROM debian:stable-slim

RUN apt-get update && \
    apt install jq curl gosu -y

RUN curl -fLo /usr/bin/yq https://github.com/mikefarah/yq/releases/download/3.4.0/yq_linux_amd64 && chmod +x /usr/bin/yq
ADD ./ga /usr/bin/ga
RUN chmod +x /usr/bin/ga


ENV USER=wesley USER_ID=1000 USER_GID=1000

# now creating user
RUN groupadd --gid "${USER_GID}" "${USER}" && \
    useradd \
      --uid ${USER_ID} \
      --gid ${USER_GID} \
      --shell /bin/bash \
      --create-home \
      ${USER}

COPY ./dockerfile/entrypoint.sh /
RUN  chmod u+x /entrypoint.sh

WORKDIR /workdir/
ENTRYPOINT ["/entrypoint.sh"]
