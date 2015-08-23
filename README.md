# golookup

Super-simple DNS lookup tool in Go

## Raison d'etre

Ubuntu installs three million dependencies for the `bind9-host` package and I needed a simple way to generate a comma-separated unicast list of IPs for Elasticsearch node discovery in a Rancher/Docker environment.

## Usage

    $ golookup www.google.com
    173.194.116.112,173.194.116.113,173.194.116.114,173.194.116.116,173.194.116.115
