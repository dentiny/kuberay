package util

const (
	RayVersion = "2.41.0"
	RayImage   = "rayproject/ray:" + RayVersion

	RayClusterLabelKey   = "ray.io/cluster"
	RayIsRayNodeLabelKey = "ray.io/is-ray-node"
	RayNodeGroupLabelKey = "ray.io/group"
	RayNodeTypeLabelKey  = "ray.io/node-type"

	ResourceNvidiaGPU = "nvidia.com/gpu"
	ResourceGoogleTPU = "google.com/tpu"

	FieldManager = "ray-kubectl-plugin"

	// NodeSelector
	NodeSelectorGKETPUAccelerator = "cloud.google.com/gke-tpu-accelerator"
	NodeSelectorGKETPUTopology    = "cloud.google.com/gke-tpu-topology"
)
