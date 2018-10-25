set -x

kubectl create secret generic hmac-token --from-file=hmac=hmac.txt
kubectl create secret generic oauth-token --from-file=oauth=oauth.txt

kubectl apply -f https://raw.githubusercontent.com/kubernetes/test-infra/master/prow/cluster/starter.yaml
