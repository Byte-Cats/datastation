# Running a Project with Docker

## Prerequisites

To run a project with Docker, first ensure you have both docker and docker-compose installed. 
See their [documentation](https://docs.docker.com/compose/install/) for information.

## Clone the Project Repository

Create a directory to store your cloned repositories. Move into the directory, then clone the repository:

```bash
$ git clone [repository URL]
```

Once the repository has been cloned, navigate into the project directory that the clone creates:

```
  cd [project directory]
```

If you have cloned the repository previously, update it prior to installing/re-installing using Docker:

```
git pull
```

Configure the Project

Edit the configuration files located in the project directory as needed.
Build the Container

Note: these steps may take some time to complete depending on the size and complexity of the project.

    Pull the required images:

```
docker-compose pull
```

    Create a directory for sharing resources between your computer and the container:

```
mkdir ~/[project_data]
mkdir ~/[project_data]/share
```

i.e. a directory called [project_data]/share in your home directory

    Build the container:

```
docker-compose build --no-cache
```

    Complete the build process:
        The first time you do this, it will complete the build process, which may include populating required databases or installing dependencies.
        Build time may vary depending on the size and complexity of the project.
        The build process is complete when you see a message indicating that the container has been named.

Run the Container

To start the container, run the following command:
```
docker-compose up -d
```

To stop the container, run the following command:
```
docker-compose down
```

To view the logs for the container:
```
docker-compose logs
```

Other Docker Commands

To list all running containers:
```
docker ps
```

To list all containers, including stopped ones:
```
 docker ps -a
```
To stop a specific container:
```
docker stop [container name or ID]
```
To remove a specific container:
```
docker rm [container name or ID]
```

To remove all containers:
```shell
docker rm $(docker ps -a -q)
```
