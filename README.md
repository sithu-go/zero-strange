# zero-strange
![image](https://user-images.githubusercontent.com/112534208/191252324-33debe57-887a-4eb1-8123-3e83edaaec2d.png)

There are two services. Authentication and Broker.

Both services use mysql for user data.
broker use mongo for writing log.

Broker Service
Broker accepts rpc and api requests.
api
get **/v2/users**
post **/v2/users/**
get **/v2/users/:id**

rpc
write log of Login from authentication service with method **"logLogin"**


Authentication Service

Authentication accepts api request.
api
post **/user/login**

When the **/user/login** enpoint hits, it calls **"logLogin"** rpc and write log in mongodb by broker service.
