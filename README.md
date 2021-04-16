## Craft-demo
A `cli` and api `server` that returns country base on country code. 



## Build 
```bash
go build -o lookup *go
```

## Usage
```bash
$ ./lookup -h
Usage:
./lookup --countryCode=AU
To start a server do ./lookup -s
```

## Deploying to kubernetes
After building the image, we can push it to ecr repo. Then use helm to deploy it to eks cluster via nodeport, use terraform to provision nlb/alb that listens to the nodeport to expose it to internet. 


## Logging and monitoring 
By default when you run it as a `server` there will be `/metrics` where prom can scrap it off of. Alerts/graphs can be made base on that


## Stateless design
The `data.json` can be stored in s3, which will get periodically updated by a labda functino. This service should be hosted in a pod where it has permission to the s3 bucket. [Doc](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html)
