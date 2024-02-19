import grpc
from concurrent import futures
import inference_pb2
import inference_pb2_grpc
import numpy as np
from PIL import Image
import io
from tensorflow import keras
import tensorflow as tf
import numpy as np
from keras.applications.inception_v3 import InceptionV3
from keras.applications.inception_v3 import preprocess_input
from keras.models import Model


def get_np_argmax(image_path, model_path):
    model = tf.keras.models.load_model(model_path)
    x = preprocess_input(
        np.expand_dims(
            keras.utils.img_to_array(
                keras.utils.load_img(image_path, target_size=(224, 224))
            ),
            axis=0,
        )
    )
    features = model.predict(x)
    # confidence
    confidence = int(np.max(features) * 100 / np.sum(features))
    return [np.argmax(features), confidence]


def main(image_path, model_path):
    print(get_np_argmax(image_path, model_path))


class InferenceServicer(inference_pb2_grpc.InferenceServicer):
    def Predict(self, request, context):
        image = Image.open(io.BytesIO(request.content))
        image.save("/tmp/image.jpg")
        prediction, confidence = get_np_argmax("/tmp/image.jpg", "zitch_v1")
        return inference_pb2.Prediction(
            value=int(prediction), confidence=int(confidence)
        )


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    inference_pb2_grpc.add_InferenceServicer_to_server(InferenceServicer(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
