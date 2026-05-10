// Package v1 contains API Schema definitions for the ani.kubercloud.io v1 API group.
package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// ── Finalizers ────────────────────────────────────────────────────────────────

const (
	// FinalizerInferenceService is set on every InferenceService at creation.
	// The controller removes it only after all owned resources (Deployment, Service,
	// ConfigMap, the K8s Secret for decryption) have been deleted and all in-flight
	// requests have been drained.
	FinalizerInferenceService = "ani.kubercloud.io/inference-service-cleanup"
)

// ── Phase constants (the only valid Phase transitions) ────────────────────────

// Valid Phase values and their allowed predecessor states:
//
//	Pending     ← initial state (create)
//	Downloading ← Pending (model pull started)
//	Decrypting  ← Downloading (all shards pulled; encrypted model only)
//	Deploying   ← Downloading | Decrypting (vLLM pod scheduled)
//	Running     ← Deploying (health-check passed)
//	Stopping    ← Running | Deploying (delete requested)
//	Stopped     ← Stopping (all pods terminated, drain complete)
//	Failed      ← any state (unrecoverable error, max retries exceeded)
//
// Forbidden transitions (controller must reject):
//   - any state → Running  (only allowed from Deploying via health-check event)
//   - Stopped   → anything (terminal; user must delete and recreate)
//   - Failed    → anything (terminal; user must delete and recreate)
const (
	PhasePending     = "Pending"
	PhaseDownloading = "Downloading"
	PhaseDecrypting  = "Decrypting"
	PhaseDeploying   = "Deploying"
	PhaseRunning     = "Running"
	PhaseStopping    = "Stopping"
	PhaseStopped     = "Stopped"
	PhaseFailed      = "Failed"
)

// ── Condition types ───────────────────────────────────────────────────────────

const (
	ConditionModelReady    = "ModelReady"    // model downloaded (and decrypted if encrypted)
	ConditionPodScheduled  = "PodScheduled"  // vLLM pod exists and is scheduled
	ConditionHealthy       = "Healthy"       // vLLM health endpoint returns 200
	ConditionDrainComplete = "DrainComplete" // all in-flight requests finished before stop
)

// ── Spec ──────────────────────────────────────────────────────────────────────

// InferenceServiceSpec defines the desired state of InferenceService.
type InferenceServiceSpec struct {
	// Model is the model ID from the ANI model registry, format "name:version".
	// Immutable after creation; change requires delete + recreate.
	// +kubebuilder:validation:Pattern=`^[a-z0-9.-]+:[a-z0-9.-]+$`
	Model string `json:"model"`

	// Replicas is the desired number of inference pod replicas.
	// +kubebuilder:default=1
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=32
	Replicas int32 `json:"replicas,omitempty"`

	// GPUType specifies the required GPU model (e.g., "A100", "A10", "910B").
	// Leave empty for CPU-only inference.
	GPUType string `json:"gpuType,omitempty"`

	// GPUCountPerPod is the number of GPUs allocated to each replica pod.
	// +kubebuilder:default=1
	// +kubebuilder:validation:Minimum=1
	GPUCountPerPod int32 `json:"gpuCountPerPod,omitempty"`

	// MaxConcurrency is the maximum number of in-flight inference requests across all replicas.
	// The gateway queues requests beyond this limit rather than returning 429.
	// +kubebuilder:default=8
	// +kubebuilder:validation:Minimum=1
	MaxConcurrency int32 `json:"maxConcurrency,omitempty"`

	// Placement controls which Region/AZ the service is deployed to.
	// Karmada PropagationPolicy is derived from this field.
	Placement *PlacementSpec `json:"placement,omitempty"`

	// EncryptionKeyRef references a K8s Secret that contains the model decryption password.
	// Required only when the referenced model version has is_encrypted=true.
	// The Secret is owned by this InferenceService and deleted when the service is deleted.
	EncryptionKeyRef *SecretKeyRef `json:"encryptionKeyRef,omitempty"`

	// DrainTimeoutSeconds is how long the controller waits for in-flight requests to finish
	// before force-deleting pods during shutdown. Default 30 seconds.
	// +kubebuilder:default=30
	DrainTimeoutSeconds int32 `json:"drainTimeoutSeconds,omitempty"`
}

// PlacementSpec defines the Region/AZ placement preference.
type PlacementSpec struct {
	// Region is the ANIRegion name (e.g., "cn-east"). Optional.
	Region string `json:"region,omitempty"`
	// AZ is the ANIAvailabilityZone name (e.g., "cn-east-1a"). Optional.
	AZ string `json:"az,omitempty"`
}

// SecretKeyRef identifies a key within a K8s Secret.
type SecretKeyRef struct {
	// SecretName is the name of the K8s Secret in the same namespace.
	SecretName string `json:"secretName"`
	// Key is the data key within the Secret that holds the decryption password.
	Key string `json:"key"`
}

// ── Status ────────────────────────────────────────────────────────────────────

// InferenceServiceStatus defines the observed state of InferenceService.
type InferenceServiceStatus struct {
	// Phase is the current lifecycle phase. See phase constants above for valid transitions.
	// +kubebuilder:validation:Enum=Pending;Downloading;Decrypting;Deploying;Running;Stopping;Stopped;Failed
	Phase string `json:"phase,omitempty"`

	// ObservedGeneration is the .metadata.generation that was last acted on.
	// Used to detect spec changes that have not yet been reconciled.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// EndpointURL is the internal Kubernetes Service URL reachable by the ANI Gateway.
	// Set only when Phase == Running.
	EndpointURL string `json:"endpointURL,omitempty"`

	// ReadyReplicas is the number of replicas currently passing the health check.
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// Message is a human-readable description of the current state or error.
	// Required when Phase == Failed.
	Message string `json:"message,omitempty"`

	// Conditions provides detailed status for each lifecycle gate.
	// Standard condition types: ModelReady, PodScheduled, Healthy, DrainComplete.
	// Each condition follows the Kubernetes meta/v1.Condition conventions:
	//   - LastTransitionTime is always set when Status changes.
	//   - Reason is a CamelCase identifier for the cause.
	//   - Message is a human-readable sentence.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// ── Root object ───────────────────────────────────────────────────────────────

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=isvc
// +kubebuilder:printcolumn:name="Model",type=string,JSONPath=`.spec.model`
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Ready",type=integer,JSONPath=`.status.readyReplicas`
// +kubebuilder:printcolumn:name="Endpoint",type=string,JSONPath=`.status.endpointURL`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// InferenceService is the Schema for the inferenceservices API.
type InferenceService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InferenceServiceSpec   `json:"spec,omitempty"`
	Status InferenceServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InferenceServiceList contains a list of InferenceService.
type InferenceServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InferenceService `json:"items"`
}
