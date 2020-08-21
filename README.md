# amazingChat
test application for internship

# to build application:
`sudo docker build -t alex https://github.com/saskamegaprogrammist/amazingChat.git`

# to run application:
` sudo docker run -p 9000:9000 --name alex -t alex`

# API

# Add user
"/users/add" **POST**

### Answers

- 201 - Created user
- 409 - Already registered
- 500 - Internal error

## CURL request example

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username": "user_1"}' \
  http://localhost:9000/users/add

# Add chat
"/chats/add" **POST**

### Answers

- 201 - Created chat
- 400 - One of the users doesn't exist
- 409 - Already exists
- 500 - Internal error

## CURL request example

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "chat_1", "users": ["<USER_ID_1>", "<USER_ID_2>"]}' \
  http://localhost:9000/chats/add

# Add message
"/messages/add" **POST**

### Answers

- 201 - Created message
- 400 - Chat or user doesn't exist
- 500 - Internal error

## CURL request example

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": "<CHAT_ID>", "author": "<USER_ID>", "text": "hi"}' \
  http://localhost:9000/messages/add

# Get user's chats sorted by last created message
"/chats/get" **POST**

### Answers

- 200 - OK
- 400 - User doesn't exist
- 500 - Internal error

## CURL request example

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"user": "<USER_ID>"}' \
  http://localhost:9000/chats/get

# Get chat's messages sorted by last created 
"/messages/get" **POST**

### Answers

- 200 - OK
- 400 - Chat doesn't exist
- 500 - Internal error

## CURL request example

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": "<CHAT_ID>"}' \
  http://localhost:9000/messages/get