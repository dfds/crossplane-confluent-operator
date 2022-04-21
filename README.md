# Crossplane Confluent operator

This operator makes sure prerequisites for accessing Confluent Kafka topics are in place.

Prerequisite set:
* Confluent service-account
* Confluent ACL's
* Confluent API-keys

The operator will detect topic manifests deployed per environment and cluster and provision/reconcile the above set leaving an secret in the end users namespace for topic interactions.

## Development

### Requirement

* A Kubernetes cluster
* Helm
* Crossplane install on the cluster
* Provider Confluent installed on the cluster

``` bash
# Setup Kubernetes cluster (I use minikube, but use whatevaaaahhhhhhh you like)
minikube start

# Create crossplane-system namespace
kubectl create namespace crossplane-system

# Add and install crossplane helm chart into crossplane-system namespace
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update
helm install crossplane --namespace crossplane-system crossplane-stable/crossplane

# Install provider confluent (make sure that the version match the one in go.mod)
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-confluent
spec:
  package: "dfdsdk/provider-confluent:v0.0.4"
EOF

# Run the operator locally
make run

# Apply a topic manifest
cat <<EOF | kubectl apply -f -
apiVersion: kafka.confluent.crossplane.io/v1alpha1
kind: Topic
metadata:
  name: confluent-test
spec:
  forProvider:
    cluster: abc-00000
    environment: env-00000
    topic:
      name: confluent-test
      partitions: 1
      config:
        retention: 259200000
  providerConfigRef:
    name: confluent-provider
EOF
```