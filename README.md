# Monitoring Apache Mesos Resources
The adoption of container based solutions is increasing dramatically among IT companies. Reports recently presented by Docker shows this adoption in terms of [numbers](http://blog.docker.com/2015/01/docker-project-2014-a-whirlwind-year-in-review/).
But as more containers are being executed by a company, more server hosts are needed. Thereby, a new challenge arises, which is manage the containers deployments and its hosts computational resources (e.g., cpus, memory, storage). For such a scenario, the project Apache Mesos has a solution. Mesos magane all server hosts (slaves) computational resources and delivers these resource as a big cluster. This cluster computational resources are the sum of all slaves computational resources, so Mesos gives an unified view of the resources. While there are available resources, it's possible to deploy a new container.

However, Mesos provides only a view of the resource usage current status and does not keep its history. Mesos users are unable to understand the datacenter behavior when, as example, a Mesos node is not working. Or the period of the day/week/month when the computational resources are been heavily consumed. In this context, if a snapshot of the current view were stored in an given rate, such as one snapshot per minute, it would be possible to better understand the cluster behavior. And more! What if these informations were stored at Elasticsearch? It would be possible to build beautiful resource time series graphs and to do some data analysis.

This tutorial is about motoring Apache Mesos cluster using Elasticsearch. But before continue, let's introduce projects Docker, Mesos and Elasticsearch.

##What is [Docker](https://www.docker.com/whatisdocker/)?
Docker is an open platform for DevOps to build, ship and run distributed applications in a form of lightweight containers that can be deployed in the Docker Engine. As a result, it's possible to provide a standardized environment in such a way, the exactly same code that runs in production, can also run in developers laptop.

##What is [Mesos](http://mesos.apache.org/)?
Apache mesos abstracts computational resources (i.e. cpus, memory, disk) away from the hosts (physical or virtual) and delivers a cluster to run elastic distributed systems, for example, a containerized application.

##What is [Elasticsearch](http://www.elasticsearch.org/overview/elasticsearch)?
Elasticsearch is a highly scalable open source real-time search and analytics engine. Document oriented, Elasticsearch stores data as structured JSON documents, whose all fields are indexed and can be used in a single query.
