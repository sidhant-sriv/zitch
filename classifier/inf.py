from tensorflow import keras
import tensorflow as tf
import numpy as np
from keras.applications.inception_v3 import InceptionV3
from keras.applications.inception_v3 import preprocess_input
from keras.models import Model

def get_np_argmax(image_path, model_path):
    model = tf.keras.models.load_model(model_path)
    x = preprocess_input(np.expand_dims(keras.utils.img_to_array(keras.utils.load_img(image_path, target_size=(224, 224))), axis=0)) 
    features = model.predict(x)
    return np.argmax(features)

def main(image_path, model_path):
    print(get_np_argmax(image_path, model_path))
