### Description
Minimalistic web application written in Go that acts as an interface for Redis. Uses Nginx as a reverse proxy. Redis and application communicate using mutual TLS.
### Usage example
```
git clone https://github.com/ScarletBlizzard/Go_Redis_Nginx.git
cd Go_Redis_Nginx
```
Set Redis password:
```
echo "REDIS_PASSWORD=<password for redis>" > .env
```
Create TLS keys and certificates, move them into Docker build contexts:
```
openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 \
  -out app.crt -keyout app.key \
  -subj "/CN=app/" -addext "subjectAltName = DNS:app"
openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 \
  -out redis.crt -keyout redis.key \
  -subj "/CN=redis/" -addext "subjectAltName = DNS:redis"
cp *.crt app/
mv app.key app/
mv *.crt redis.key redis/
```
Build and run:
```
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
