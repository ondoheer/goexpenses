# Expenses backend API

The main goal of this project is to build a small backend in Go for my defunct expenses APP

## Models

- User
- Category
- Expense
- Month ?

I will first try to implement it without the Month model by just using SQL date
queries.

### User Model

* id
* username
* password
* email
* name

#### Relationships:

* expenses
* categories


### Category Model

* id
* name
* label

#### Relationships:
* user_id

### Expense Model

* id
* name
* date
* amount

#### Relationships:
* user_id
* category_id


## Routes

Endpoint | method | action
---------|--------|---------|
/login | GET | shows login form
/login | POST | logs the user in
/logout | GET | logs the user out
/   | GET | shows lates expenses, main menu, add expense button
/expense | GET | list of all expenses, search filter, latest expense, form to add expenses
/expense | POST | creates a new expense
/expense/{id} | GET | edit form for the expense id={id}
/expense/{id} | PUT | update expense with id={id}
/expense/{id} | DELETE | delete expense with id={id}
/category | GET | list of categories and form to add new one
/category   | POST | create a new category
/category/{id} | GET | edit form for the category id={id}
/category/{id} | PUT | update category with id={id}
/category/{id} | DELETE | delete category with id={id}

