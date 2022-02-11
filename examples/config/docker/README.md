### Start exporter with default settings

The container executes the application with the default numner of worker processes (2) and with the SSL warning enabled

'''shell

docker run -d -p 9491:9491  --rm --name pure-exporter pure-fb-prometheus-exporter:<version>
'''
