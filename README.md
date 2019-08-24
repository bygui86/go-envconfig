
# Go envconfig sample project

## Intructions

1. Export some environment variables

	```
	export MYAPP_DEBUG=false
	export MYAPP_PORT=8080
	export MYAPP_USER=Kelsey
	export MYAPP_RATE="0.5"
	export MYAPP_TIMEOUT="3m"
	export MYAPP_USERS="rob,ken,robert"
	export MYAPP_COLORCODES="red:1,green:2,blue:3"
	```

2. Run

	```
	go run main.go
	```

3. Results should be like following

	```
	Debug: false
	Port: 8080
	User: Kelsey
	Rate: 0.500000
	Timeout: 3m0s
	Users:
	  rob
	  ken
	  robert
	Color codes:
	  red: 1
	  green: 2
	  blue: 3
	```

---

## Links

* https://github.com/kelseyhightower/envconfig
