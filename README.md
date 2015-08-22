[![Build Status](https://travis-ci.org/halk/in-common.svg?branch=master)](https://travis-ci.org/halk/in-common)
[![GitHub license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/halk/in-common/blob/master/LICENSE)

# In Common: data-agnostic, graph-based collaborative recommendation engine

In Common is a collaborative recommendation engine which utilizes the graph database system [Neo4j](http://neo4j.com/). It is data-agnostic in the sense that the engine has no knowledge about the data model and requires no data migration or set up. With its minimal RESTful API, it is almost a drop-in solution.

Since this is part of my ongoing [MSc project](https://github.com/halk/msc-project-report), README will be improved by October.

## Concept

### Data Model

The main model behind this engine is an event of the type **Relationship** which has the following data fields:

- a **subject** type (e.g. *user*)
- a **subjectID** (e.g. *123*)
- a **relationship** (e.g. *viewed*)
- a **object** type (e.g. *product*)
- a **objectID** (e.g. *SKU001*)

Based on this model two graph nodes with the appropriate label and IDs are created (if they do not exist already) and a relationship with the appropriate label added (or a counter increased if there is already one).

### Recommendation

The recommendation are computed based on **RecommendationRequest** which has the following data fields:

- a **subject** type (e.g. *user*)
- a **subjectID** (e.g. *123*)
- a **relationship** (e.g. *viewed*)
- a **object** type (e.g. *product*)

Notice that the **objectID** is missing compared to a **Relationship** model.

The following query is then performed:

> Return objects of type **object** which subjects of type **subject**, who have **relationship** the same objects of type **object** as the subject of type **subject** and ID **subjectID** has **relationship**, have **relationship** which **subjectID** has not yet.

## API

Events are managed by POST and DELETE requests. Recommendations are fetched via GET.

**POST /engine** Adds an event. Request body should have a JSON representation of the **Relationship** model.

**DELETE /engine** Removes an event. Request body should have a JSON representation of the **Relationship** model.

**GET /recommendation** Gets recommendations. The fields of the **RecommendationRequest** model are passed as a query string.

### Installation

[Go](https://golang.org/) 1.4 and [Neo4j](http://neo4j.com/) are required.

```bash
$ git clone https://github.com/halk/in-common
$ cd in-common
$ go get github.com/mattn/gom
$ gom install
$ gom build
$ inCommon
# will be looked up in $GOPATH
```

Please see [recowise-vagrant](https://github.com/halk/recowise-vagrant) for provisioning details. You can use [supervisord](http://supervisord.org/) to daemonize it.

## Tests

```bash
$ cd in-common
$ gom test ./...
```

## Documentation

Documentation can be found on [GoDoc](https://godoc.org/github.com/halk/in-common).
