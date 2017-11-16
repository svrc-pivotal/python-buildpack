#!/usr/bin/env python

from http.server import BaseHTTPRequestHandler, HTTPServer
from animals import Mammals

class testHTTPServer_RequestHandler(BaseHTTPRequestHandler):
  def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type','text/html')
        self.end_headers()

        mammals = Mammals()
        self.wfile.write(bytes("Hello world!", "utf8"))
        for member in mammals.getMembers():
           self.wfile.write(bytes(member, "utf8"))

        return

def run():
  server_address = ('0.0.0.0', 8080)
  httpd = HTTPServer(server_address, testHTTPServer_RequestHandler)
  print('running server...')
  httpd.serve_forever()

run()
