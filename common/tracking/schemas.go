package tracking

import (
	"time"

	"github.com/bentoml/yatai-schemas/modelschemas"
)

type YataiEventType string

const (
	YataiDeploymentCreate    YataiEventType = "yatai_deployment_create"
	YataiDeploymentUpdate    YataiEventType = "yatai_deployment_update"
	YataiDeploymentTerminate YataiEventType = "yatai_deployment_terminate"
	YataiDeploymentDelete    YataiEventType = "yatai_deployment_delete"
	YataiBentoPull           YataiEventType = "yatai_bento_pull"
	YataiBentoPush           YataiEventType = "yatai_bento_pull"
	YataiModelPull           YataiEventType = "yatai_model_pull"
	YataiModelPush           YataiEventType = "yatai_model_push"
	YataiLifeCycleStartup    YataiEventType = "yatai_lifecycle_startup"
	YataiLifeCycleShutdown   YataiEventType = "yatai_lifecycle_shutdown"
	YataiLifeCyclePeriodic   YataiEventType = "yatai_lifecycle_periodic"
)

type CommonProperties struct {
	EventType       string    `json:"event_type"`
	OrganizationUID string    `json:"organization_uid"`
	Timestamp       time.Time `json:"timestamp"`
	YataiVersion    string    `json:"yatai_version"`
}

func NewCommonProperties(eventType YataiEventType, organizationUID string, yataiVersion string) CommonProperties {
	return CommonProperties{
		EventType:       string(eventType),
		OrganizationUID: organizationUID,
		YataiVersion:    yataiVersion,
		Timestamp:       time.Now(),
	}
}

// type LifecycleEvent struct {
// 	LifecycleEventType string        `json:"lifecycle_eventtype"`
// 	Uptime             time.Duration `json:"uptime"`
// }

type DeploymentEvent struct {
	CommonProperties
	UserUID               string                                              `json:"user_uid"`
	ClusterUID            string                                              `json:"cluster_uid"`
	DeploymentUID         string                                              `json:"deployment_uid"`
	DeploymentStatus      modelschemas.DeploymentStatus                       `json:"deployment_status"`
	DeploymentRevisionID  string                                              `json:"deployment_revision_id,omitempty"`
	DeploymentTargetTypes []modelschemas.DeploymentTargetType                 `json:"deployment_target_types,omitempty"`
	ApiServerResources    []modelschemas.DeploymentTargetResources            `json:"api_server_resources,omitempty"`
	ApiServerHPAConfig    []modelschemas.DeploymentTargetHPAConf              `json:"api_server_hpa_config,omitempty"`
	RunnerResourcesList   []map[string]modelschemas.DeploymentTargetResources `json:"runner_resources_list,omitempty"`
	RunnerHPAConfigList   []map[string]modelschemas.DeploymentTargetHPAConf   `json:"runner_hpa_config_list,omitempty"`
	// DeploymentTargetCanaryRuleTypes [][]modelschemas.DeploymentTargetCanaryRuleType
}

type BentoEvent struct {
	CommonProperties
	UserUID                   string                            `json:"user_uid"`
	BentoRepositoryUID        string                            `json:"bentorepository_uid"`
	BentoVersion              string                            `json:"bento_version"`
	BentoUploadStatus         modelschemas.BentoUploadStatus    `json:"bento_upload_status"`
	BentoUploadFinishedReason string                            `json:"bento_upload_finished_reason"`
	BentoTransmissionStrategy modelschemas.TransmissionStrategy `json:"bento_transmission_strategy"`
	BentoSizeBytes            uint                              `json:"bento_size_bytes"`
	NumModels                 int                               `json:"num_models"`
	NumRunners                int                               `json:"num_runners"`
}


type ModelEvent struct {
	CommonProperties
	UserUID                   string                            `json:"user_uid"`
	ModelUID                  string                            `json:"model_uid"`
	ModelUploadStatus         modelschemas.ModelUploadStatus    `json:"model_upload_status"`
	ModelUploadFinishedReason string                            `json:"model_upload_finished_reason"`
	ModelTransmissionStrategy modelschemas.TransmissionStrategy `json:"model_transmission_strategy"`
	ModelSizeBytes            uint                              `json:"model_size_bytes"`
}

type LifeCycleEvent struct {
	CommonProperties
	Uptime time.Duration
	YataiEventType
}
