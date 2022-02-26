#!/usr/bin/env python

from flask import Flask, request, abort, make_response
from flask_httpauth import HTTPTokenAuth
from prometheus_client import generate_latest, CollectorRegistry, CONTENT_TYPE_LATEST
from pure_fb_openmetrics_exporter.flashblade_collector.collector import FlashbladeCollector
from pure_fb_openmetrics_exporter.flashblade_client.client import FlashbladeClient

import re
import logging

def create_app(disable_ssl_warn=False):

    app = Flask(__name__)
    app.logger.setLevel(logging.INFO)
    auth = HTTPTokenAuth(scheme='Bearer')

    @auth.verify_token
    def verify_token(token):
        pattern_str = "^T-[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
        regx = re.compile(pattern_str)
        match = regx.search(token)
        return token if match is not None else False

    @app.route('/')
    def route_index():
        """Display an overview of the exporters capabilities."""
        return '''
<h1>Pure Storage Flashblade OpenMetrics Exporter</h1>
<table>
    <thead>
        <tr>
        <td>Type</td>
        <td>Endpoint</td>
        <td>GET parameters</td>
        <td>Description</td>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Full metrics</td>
            <td><a href="/metrics?endpoint=host">/metrics</a></td>
            <td>endpoint</td>
            <td>All array metrics. Expect slow response time.</td>
        </tr>
        <tr>
            <td>Array metrics</td>
            <td><a href="/metrics/array?endpoint=host">/metrics/array</a></td>
            <td>endpoint</td>
            <td>Provides only array related metrics.</td>
        </tr>
        <tr>
            <td>Client metrics</td>
            <td><a href="/metrics/clients?endpoint=host">/metrics/clients</a></td>
            <td>endpoint</td>
            <td>Provides only client related metrics. This is the most time expensive query</td>
        </tr>
        <tr>
            <td>Quota metrics</td>
            <td><a href="/metrics/usage?endpoint=host">/metrics/usage</a></td>
            <td>endpoint</td>
            <td>Provides only quota related metrics.</td>
        </tr>
    </tbody>
</table>
'''

    @app.route('/metrics/<m_type>', methods=['GET'])
    @auth.login_required
    def route_flashblade(m_type: str):
        """Produce FlashBlade metrics."""
        if not m_type in ['all', 'array', 'clients', 'usage']:
            abort(400)
        registry = CollectorRegistry()
        collector = FlashbladeCollector
        try:
            endpoint = request.args.get('endpoint', None)
            token = auth.current_user()
            fb_client = FlashbladeClient(endpoint, token, disable_ssl_warn)
            registry.register(collector(fb_client, request=m_type))
        except Exception as e:
            app.logger.warning('%s: %s', collector.__name__, str(e))
            abort(500)

        resp = make_response(generate_latest(registry), 200)
        resp.headers['Content-type'] = CONTENT_TYPE_LATEST
        return resp

    @app.route('/metrics', methods=['GET'])
    def route_flashblade_all():
        return route_flashblade('all')

    @app.errorhandler(400)
    def route_error_400(error):
        """Handle invalid request errors."""
        return 'Invalid request parameters', 400

    @app.errorhandler(404)
    def route_error_404(error):
        """ Handle 404 (HTTP Not Found) errors."""
        return 'Not found', 404

    @app.errorhandler(500)
    def route_error_500(error):
        """Handle server-side errors."""
        return 'Internal server error', 500

    return app

# Run in debug mode when not called by WSGI
if __name__ == "__main__":
    app = create_app(disable_ssl_warn=True)
    app.logger.setLevel(logging.DEBUG)
    app.logger.debug('running in debug mode...')
    app.run(host="0.0.0.0", port=8080, debug=True)
