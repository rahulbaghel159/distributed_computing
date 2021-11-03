from pip._vendor import requests
from datetime import datetime 
from http.server import BaseHTTPRequestHandler, HTTPServer
import http.server
import threading

def local_time(counter):
    return ' (LAMPORT_TIME={}, LOCAL_TIME={})'.format(counter,
                                                     datetime.now())

def calc_recv_timestamp(recv_time_stamp, counter):
    return max(recv_time_stamp, counter) + 1

def event(pid, counter):
    counter += 1
    print('Something happened in {} !'.\
          format(pid) + local_time(counter))
    return counter

def send_message(pid, counter):
    # send message code
    print('Message sent from ' + str(pid) + local_time(counter))
    
    # api-endpoint
    URL = "http://localhost:8000"

    # sending get request and saving the response as response object
    r = requests.get(url = URL)

def recv_message(pipe, pid, counter):
    # recieve message code
    print('Message received at ' + str(pid)  + local_time(counter))

class S(BaseHTTPRequestHandler):
    def _set_response(self):
        self.send_response(200)
        self.end_headers()

    def do_GET(self):
        self._set_response()
        print("get from handler")

def start_server(handler_class=S):
    hostName = "localhost"
    serverPort = 8000

    webServer = HTTPServer((hostName, serverPort), handler_class)
    print("Server started http://%s:%s" % (hostName, serverPort))

    try:
        webServer.serve_forever()
    except KeyboardInterrupt:
        pass

    webServer.serve_forever()

if __name__ == "__main__":
    daemon = threading.Thread(name='daemon_server',
                        target=start_server,
                        args=())
    daemon.setDaemon(True)
    daemon.start()

    send_message(1, 1)
