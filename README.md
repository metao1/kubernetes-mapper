# k8s Mapper 

Kubernete Mapper is a tool to generate Kubernetes architecture diagrams from the actual state in a namespace.

## Prerequisites
Go version 1.15

### Go version
`k8mapper` only depends dot (graphviz) command.

## Installation

```
$ export GO115MODULE=on
$ go build -o k8mapper .
```

`icons` directory needs to be in the same directory to the binary.

## Usage
### Bash script version
```
$ ./k8mapper -h
Usage of ./k8sviz:
  -kubeconfig string
        absolute path to the kubeconfig file (default "/root/.kube/config")
  -n string
        namespace to visualize (shorthand) (default "namespace")
  -namespace string
        namespace to visualize (default "namespace")
  -o string
        output filename (shorthand) (default "k8sviz.out")
  -outfile string
        output filename (default "k8sviz.out")
  -t string
        type of output (shorthand) (default "dot")
  -type string
        type of output (default "dot")
```

## Examples
```
- Generate png file for namespace `default`
```
$ ./k8mapper -n topic-clustering -t png -o topic-clustering.png
```
- Output for [an example topic clustering namespace
   - [topic-clustering.png](./topic-clustering.png):
```