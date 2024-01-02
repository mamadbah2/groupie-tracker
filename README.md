# Groupie Trackers

Groupie Trackers is a project that involves receiving a given API and manipulating the data contained in it to create a user-friendly website, displaying information about bands and artists, their concert locations, dates, and the relationships between them.

## Project Overview

The API consists of four parts:

1. **Artists:** Contains information about bands and artists, including their name(s), image, inception year, date of their first album, and members.
2. **Locations:** Contains information about the last and/or upcoming concert locations.
3. **Dates:** Contains information about the last and/or upcoming concert dates.
4. **Relation:** Establishes the link between artists, dates, and locations.

## Features

- Display bands' information through various data visualizations such as blocks, cards, tables, lists, pages, graphics, etc.
- Creation of events/actions triggered by the client, time, or any other factor.
- Client-server communication for handling features like client calls to the server.

## Backend

- Written in Go.
- Utilizes only standard Go packages.
- Ensures the site and server do not crash at any time.
- Implements error handling for all pages.
- Follows good coding practices.
- Includes test files for unit testing.

## Usage

1. Clone the repository: `git clonehttps://learn.zone01dakar.sn/git/mamadbah/groupie-tracker.git`
2. Navigate to the project directory: `cd groupie-trackers`
3. Run the backend server: `go run ./cmd/web/.`
4. Access the website in your browser: `http://0.0.0.0:1111`

## Example API

An example of a RESTful API can be found [here](https://groupietrackers.herokuapp.com/api).

