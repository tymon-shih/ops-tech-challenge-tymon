kubectl create secret generic ops-tech-challenge-secret --from-literal=SECRET=yoursharedsecret

kubectl apply -f ops-tech-challenge-deployment.yaml