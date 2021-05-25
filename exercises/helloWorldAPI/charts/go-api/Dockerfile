FROM alpine/git:latest

WORKDIR /

RUN wget https://github.com/mikefarah/yq/releases/download/v4.9.3/yq_linux_amd64
RUN mv yq_linux_amd64 yq
RUN chmod 755 yq

ENTRYPOINT ["yq"]
CMD ["-help"]
