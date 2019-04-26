# ForgeRock SaaS Software Engineer Coding Challenge

## Notes

First, thank you for the opportunity.


Due to a time crunch, I didn't have enough time to do this part of the challenge to the best of my ability. There were many external variables in play. 

Since I didn't have a lot of time to do this, I spent a total of about 45 minutes that I did have available to complete the two items in this repo.

First is workert-test.go. This proves that I know how to write a worker type project. No, it's no elegant, and no it's not useful. But it is basic and does what is needed.

Second, I have go-redis.go.  This short script makes a publish subscribe connection to redis and returns the message. Again, is is elegant - HELL NO. But it is something.

To install that which is needed by go-redis, you need to install redis locally.

## Installation Notes:
### Install redis
#### On Mac
 brew install redis

#### From source
$ mkdir redis && cd redis
$ curl -O http://download.redis.io/redis-stable.tar.gz
$ tar xzvf redis-stable.tar.gz
$ cd redis-stable
$ make
$ make test
$ sudo make install

#### Start the redis server
redis-server & 

#### TO stop the redis server 
type fg then ctrl-c the process.

#### Running the two scripts
##### go-redis.gp
First run this: 

go get github.com/garyburd/redigo/redis

go run go-redis.go


This will publish and return: 
    Hello ForgeRock 

##### workers-test.go

This is a simple script with no real prerequiests. 

just run:

go run workers-test.go




I'm sure I may have left something out since my environment has been working locally for sometime.

##### Threat Analysis of this.

It's pretty simple so there's really no threat. There could be some auth and auth added to the redis connection to disallow random users from adding entries to the channel. 

##### Improvements
There is a ton of room for improvements. I would like to create something like that that would allow me to have a chat bot like functionality that stores messages in redis and makes changes to an instance or pod in k8s on the fly.  



----


## Project Background

We have recently decided that all new projects will be done in Typescript for front-end work and Golang for back-end and utility work.  We obviously understand that not every solution will fit into those two buckets, but any outliers will be addressed on a case-by-case basis.  This coding challenge is meant to be non-trivial example of a problem we recently worked on originally in nodejs and then in golang.

Often times we will have tasks that need to be executed (some in parallel and some in sequential order -- in our case allocating cloud resources).  These tasks could take a few seconds or a few minutes and may or may not fail in various ways.

## The Challenge

We are asking that you write a simple worker pattern type project using a modern language. We prefer golang but if you are more comfortable in another language, we are also polyglot and will accept a language you are most proficient with.  These workers can be stubs (i.e. they can just sleep for random periods of time -- bonus points for doing an actual task) but they should feed back into the system in some way -- i.e. tracking and reporting.  We suggest using redis for both the pub/sub worker interactions as well as datastore -- but the technology choice is yours as long as we can build it and it fits a similar pattern.

This challenge is not meant to be a production ready version, but it is meant to show thought process, concepts, and problem solving skills.  You are free to use any 3rd party libraries or code to accomplish the task as long as it is LICENSED accordingly.

## Requirements / Concepts

Some requirements / concepts we would like to see (either in code or at least outlined in the README):

- Build instructions (any 3rd party requirements and how to generally get them setup on either linux or mac) -- docker-compose or something similar is suggested
- Usage instructions (i.e. samples to actually show how it works)
- Unit tests
- Integration tests
- Are there any shortcomings of the code?
- How might this project be scaled?
- How might one approach doing sequential versus parallel tasks?

## Completion Format

- [ ] To do this challenge, please Fork this repository to your own public github repository.
- [ ] Build all the things, and have fun! 
- [ ] Check your code into the fork when you are complete.
- [ ] :shipit:  Send the recruiter at ForgeRock the link to your repository.  

We will review what you did during one of your interview sessions!


