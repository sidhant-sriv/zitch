import grpc
import inference_pb2
import inference_pb2_grpc
import time


def run():
    with open("image.jpg", "rb") as f:
        image_bytes = f.read()

    with grpc.insecure_channel("35.192.204.153") as channel:
        stub = inference_pb2_grpc.InferenceStub(channel)
        response = stub.Predict(inference_pb2.Image(content=image_bytes))
    s = time.time()
    print("Prediction:", response.value)
    print("Confidence:", response.confidence)
    print("Time:", time.time() - s)


if __name__ == "__main__":
    run()
