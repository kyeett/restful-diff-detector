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
    options = [('grpc.max_reconnect_backoff_ms', 5000)]
    channel = grpc.insecure_channel('localhost:50051', options=options)

    stub = hello_pb2_grpc.DiffSubscriberStub(channel)

    subscribeMessage = hello_pb2.DiffSubscribe(
        path="/user/3",
        period=1,
        subscriberId="orm")

    while(True):
        start = time.time()

        try:
            for notification in stub.SubscribeStream(subscribeMessage):
                print("Received a notification: %s" % notification.responseData)

        except grpc.RpcError as e:
            if e.code() == grpc.StatusCode.UNAVAILABLE:
                additional_info = "Server probably not up. "
            elif e.code() == grpc.StatusCode.UNKNOWN:
                additional_info = "Server shutdown or crashed. "
            else:
                additional_info = "Unknown reason. "
            print("gRPC error '%s'. %s" % (e.details(), additional_info), end='', flush=True)

        print("Sleep for 1 second then reconnect")
        time.sleep(1)

        done = time.time()
        elapsed = done - start
        print("Elapsed time", elapsed)

        grpc


if __name__ == '__main__':
    run()
