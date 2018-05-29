# cni

## Requirement

IPAM `static` binary. see [cni plugin](https://github.com/containernetworking/plugins/tree/master/plugins/ipam/static)

## Underlay OVS with static IPAM

Prepare a json file named `static.conf`

```json
{
        "name": "mynet",
        "type": "bridge",
        "ipam": {
                "type": "static",
                "addresses": [{
                        "address": "10.1.14.201/24",
                        "gateway": "10.1.14.1"
                }],
                "dns": {
                        "nameservers": ["8.8.8.8"]
                }
        },
        "bridge": "ovsbr0",
        "device": "enp0s8"
}
```
device field is optional.

## Usage

```bash
sudo ip netns add ns1

sudo CNI_COMMAND=ADD CNI_CONTAINERID=ns1 CNI_NETNS=/var/run/netns/ns1 CNI_IFNAME=net0 CNI_PATH=`pwd` ./ovsbridge <static.conf
```
