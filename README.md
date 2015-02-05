# Monitoring Apache Mesos Resources
The adoption of container based solutions is increasing dramatically among IT companies. Reports recently presented by Docker shows this adoption in terms of numbers (http://blog.docker.com/2015/01/docker-project-2014-a-whirlwind-year-in-review/).
But as more containers are being executed by a company, more server hosts are needed. Thereby, a new challenge arises, which is manage the containers deployments and its hosts computational resources (e.g., cpus, memory, storage). For such a scenario, the project Apache Mesos has a solution. Mesos magane all server hosts (slaves) computational resources and delivers these resource as a big cluster. This cluster computational resources are the sum of all slaves computational resources, so Mesos gives an unified view of the resources. While there are available resources, it's possible to deploy a new container.

However, Mesos provides only a view of the resource usage current status and does not keep its history. Mesos users are unable to understand the datacenter behavior when, as example, a Mesos node is not working. Or the period of the day/week/month when the computational resources are been heavily consumed. In this context, if a snapshot of the current view were stored in an given rate, such as one snapshot per minute, it would be possible to better understand the cluster behavior. And more! What if these informations were stored at Elasticsearch? It would be possible to build beautiful resource graphs and to do some data analysis.

This project is about motoring Apache Mesos cluster using Elasticsearch. But before continue, let's introduce projects Docker, Mesos and Elasticsearch.

##What is Docker?

##What is Mesos?

##What is Elasticsearch?
