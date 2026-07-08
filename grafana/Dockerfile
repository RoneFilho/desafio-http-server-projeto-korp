FROM grafana/grafana:latest

COPY dashboards.yml /etc/grafana/provisioning/dashboards/

COPY http-server-projeto-korp-dashboard.json /etc/grafana/provisioning/dashboards/

COPY datasources.yml /etc/grafana/provisioning/datasources/

EXPOSE 3000