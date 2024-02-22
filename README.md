# Zitch

#### Confirmed
1. Run docker compose up.
2. Wait like 5-6 minutes because it takes a longgg time to initially download everthing
3. Check docker ps. You should have two services running
4. Confirm healthcheck by running `curl 0.0.0.0:4000/v1/healthcheck`
5. Copy the file into the docker go container with `docker cp ./classifier/image.jpg [container-id]:/app/image.jpg`

#### Kinda untested
1. `curl -X GET 0.0.0.0:4000/v1/classifier -H 'Content-Type: application/json' -d '{"imgurl": "image.jpg"}'`
2. Run the inference part with the above command, this should theoretically work but this crashed my docker so I'm guessing my laptop is not meant to run classifiers. Or maybe the command is wrong. This is why it's untested. 