FROM gradle:latest
RUN apt-get update
RUN apt-get install postgresql sudo -y

RUN git clone https://github.com/C-Otto/lnd-manageJ.git && \
    cd lnd-manageJ && \
    git checkout 493b11cbd325e00a84a01ce118b24dbdfcb6f9d1

WORKDIR lnd-manageJ
RUN gradle application:bootJar

RUN mkdir -p /root/.config
RUN echo "[lnd]" >> /root/.config/lnd-manageJ.conf
RUN echo "host=node-start" >> /root/.config/lnd-manageJ.conf
RUN echo "macaroon_file=/cfg/start/admin.macaroon" >> /root/.config/lnd-manageJ.conf
RUN echo "cert_file=/cfg/start/tls.cert" >> /root/.config/lnd-manageJ.conf

RUN echo "[pickhardt-payments]" >> /root/.config/lnd-manageJ.conf
RUN echo "quantization=5000" >> /root/.config/lnd-manageJ.conf

RUN echo "server.address=0.0.0.0" >> /root/override.properties

EXPOSE 8081
CMD /etc/init.d/postgresql start && \
    sudo -u postgres psql -c "CREATE USER bitcoin WITH PASSWORD 'unset'" && \
    sudo -u postgres createdb lndmanagej -O bitcoin && \
    java -jar application/build/libs/application-boot.jar --spring.config.location=classpath:application.properties,/root/override.properties


# docker build -t lnd-managej .
# docker run --network host -v /home/xxx/.lnd/:/root/.lnd lnd-managej
