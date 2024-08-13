import json
import pandas as pd
import numpy as np
from sklearn.cluster import DBSCAN
from sklearn.preprocessing import StandardScaler
from joblib import Parallel, delayed

# Function to process chunks of data
def process_chunk(chunk):
    # Preprocess data
    df = pd.DataFrame(chunk)
    scaler = StandardScaler()
    df[['x', 'y']] = scaler.fit_transform(df[['x', 'y']])
    
    # Cyclical time features
    df['hour'] = (df['time'] % 24).astype(float)
    df['hour_sin'] = np.sin(2 * np.pi * df['hour'] / 24)
    df['hour_cos'] = np.cos(2 * np.pi * df['hour'] / 24)
    
    return df.drop(columns=['time', 'hour'])

# Stream JSON data in chunks
def json_streamer(file_path, chunk_size=1000):
    with open(file_path, 'r') as file:
        data_chunk = []
        for line in file:
            data_chunk.append(json.loads(line))
            if len(data_chunk) >= chunk_size:
                yield process_chunk(data_chunk)
                data_chunk = []
        if data_chunk:  # Yield any remaining data
            yield process_chunk(data_chunk)

# Function to apply DBSCAN in parallel
def parallel_dbscan(chunk, eps, min_samples):
    db = DBSCAN(eps=eps, min_samples=min_samples).fit(chunk[['x', 'y', 'hour_sin', 'hour_cos']])
    return db.labels_

# Main processing function
def process_large_json(file_path, eps=0.3, min_samples=5, n_jobs=-1):
    all_labels = {}
    
    # Process in parallel
    for i, chunk in enumerate(json_streamer(file_path)):
        labels = Parallel(n_jobs=n_jobs)(delayed(parallel_dbscan)(segment_data, eps, min_samples) 
                                         for _, segment_data in chunk.groupby('segment'))
        for segment, label in zip(chunk['segment'].unique(), labels):
            if segment not in all_labels:
                all_labels[segment] = []
            all_labels[segment].extend(label)
    
    # Here, you would typically save these labels or process them further for evolution tracking
    # Note: Tracking evolution and HMM prediction might require additional steps for large-scale data

# Usage
file_path = 'collisions1.json'
process_large_json(file_path)