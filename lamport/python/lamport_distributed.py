from pip._vendor import requests
from datetime import datetime 
from flask import Flask
from flask_restful import Resource,  Api, reqparse
import pandas as pd
import argparse
from multiprocessing import Process
import time

def local_time(counter):
    return ' (LAMPORT_TIME={}, LOCAL_TIME={})'.format(counter,datetime.now())

def event(pid, counter):
    counter += 1
    print('Something happened in {} !'.\
          format(pid) + local_time(counter))
    return counter

def recv_message(counter, localtime):
    # recieve message code
    print('Message received at ' + counter  + local_time(localtime))

def send_message(port, counter):
    # api-endpoint
    URL = "http://127.0.0.1:" + str(port) + "/message"
    params = {"lamport_time":counter, "local_time":datetime.now()}

    # send message code
    print('Message sent from ' + str(port) + local_time(counter))
    print('URL: ' + URL)
    
    # sending get request and saving the response as response object
    r = requests.get(url = URL, params=params)

#server code
app=Flask(__name__)
api=Api(app)

class Message(Resource):
    def get(self):
        parser = reqparse.RequestParser()
        parser.add_argument("lamport_time", required=True)
        parser.add_argument("local_time", required=True)
        args = parser.parse_args()

        print("message recieved")
        print("request param", args["lamport_time"])
        print("request param", args["local_time"])

        recv_message(args["lamport_time"], args["local_time"])

        return {"status":"success"},200
    pass

api.add_resource(Message, "/message")

def api_process(port):
    app.run(port=port)

if __name__ == "__main__":
    # Construct the argument parser
    ap = argparse.ArgumentParser()

    # Add the arguments to the parser
    ap.add_argument("-p", "--port", required=True,help="port")
    
    args = vars(ap.parse_args())

    port = format(int(args['port']))

    process1 = Process(target=api_process, args=(port,))

    process1.start()
    time.sleep(2)
    
    send_message(8000, 1)