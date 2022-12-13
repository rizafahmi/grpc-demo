import logging
import grpc

import content_pb2
import content_pb2_grpc

def run():
    with grpc.insecure_channel('localhost:9000') as channel:
        stub = content_pb2_grpc.ContentManagementStub(channel)

if __name__ == '__main__':
    logging.basicConfig()
    run()
