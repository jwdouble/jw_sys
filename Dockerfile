FROM www.jwdouble.top:10443/k8s/base-image:v1

COPY ./bin/app ./config.yaml* /

CMD exec /app