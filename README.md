# zero-strange
![image](https://user-images.githubusercontent.com/112534208/191252324-33debe57-887a-4eb1-8123-3e83edaaec2d.png)

**REQUIREMENTS**: etcd, redis-server running

Redis is required cause of database model **embed** cache as default behaviour. 
You can see in generated usermodel_gen.go.

````
	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}
````


There are two services. Authentication and Broker.

Both services use mysql for user data.

broker use mongo for writing log.

**Broker Service**
Broker accepts rpc and api requests.

**API**

get **/v2/users**

post **/v2/users/**

get **/v2/users/:id**


**RPC**

write log of Login from authentication service with method **"logLogin"**


**Authentication Service**

Authentication accepts api request.

**api**

post **/user/login**

When the **/user/login** enpoint hits, it calls **"logLogin"** rpc and write log in mongodb by broker service.

If you want to write cutom database functions, u can add model/usermodel.go

Example:

Add **FindAll** method in UserModel interface

````
	UserModel interface {
		userModel
		FindAll(ctx context.Context) ([]*User, error)
	}
````

Implement **FindAll** funtionin the same file.

````
func (m *customUserModel) FindAll(ctx context.Context) ([]*User, error) {
	var resp = []*User{}

	err := m.QueryRowsNoCacheCtx(ctx, &resp, "SELECT * FROM `user`;")
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
````
