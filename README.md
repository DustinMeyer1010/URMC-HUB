
# Introduction

![Static Badge](https://img.shields.io/badge/-gray?style=for-the-badge&logo=Go&logoSize=auto)
![Static Badge](https://img.shields.io/badge/React-gray?style=for-the-badge&logo=React&logoSize=auto)
![Static Badge](https://img.shields.io/badge/Vite-gray?style=for-the-badge&logo=Vite&logoSize=auto)
![Static Badge](https://img.shields.io/badge/Tyepscript-gray?style=for-the-badge&logo=Typescript&logoSize=auto)
![Static Badge](https://img.shields.io/badge/CSS-gray?style=for-the-badge&logo=CSS&logoSize=auto)
![Static Badge](https://img.shields.io/badge/HTML5-gray?style=for-the-badge&logo=html5&logoSize=auto)
![Static Badge](https://img.shields.io/badge/SQLITE-gray?style=for-the-badge&logo=Sqlite&logoSize=auto)



The URMC-HUB is a simple web applications that converts tools for the URMC service desk into one fast application.

- Active Directory Searching
- Network Drive Searching
- Network Printer Searching
- Bookmarks for common links
- System summary index for common KBs

# Motivation

URMC Service desk agents are requred to use many application that are slow, require other information before getting what is needed and require a ton of application open to get all information. Since most of our work flow in done in the browser. A better solution to this was building an application that was faster, requiring less clicks and can be used in one spot.


# What's different from 1.0

Since the release of URMC-HUB 1.0. There are many things that can be changed to better manage the project and for better usage.

## Frontend moved to React

The application was hard to manage with all the files being static html, css, and js files. The project will be in react to better manage the files.

## Frontend being served from backend

Orginally the files were sitting on the share drive and to open the project you would just click on an .html file. New version will have the files embedding in the final .exe. Leaving the project to just one file rather than a bunch of files.

## Adding a database

There will now be a sqlite database for storing information about each service desk agent for them to customize parts inside of the tool. Like having custom links on the bookmarks page.

## New Features

- New layout and appearnce for better readability
- Ability to add, remove, and update links
- Login is now done on website itself
- Adding a queue for searches to cycle through for look for information

# Installation

Instructions on what is needed and how to build the project for development work.

## Required

- GO v1.24.2
- NODE v23.11.0
- NPM 10.9.2


## Build Instructions

### Navigate to backend/server/fronted
```sh
cd backend/server/frontend
```
### Install all npm packages
```sh
npm install
```
### Build the react project
```sh
npm run build
```
> Note: npm run dev to just work with frontend
### Navigate to backend/cmd/hub
```sh
cd backend/cmd/hub
```
### Install go packages
```sh
go mod tidy
```
### Run main.go or build project
```sh
go run .
```
> Project runs on localhost:8080/
```sh
go build .
```
> Run the hub.exe created
