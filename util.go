package kstatus

import (
	"strings"

	corev1 "k8s.io/api/core/v1"
	apiunstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// NewReconcilingCondition creates an reconciling condition with the given
// reason and message.
func NewReconcilingCondition(reason, message string) Condition {
	return Condition{
		Type:    ConditionReconciling,
		Status:  corev1.ConditionTrue,
		Reason:  reason,
		Message: message,
	}
}

func NewStalledCondition(reason, message string) Condition {
	return Condition{
		Type:    ConditionStalled,
		Status:  corev1.ConditionTrue,
		Reason:  reason,
		Message: message,
	}
}

// NewInProgressStatus creates a status Result with the InProgress status
// and an InProgress condition.
func NewInProgressStatus(reason, message string) *Result {
	return &Result{
		Status:     InProgressStatus,
		Message:    message,
		Conditions: []Condition{NewReconcilingCondition(reason, message)},
	}
}

func NewFailedStatus(reason, message string) *Result {
	return &Result{
		Status:     FailedStatus,
		Message:    message,
		Conditions: []Condition{NewStalledCondition(reason, message)},
	}
}

// ObjWithConditions Represent meta object with status.condition array
type ObjWithConditions struct {
	// Status as expected to be present in most compliant kubernetes resources
	Status ConditionStatus `json:"status" yaml:"status"`
}

// ConditionStatus represent status with condition array
type ConditionStatus struct {
	// Array of Conditions as expected to be present in kubernetes resources
	Conditions []BasicCondition `json:"conditions" yaml:"conditions"`
}

// BasicCondition fields that are expected in a condition
type BasicCondition struct {
	// Type Condition type
	Type string `json:"type" yaml:"type"`
	// Status is one of True,False,Unknown
	Status corev1.ConditionStatus `json:"status" yaml:"status"`
	// Reason simple single word reason in CamleCase
	// +optional
	Reason string `json:"reason,omitempty" yaml:"reason"`
	// Message human readable reason
	// +optional
	Message string `json:"message,omitempty" yaml:"message"`
}

// GetObjectWithConditions return typed object
func GetObjectWithConditions(in map[string]interface{}) (*ObjWithConditions, error) {
	out := new(ObjWithConditions)
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func HasConditionWithStatus(conditions []BasicCondition, conditionType string, status corev1.ConditionStatus) bool {
	_, found := GetConditionWithStatus(conditions, conditionType, status)
	return found
}

func GetConditionWithStatus(conditions []BasicCondition, conditionType string, status corev1.ConditionStatus) (BasicCondition, bool) {
	for _, c := range conditions {
		if c.Type == conditionType && c.Status == status {
			return c, true
		}
	}
	return BasicCondition{}, false
}

// GetStringField return field as string defaulting to value if not found
func GetStringField(obj map[string]interface{}, fieldPath string, defaultValue string) string {
	rv := defaultValue

	fields := strings.Split(fieldPath, ".")
	if fields[0] == "" {
		fields = fields[1:]
	}

	val, found, err := apiunstructured.NestedFieldNoCopy(obj, fields...)
	if !found || err != nil {
		return rv
	}

	if v, ok := val.(string); ok {
		return v
	}
	return rv
}

// GetIntField return field as string defaulting to value if not found
func GetIntField(obj map[string]interface{}, fieldPath string, defaultValue int) int {
	fields := strings.Split(fieldPath, ".")
	if fields[0] == "" {
		fields = fields[1:]
	}

	val, found, err := apiunstructured.NestedFieldNoCopy(obj, fields...)
	if !found || err != nil {
		return defaultValue
	}

	switch v := val.(type) {
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	}
	return defaultValue
}
