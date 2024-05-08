# Example 

Turn on ingress 

```shell
minikube addons enable ingress
```

Create and delete namespace:

```bash
kubectl create namespace backend
kubectl delete namespace backend
```
Sve kubernetes fajlove pokrenuti da dobijemo configmap, secret i mongo service i statefulSet
```shell
kubectl -n backend apply -f mongo-configmap.yml
kubectl -n backend apply -f mongo-secret.yml
kubectl -n backend apply -f search-configmap.yml 
kubectl -n backend apply -f mongo.yml
kubectl apply -f search-service.yml 
```
Get pods:
```shell
kubectl -n backend get pods
```

Testing load balancing and service:
```shell
kubectl -n backend run -it --rm  --image curlimages/curl:8.00.1 curl -- sh
```
Inside the container execute `curl http://hotel:8083/hotels` (hotel jer je to naziv servisa)
```shell
 curl http://hotel/hotel/hotels
 curl http://hotel/hotel/hotels/ID
 curl -X POST http://hotel/hotel/hotels -d '{"name": "New Hotel"}' -H "Content-Type: application/json"
 curl -X PUT http://hotel/hotel/hotels/ID -d '{"name": "Updated Hotel Name"}' -H "Content-Type: application/json"
 curl -X DELETE http://hotel/hotel/hotels/ID
 curl -H "Authorization: Bearer TOKEN” "http://hotel/booking/hotels"
```

Get JWT token for user
```shell
 curl -X POST -d "client_id=Istio" -d "username=hotel-user" -d "password=test" -d "grant_type=password" "http://keycloak.default.svc.cluster.local:8080/realms/Istio/protocol/openid-connect/token"
```

Get JWT token for admin
```shell
 curl -X POST -d "client_id=Istio" -d "username=hotel-admin" -d "password=test" -d "grant_type=password" "http://keycloak.default.svc.cluster.local:8080/realms/Istio/protocol/openid-connect/token"
```

Ingress setup:
Deploy ingress:
```shell
kubectl -n backend apply -f ingress.yml
kubectl -n backend describe ingress demo-ingress
```

Apply za ceo ili vise direktorijuma
```shell
kubectl -n backend apply -R -f k8s
kubectl -n backend apply -R -f istio
```

Ponisti prethodnu verziju i apply novu 
```shell
kubectl replace --force -f ingress.yml
```

Keycloak
```shell
minikube addons enable ingress
kubectl create -f https://raw.githubusercontent.com/keycloak/keycloak-quickstarts/latest/kubernetes/keycloak.yaml
minikube tunnel

browser: localhost:8080 (username: admin, password: admin)

Create Istio realm
Create Istio client 
Create hotel-user , hotel-admin (password: test)
```