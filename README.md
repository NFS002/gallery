# Gallery

### Gallery is a simple web app that builds a network based on 3 archetypes:

- Hosts. Anyone can sign up to be a host, the host role entails that you are willing to be contacted 
by artists who have signed up fot the platform and would potentially agree to host and display 
some artists work whom you have met and spoken with to the public on arranged times/dates.

- Artists are pretty much what you would expect. For example, someone who has created a small collection of
pottery can sign up to this platform, contact hosts in their area and hopefully arrange a home exhibition.00  

- Viewers, anyone with an interest in appreciating art can browse current exhibitions and their location on gallery, 
and then email the host via gallery to check if they can come for a viewing.


### A containerised **MVP** is deployed on heroku [here](https://heroku-gallery-development.herokuapp.com)


## TechStack
- PostgreSQL
- [Gobuffalo framework for golang](https://gobuffalo.io)
  - [pop: sort-of-orm database tool for gobuffalo](https://github.com/gobuffalo/pop)
  - [plush: html/golang temlplates for gobuffalo](https://github.com/gobuffalo/plush)
- [Heroku PaaS](https://www.heroku.com/platform)

## Running locally

1. [Install buffalo](https://gobuffalo.io/documentation/getting_started/installation/)

2. Start the development server
```shell
buffalo dev
```
