FROM gradle:latest
RUN apt-get update
RUN apt-get install postgresql sudo -y

RUN /etc/init.d/postgresql start && \
    sudo -u postgres psql -c "CREATE USER bitcoin WITH PASSWORD 'unset'" && \
    sudo -u postgres createdb lndmanagej -O bitcoin && \
    /etc/init.d/postgresql stop

RUN git clone https://github.com/C-Otto/lnd-manageJ.git && \
    cd lnd-manageJ && \
    git checkout 777a0e9c2412b3b7bd682b3cd1b4832472135650

WORKDIR lnd-manageJ
RUN gradle application:bootJar

RUN mkdir -p /root/.config
RUN echo "[lnd]" >> /root/.config/lnd-manageJ.conf
RUN echo "host=node-start" >> /root/.config/lnd-manageJ.conf
RUN echo "macaroon_file=/cfg/start/admin.macaroon" >> /root/.config/lnd-manageJ.conf
RUN echo "cert_file=/cfg/start/tls.cert" >> /root/.config/lnd-manageJ.conf

RUN echo "server.address=0.0.0.0" >> /root/override.properties

EXPOSE 8081
CMD /etc/init.d/postgresql start && \
    java -jar application/build/libs/application-boot.jar --spring.config.location=classpath:application.properties,/root/override.properties


# docker build -t lnd-managej .
# docker run --network host -v /home/xxx/.lnd/:/root/.lnd lnd-managej
