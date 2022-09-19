# The Space Launcher
There is a list of available launchpads and destinations: Mars, Moon, Pluto, Asteroid Belt, Europa, Titan, Ganymede.
On every day of the week from the same launchpad has to be a “flight” to a different place.
Information about available launchpads and upcoming SpaceX launches you can find by [SpaceX API](https://github.com/r-spacex/SpaceX-API)
Launchpads  are shared we cannot launch rockets from the same place on the same day.

The service provides an API to:
- Create booking (POST /space_launcher/bookings)
- Get all bookings (GET /space_launcher/bookings)
- Delete booking (DELETE /space_launcher/bookings)

Endpoint to book a ticket where client sends data like:
- First Name
- Last Name
- Gender
- Birthday
- Launchpad ID
- Destination
- Launch Date

## TODO:
- add planed time validation (not to plan in the past)
- add destination table in DB
- add swagger contracts
- add Makefile
- add linters