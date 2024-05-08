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
/ $ curl http://hotel:8080/booking/hotels
/ $ curl http://hotel:8080/booking/hotels/ID
/ $ curl -X POST http://hotel:8080/booking/hotels -d '{"name": "New Hotel"}' -H "Content-Type: application/json"
/ $ curl -X PUT http://hotel:8080/booking/hotels/ID -d '{"name": "Updated Hotel Name"}' -H "Content-Type: application/json"
/ $ curl -X DELETE http://hotel:8080/booking/hotels/ID
/ $ curl http://search:8000/search/proba
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
```

Ponisti prethodnu verziju i apply novu 
```shell
kubectl replace --force -f ingress.yml
```