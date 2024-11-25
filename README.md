# cyber-forge-xds-balancer

This repository provides a hands-on demonstration of setting up an xDS (eXtensible Discovery Service) server tailored for use with the Envoy Proxy. It showcases a basic yet functional configuration to help you get started with service discovery and dynamic configuration.

## Sample Config File

This a sample YAML file that defines the xDS configuration. This includes listener setup, routing logic, and cluster endpoint specifications.

```yaml
name: sample-config
spec: 
  listeners:
  - name: listener_0
    address: 0.0.0.0
    port: 9000
    routes:
    - name: echo-route
      prefix: /
      clusters:
      - echo-cluster
  clusters:
  - name: echo-cluster
    endpoints:
    - address: 127.0.0.1
      port: 9101
    - address: 127.0.0.1
      port: 9102
```
This configuration initializes a listener on port 9000, routes requests to an endpoint cluster named echo-cluster, and specifies two backend services on ports 9101 and 9102.

## Cleaning Up

Stop all the sample endpoints created in the previous step:
```
docker stop echo-9100
docker stop echo-9101
```
