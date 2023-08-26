### Description
Minimalistic web application written in Go that acts as an interface for Redis. Uses Nginx as a reverse proxy.
### Usage example
Install and run:
```
git clone https://github.com/ScarletBlizzard/Go_Redis_Nginx.git
cd Go_Redis_Nginx
docker compose up
```
Set key-value pair:
```
curl -X POST -H "Content-Type: application/json" -d '{"<key>":"<value>"}' localhost/set_key
```
Get value by key:
```
curl localhost/get_key?key=<key>
```
Delete key-value pair:
```
curl -X POST -H "Content-Type: application/json" -d '{"key":"<key>"}' localhost/del_key
```
