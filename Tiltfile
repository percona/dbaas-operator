# Generate the DBaaS operator manifests
local_resource(
  'dbaas-operator-gen-manifests',
  'make fmt generate manifests',
  deps=['./api'],
  ignore=['./api/v1/zz_generated.deepcopy.go'],
)

# Build the DBaaS operator manager locally to take advantage of the go cache
local_resource(
  'dbaas-operator-build',
  'make fmt; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/manager main.go',
  deps=['./api', './controllers', './main.go'],
)

# Live update the DBaaS operator manager without generating a new pod
load('ext://restart_process', 'docker_build_with_restart')
docker_build_with_restart('percona/dbaas-operator', '.',
  dockerfile='tilt.Dockerfile',
  entrypoint='/home/tilt/manager',
  only=['./bin/manager'],
  live_update=[
    sync('./bin/manager', '/home/tilt/manager'),
  ]
)

# Apply DBaaS operator manifests
k8s_yaml(kustomize('./config/default', kustomize_bin='./bin/kustomize'))
k8s_resource(workload='dbaas-operator-controller-manager',
  new_name='dbaas-operator',
  resource_deps = [
    'pxc-operator-crds-ready',
    'psmdb-operator-crds-ready'
  ]
)

# Deploy PXC and PSMDB operators
# Define the operator versions with PXC_OPERATOR_VERSION and
# PSMDB_OPERATOR_VERSION environment variables
load('ext://helm_remote', 'helm_remote')
helm_remote('pxc-operator',
  repo_name='percona',
  repo_url='https://percona.github.io/percona-helm-charts',
  namespace='dbaas-operator-system',
  version=os.getenv('PXC_OPERATOR_VERSION', '') ,
)
helm_remote('psmdb-operator',
  repo_name='percona',
  repo_url='https://percona.github.io/percona-helm-charts',
  namespace='dbaas-operator-system',
  version=os.getenv('PSMDB_OPERATOR_VERSION', '') ,
)
