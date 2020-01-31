echo 'TODO may take up to 5 minutes'

hostnamectl set-hostname demo
rm -f original-ks.cfg

curl -Lo /usr/bin/kind "https://github.com/kubernetes-sigs/kind/releases/download/v0.7.0/kind-$(uname)-amd64"
chmod +x /usr/bin/kind

curl -Lo /usr/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/v1.17.0/bin/linux/amd64/kubectl
chmod +x /usr/bin/kubectl

echo '
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
' > kind.yaml

kind create cluster --config kind.yaml

for node in $(docker ps | awk '{if(NR>1) print $1}'); do
    docker exec $node sh -c 'apt-get update && apt-get -y install systemd network-manager && touch /etc/NetworkManager/conf.d/10-globally-managed-devices.conf && systemctl start network-manager && ip link add eth1 type veth peer name eth2' &
    waitPids[${i}]=$!
done

for pid in ${waitPids[*]}; do
    wait $pid
done

kubectl apply -f https://raw.githubusercontent.com/kubevirt/cluster-network-addons-operator/master/manifests/cluster-network-addons/0.25.0/namespace.yaml
kubectl apply -f https://raw.githubusercontent.com/kubevirt/cluster-network-addons-operator/master/manifests/cluster-network-addons/0.25.0/network-addons-config.crd.yaml
kubectl apply -f https://raw.githubusercontent.com/kubevirt/cluster-network-addons-operator/master/manifests/cluster-network-addons/0.25.0/operator.yaml
cat <<EOF | kubectl create -f -
apiVersion: networkaddonsoperator.network.kubevirt.io/v1alpha1
kind: NetworkAddonsConfig
metadata:
  name: cluster
spec:
  nmstate: {}
EOF
kubectl wait networkaddonsconfig cluster --for condition=Available

kubectl get nns
kubectl get nncp
kubectl get nnce

bash
