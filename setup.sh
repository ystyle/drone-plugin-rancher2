kubectl config set-cluster default --server=$SERVER
kubectl config set-credentials default --token=$TOKEN
kubectl config set-context default --user=default --cluster=default
kubectl config use-context default
