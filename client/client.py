import logging
import grpc

import content_pb2
import content_pb2_grpc

def run():
    with grpc.insecure_channel('localhost:9000') as channel:
        stub = content_pb2_grpc.ContentManagementStub(channel)
        response = stub.CreateContent(content_pb2.NewContent(text="Hello from 🐍"))
        print("Python received: ")
        print("Text: " + response.text)
        print("Id: " + str(response.id))

if __name__ == '__main__':
    logging.basicConfig()
    run()
