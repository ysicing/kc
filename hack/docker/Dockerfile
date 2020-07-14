FROM ysicing/alpine as build

RUN apk add lxcfs containerd

FROM ysicing/alpine

COPY --from=build /usr/bin/lxcfs /usr/bin/lxcfs
COPY --from=build /usr/lib/*fuse* /usr/lib/
COPY --from=build /usr/bin/ctr /usr/bin/ctr

RUN apk update \
    && apk add --no-cache iputils iproute2 drill busybox-extras mtr \
    && rm -rf /var/cache/apk/*