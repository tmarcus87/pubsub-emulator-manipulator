pubsub-emulator-manipulator
===========================

Manipulator for cloud pubsub emulator

## Usage
```
# Setup
$ docker run --rm -p '8085:8085' -d --name pubsub knarz/pubsub-emulator:latest
$ export PUBSUB_EMULATOR_HOST=localhost:8085


# Checkout
$ git clone git@ghe.ca-tools.org:ono-takahiko/pubsub-emulator-manipulator.git
$ cd pubsub-emulator-manipulator
$ ln -s pem-${YOUROS}-${YOURARCH} bin/pem

# List topics
$ bin/pem topics list -project localproject-123

# Create topics
$ bin/pem topics create -project localproject-123 -topic mytopic

# Create subscription for topics
$ bin/pem topics subscription -project localproject-123 -topic mytopic -subscription mytopic-sub

# Publish message
$ bin/pem topics publish -project localproject-123 -topic mytopic \
    -data 'Hello, world!' -attribute 'key1=value1,key2=value2'

# List subscription
$ bin/pem subscriptions list -project localproject-123

$ bin/pem subscriptions consume -project localproject-123 -subscription mytopic-sub
```