# space_launcher

## The Challenge
Imagine it’s 2049 and you are working for a company called SpaceTrouble that sends people to different places in our solar system. 
You are not the only one working in this industry. Your biggest competitor is a less known company called SpaceX. 
Unfortunately you both share the same launchpads and you cannot launch your rockets from the same place on the same day. 
There is a list of available launchpads and your spaceships go to places like: Mars, Moon, Pluto, Asteroid Belt, Europa, Titan, Ganymede. 
Every day you change the destination for all the launchpads. Basically on every day of the week from the same launchpad has to be a “flight” to a different place.

Information about available launchpads and upcoming SpaceX launches you can find by SpaceX API: https://github.com/r-spacex/SpaceX-API

Your task is to create an API that will let your consumers book tickets online.

In order to do that you have to create 2 endpoints:

Endpoint to book a ticket where client sends data like:
- First Name 
- Last Name 
- Gender 
- Birthday
- Launchpad ID
- Destination ID
- Launch Date

You have to verify if the requested trip is possible on the day from provided launchpad ID and do not overlap with SpaceX launches, if that’s the case then your flight is cancelled.

Endpoint to get all created Bookings.

## Extra points:
When you use docker/docker-compose to run the project.
When you write unit/functional tests.
When you create an endpoint to delete booking.

## Technical requirements:
Please use Golang and Postgres.
Please use github.
Commit your changes often. Do not push the whole project in one commit.

## TODO:
- add validation on used Destination per launchpad per week
- add planed time validation (not to plan in the past)
- add test coverage
- improve logging
- make the naming conventional & consistent
- revise the package structure
- add destination table in DB
- add booking status
- add delete endpoint
- add swagger contracts
- add documentation
- add Makefile
- add linters