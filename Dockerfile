FROM alpine
ADD renderer /renderer
ENTRYPOINT [ "/renderer" ]
