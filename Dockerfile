FROM moby/buildkit:v0.9.3
WORKDIR /game
COPY game README.md /game/
ENV PATH=/game:$PATH
ENTRYPOINT [ "/bhojpur/game" ]