import grpc
import sys
from os import path
import time

# Include protobuf definitions from parallel directory
proto_path = path.abspath(path.dirname(path.abspath(__file__)) + "/../proto/")
sys.path.append(proto_path)
import hello_pb2
import hello_pb2_grpc


def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = hello_pb2_grpc.DiffSubscriberStub(channel)

    subscribeMessage = hello_pb2.DiffSubscribe(
        path="/user/3",
        period=1,
        subscriberId="orm")

    try:
        for response in stub.SubscribeStream(subscribeMessage):
            print("Response: %s" % response.responseData)

    except grpc.RpcError as e:
        status_code = e.code()

        if e.code() == grpc.StatusCode.UNKNOWN:
            additional_info = "Server shutdown or crashed."
        elif e.code() == grpc.StatusCode.UNAVAILABLE:
            additional_info = "Server probably not up."
        else:
            additional_info = "Unknown reason."
        print("gRPC error '%s'. %s" % (e.details(), additional_info))

if __name__ == '__main__':
    run()
