from nginx:alpine

run mkdir /msb

copy default.conf msb

CMD /bin/sh -c "cat /msb/default.conf | envsubst > /etc/nginx/conf.d/default.conf && exec nginx -g 'daemon off;'"
