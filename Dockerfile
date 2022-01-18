FROM python:slim

# Run as non-root user
RUN groupadd app && \
    useradd -d /app -g app app && \
    mkdir /app && \
    chown -R app:app /app

WORKDIR /app

USER app

COPY dist/pure-fb-prometheus-exporter-0.0.1.tar.gz .
COPY src/pure_fb_prometheus_exporter/pure_fb_exporter.py .

# Install package - dependencies are already included in the package itself
RUN pip install --upgrade pip && \
    pip install --no-cache-dir pure-fb-prometheus-exporter-0.0.1.tar.gz

RUN rm pure-fb-prometheus-exporter-0.0.1.tar.gz

ENV PATH "$PATH:/app/.local/bin"

# Configure the image properties
# gunicorn settings: bind any, 2 threads, log to
# stdout/stderr (docker/k8s handles logs), anonymize request URL
# end of log shows request time in seconds and size in bytes
ENV GUNICORN_CMD_ARGS="--bind=0.0.0.0:9491 \
    --workers=16 \
    --access-logfile=- \
    --error-logfile=- \
    --access-logformat=\"%(t)s %(h)s %(U)s %(l)s %(T)s %(B)s\""
EXPOSE 9491
ENTRYPOINT ["gunicorn"]
CMD ["pure_fb_exporter:create_app()"]
