FROM nginx:1.10.3

RUN DEBIAN_FRONTEND=noninteractive \
    apt-get update -qq && \
    apt-get -y install curl runit unzip && \
    rm -rf /var/lib/apt/lists/*

ENV CT_URL https://releases.hashicorp.com/consul-template/0.18.2/consul-template_0.18.2_linux_amd64.zip
RUN curl -o consul-template.zip $CT_URL
RUN unzip consul-template.zip
RUN chmod a+x consul-template
RUN mv consul-template /usr/bin/consul-template

ADD nginx.service /etc/service/nginx/run
ADD consul-template.service /etc/service/consul-template/run

RUN rm -v /etc/nginx/conf.d/*
ADD nginx.conf /etc/consul-templates/nginx.conf

CMD ["/usr/bin/runsvdir", "/etc/service"]
