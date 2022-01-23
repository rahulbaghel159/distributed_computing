from flask import Flask
from flask_restful import Resource,  Api, reqparse
import pandas as pd
import argparse

app=Flask(__name__)
api=Api(app)

class Message(Resource):
    def get(self):
        parser = reqparse.RequestParser()
        parser.add_argument("lamport_time", required=True)
        parser.add_argument("local_time", required=True)
        args = parser.parse_args()

        print("request param", args["lamport_time"])
        print("request param", args["local_time"])

        return {"status":"success"},200
    pass

api.add_resource(Message, "/message")

if __name__ == "__main__":
    # Construct the argument parser
    ap = argparse.ArgumentParser()

    # Add the arguments to the parser
    ap.add_argument("-p", "--port", required=True,help="port")
    
    args = vars(ap.parse_args())

    port = format(int(args['port']))

    app.run(port=port)