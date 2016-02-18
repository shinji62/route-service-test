##Description
Reverse proxy to be use with the Cloudfoundry route service.

Features:
* User proper go Reverse proxy
* Send back to the goRouter a  new header indicating the response from the forwarded URL `X-Response-Forwarding`
* Validate `X-CF-Proxy-Signature`  `X-CF-Forwarded-Url` `X-CF-Proxy-Metadata`

##Creating Route-Service
### Using User provided cf >= 229 and use edge cf cli (unrelease build)

[![asciicast](https://asciinema.org/a/14.png)](https://asciinema.org/a/14)


##Usage

### Build (tested with go 1.5.3)
```
$godep go build
```

###Just push to cloudfoundry
```
$cf push application-name -b go_buildpack -m 128M
```




## TBD / idea
* Add Signal handler and so on. (listenAndService is too simple)
* Embedded service-broker
* More Configuration
* Worker / pool (offload work)