![img](https://imgs.search.brave.com/5aouW_db_X1s2jLw-XcjUDSTpv9aKjZmYQLXUU7lTCU/rs:fit:860:0:0/g:ce/aHR0cHM6Ly9pLnBp/bmltZy5jb20vb3Jp/Z2luYWxzLzI2L2Vk/LzJlLzI2ZWQyZTYx/MTA2MGMwNzZhZDFk/MTY5MTVlZGJiYzgy/LmpwZw)

# Employee Management System

The Employee Management System is a simple application designed to
manage employee information within an organisation. It has four employee
positions - admin, manager, developer and tester. It allows you to
perform essential Admin tasks including adding new
employees, updating their details, and viewing employee records. It uses
firestore as database. This README provides an
overview of the project and how to get started.

# Features

- **_Employee Records:_** Store and manage employee information, including
  names, IDs, contact details, and positions, etc.
- **_Add and Edit:_** Easily
  add new employees and edit existing records.
- **_Search and Filter:_**
  Quickly find employees using search and filter options.

### List of operations admins can perform

1. Add employee details.
2. View all employees.
3. Update employee
   details.
4. Delete employee details.
5. Search employees by their first name, last name, email and role (with one or multiple fields).
6. Find employee by employee id.

### Assumptions and Notes

- Google Firestore is used as a database.
- When a new employee is
  added, his/her employee id will be the maximum employee id presentin the database + 1. It will be generated automatically upon addition of each
  employee.

## Getting Started

- Clone this repository into your machine.
- navigate to dist folder
- run index.html on server to render swagger ui
