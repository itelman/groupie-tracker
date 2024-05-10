## groupie-tracker

## FOR AUDITORS: IMPORTANT!!!

WHEN SUBMITTING A REQUEST TO "/artists", IF A VALUE OF ID IS "+25", "-25" OR "00025", THE WEBSITE WILL RETURN ERROR 404.

## Authors:

Ilyas Telman (@itelman), Arshat Aitkozha (@araitkozha).

## Objectives

Groupie Trackers consists of receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.

- It will be given an [API](https://groupietrackers.herokuapp.com/api), that consists of four parts:

  - The first one, `artists`, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

  - The second one, `locations`, consists in their last and/or upcoming concert locations.

  - The third one, `dates`, consists in their last and/or upcoming concert dates.

  - And the last one, `relation`, does the link between all the other parts, `artists`, `dates` and `locations`.

- Given all this we built a user friendly website where we displayed the bands info through several data visualizations (examples : blocks, cards, tables, list, pages, graphics, etc).

- This project also focuses on the creation of events/actions and on their visualization.

  - The event/action we worked with is known as a client call to the server (client-server). It is a feature that needs to trigger an action. This action communicates with the server in order to recieve information, ([request-response])(https://en.wikipedia.org/wiki/Request%E2%80%93response)
  - An event consists in a system that responds to some kind of action triggered by the client, time, or any other factor.

### Geolocalization

- Groupie Tracker Geolocalization consists of mapping the different concerts locations of a certain artist/band given by the Client.

- It applies a process of converting addresses (ex: Germany Mainz) into geographic coordinates (ex: 49,59380 8,15052), which then uses to place markers for the concerts locations of a certain artist/band on a map.

- For this feature, [Google Maps JavaScript API](https://developers.google.com/maps/documentation/javascript/) was used.

This project focuses on:

- Manipulation and storage of data.
- [JSON](https://www.json.org/json-en.html) files and format.
- HTML.
- Manipulation of Maps API
- Geolocation, geocoding, etc
- [Event](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Building_blocks/Events) creation and display.
- [Client-server](https://developer.mozilla.org/en-US/docs/Learn/Server-side/First_steps/Client-Server_overview).

## Usage: how to run

- Download the repository to your local machine.
- Open the repository via Terminal.
- Run the following command:
```console
go run main.go
```
- Run the server on:
```console
http://localhost:8080/
```
