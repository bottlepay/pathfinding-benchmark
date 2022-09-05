FROM gradle:latest

RUN git clone https://github.com/C-Otto/lnd-manageJ.git && \
    cd lnd-manageJ && \
    git checkout c57eca472aafe2546768a6f68ffe39499919e36b

WORKDIR lnd-manageJ
RUN gradle application:bootJar

RUN mkdir -p /root/.config
RUN echo "[lnd]" >> /root/.config/lnd-manageJ.conf
RUN echo "host=node-start" >> /root/.config/lnd-manageJ.conf
RUN echo "macaroon_file=/cfg/start/admin.macaroon" >> /root/.config/lnd-manageJ.conf
RUN echo "cert_file=/cfg/start/tls.cert" >> /root/.config/lnd-manageJ.conf

RUN echo "[pickhardt-payments]" >> /root/.config/lnd-manageJ.conf
RUN echo "enabled=true" >> /root/.config/lnd-manageJ.conf

RUN echo "server.address=0.0.0.0" >> /root/override.properties

EXPOSE 8081
CMD java -Dspring.profiles.active=h2 -jar application/build/libs/application-boot.jar --spring.config.location=classpath:application.properties,/root/override.properties
