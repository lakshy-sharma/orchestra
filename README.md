# Orchestra
A complete orchestration tool.

## Components Description
1. Gallery: Frontend for the complete ecosystem management.
    1. Communicates with conductor API to provide a GUI for management of cluster.
2. Conductor: Cluster Manager service.
    1. Sets up the security components of the cluster and centralizes all communication to musicians.
    2. Can be setup on multiple nodes securely for high availability.
    3. Acts as a Load balancer for the complete cluster.
    4. Acts as a central repository of tasks to be scheduled.
3. Musician: Node Manager service.
    1. Sets up required services on a node to utilize the resources. (Installs packages and runtimes as required)
    2. Provides a unified api to provide instructions for orchestration of applications. (Called Symphonies)
    3. Manages the orchestration of designated tasks by communication with conductor.
4. Historian: Redis Service
    1. A collection of scripts which can help you setup a redis cluster.
    2. The cluster is useful for conductor application and instead of bundling it we allow you to set it up anywhere.

## Target Services
The goal is to provide orchestration capabilities for the following services.
1. Containers
    1. Docker (Work In Progress)
    2. runc
2. Virtual Machines
    1. Kvm (Work in Progress)
    2. Qemu
3. Cloud Infrastructure
    1. AWS  (Work in Progress)
    2. GCP
    3. Strictly No Azure support 
        * Because no developer should go through the agony of reading Azure API Docs.
4. microVMS
    1. Firecracker

## Vision
To create a simple orchestration platform which can be quickly setup and provide extensive orchestration capabilities for variety of services.