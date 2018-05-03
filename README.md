# umqtt
tiny mqtt client (1.7 Mb gzip)

```
1. go get github.com/knalum/umqtt
2. go build
3. ./umqtt
```

```
Usage of ./mqtt-cli:
  -host string
    	Hostname of the broker
  -int int
    	Interval for publish message in seconds
  -msg string
    	Message (default unix nano)
  -pass string
    	Password for the broker
  -port string
    	Port for the broker (default 1883) (default "1883")
  -topic string
    	Topic (default A) (default "A")
  -user string
    	Username for the broker

```

Example:

```
./mqtt-cli -host=<IP> -int=1 -pass=<PASSWORD> -user=<USERNAME> -topic=<TOPIC>
```

Will send unix nano every 1 second.
