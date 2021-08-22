FROM chromedp/headless-shell:latest
WORKDIR /opt/web-screenshot
MAINTAINER debug.icu
ADD ./ /opt/web-screenshot
RUN chmod 777 /opt/web-screenshot/web-screenshot-service
ENTRYPOINT ["/opt/web-screenshot/web-screenshot-service"]