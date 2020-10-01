// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type EncryptionKeyType string

// Enum values for EncryptionKeyType
const (
	EncryptionKeyTypeKms EncryptionKeyType = "KMS"
)

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusSubmitted            JobStatus = "SUBMITTED"
	JobStatusIn_progress          JobStatus = "IN_PROGRESS"
	JobStatusCompleted            JobStatus = "COMPLETED"
	JobStatusCompleted_with_error JobStatus = "COMPLETED_WITH_ERROR"
	JobStatusFailed               JobStatus = "FAILED"
	JobStatusStop_requested       JobStatus = "STOP_REQUESTED"
	JobStatusStopped              JobStatus = "STOPPED"
)

type MergeStrategy string

// Enum values for MergeStrategy
const (
	MergeStrategyOverwrite MergeStrategy = "OVERWRITE"
)

type TerminologyDataFormat string

// Enum values for TerminologyDataFormat
const (
	TerminologyDataFormatCsv TerminologyDataFormat = "CSV"
	TerminologyDataFormatTmx TerminologyDataFormat = "TMX"
)