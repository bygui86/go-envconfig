
# Go env-var sample project

## Intructions

### viper

1. Run

	```
	go run viper/main.go
	```

2. Results should be like following

	```
	Config file loaded
	config1: value1
	```

### gonfig

`TODO`

### envconfig

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
	go run envconfig/main.go
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

### Articles

* https://medium.com/@felipedutratine/best-practices-for-configuration-file-in-your-code-2d6add3f4b86

### Techs

* https://github.com/spf13/viper
* https://github.com/tkanos/gonfig
* https://github.com/kelseyhightower/envconfig
