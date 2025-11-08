
# Introduction

![Static Badge](https://img.shields.io/badge/-gray?style=for-the-badge&logo=Go&logoSize=auto)
![Static Badge](https://img.shields.io/badge/Svelte-gray?style=for-the-badge&logo=Svelte&logoSize=auto)
![Static Badge](https://img.shields.io/badge/Vite-gray?style=for-the-badge&logo=Vite&logoSize=auto)
![Static Badge](https://img.shields.io/badge/Tyepscript-gray?style=for-the-badge&logo=Typescript&logoSize=auto)
![Static Badge](https://img.shields.io/badge/CSS-gray?style=for-the-badge&logo=CSS&logoSize=auto)
![Static Badge](https://img.shields.io/badge/HTML5-gray?style=for-the-badge&logo=html5&logoSize=auto)
![Static Badge](https://img.shields.io/badge/SQLITE-gray?style=for-the-badge&logo=Sqlite&logoSize=auto)



URMC-HUB a tool for service desk agents to gather information for making information gathering easier. Converting multiple tools into one web application.

# Uses

- Active Directory Searching (User, Groups, Computer)
- Adding & Removing Groups for users
- Finding Network Drives
- Finding Network Printers
- Custom bookmarks and articles
- Bulk user information
- Bulk Network drive access

# Motivation

Working at the URMC Service Desk, I realized that many of the applications we used were outdated and slow. On top of that, I often needed to open multiple tools just to look up one piece of information. To make things easier, I decided to build a single application that could handle everything, so I didn’t need to keep so many programs open at once.


# What's different from 1.0

I built the first version of the application using Go with plain JavaScript, HTML, and CSS. It worked, but as I added features, the code became hard to maintain. The JS and HTML/CSS structure wasn’t organized well, so even small updates became messy. After seeing what worked and what didn’t, I decided to rebuild the frontend from scratch and clean up the backend so the project would be easier to maintain and improve going forward.

## Frontend moved to Svelte

I really like how Svelte handles components and keeps the CSS scoped to each one. So I rebuilt the entire frontend using Svelte. Now the structure is much cleaner, and making changes is a lot easier because everything is organized and readable.

## Frontend being served from backend

Originally, the frontend was just a static set of files on a shared drive, and it made requests to a backend API. This setup caused performance issues, because the server had to navigate through multiple file paths to serve the HTML files, which slowed everything down. It was also confusing for other agents, since they had to open two separate things for the application to work. To fix this, I embedded the frontend directly into the backend server and now serve the UI locally, which makes everything faster and easier to run.

## Adding a database

I ran into some issues with the database setup. Since each computer was making read and write requests to the same file, SQLite couldn’t handle the concurrent access well. I also didn’t have a dedicated server to host a central database. To work around this, each computer now has its own local SQLite database. When information is needed from another machine, it can read from that machine’s database instead.

## New Features

- New layout and appearance
- Add & Remove & Edit personal bookmarks
- Login done on the web application
- Search page now has filters for the different results
- Bulk search lookup
- Bulk add/remove groups from users
- Actual feedback for add and removal of groups
- MemberOf Groups now showing

# Installation

Instructions on what is needed and how to build the project for development work.

## Required

- GO v1.24.2
- NODE v22.20.0
- NPX 10.9.2


## Build Instructions

### Navigate to backend/server/fronted
```sh
cd backend/server/frontend
```
### Install all npm packages
```sh
npm install
```
### Build the svelte project
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
