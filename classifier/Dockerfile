# Use an official Python runtime as a parent image
FROM tensorflow/tensorflow:latest

# Set the working directory in the container to /app
WORKDIR /app

# Add the current directory contents into the container at /app
ADD . /app

# Install any needed packages specified in requirements.txt
RUN pip3 install -r requirements.txt

# Make port 50051 available to the world outside this container
EXPOSE 50051

# Run server.py when the container launches
CMD ["python", "serve.py"]