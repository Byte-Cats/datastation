# Running Datastation with Docker

## Prerequisites

To install Datastation via Docker, first ensure you have both docker and docker-compose installed. 
See their [documentation](https://docs.docker.com/compose/install/) for information.


## Clone the Datastation Repository
Create a directory collate your cloned repositories. Move into the directory then, clone the repository. 

```bash
$ git clone https://github.com/ByteCats/Datastation
```

Once the repository has been cloned, cd into the Datastation directory that the clone creates.

```bash
$ cd Datastation/
``` 

If you have cloned the repository previously, update it prior to installing/re-installing using Docker

```bash
$ git pull
```

## Configuring the software

Edit the env files located in `datastation/deploy/{database_name}`


## Build the container

*Note: some of these steps take >>1hr to complete depending on the speed of your internet connection*

- Pull images

```bash
$ docker-compose pull
```

- Create a directory for sharing resources between your computer and the container
```bash
$ mkdir ~/datastation_data
$ mkdir ~/datastation_data/share
```
*i.e.* a directory called `datastation_data/share` in your `home` directory

- Build

```bash
$ docker-compose build --no-cache
```

- Complete build
    - The first time you do this, it will complete the build process, for example, populating the required the databases
    - The build takes a while because the  vv databases are large. However, this is a significant improvement on previou
    s versions. Build time is ~30 minutes (depending on the speed of you computer and internet connection)
    - The build has completed when you see the message ***"naming to docker.io/library Datastation_restvv"***

