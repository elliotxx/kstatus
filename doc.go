// Package kstatus contains functionality for computing the status
// of Kubernetes resources.
//
// The statuses defined in this package are:
//   - InProgress
//   - Current
//   - Failed
//   - Terminating
//   - NotFound
//   - Unknown
//
// Computing the status of a resources can be done by calling the
// Compute function in the status package.
//
//	import (
//	  "github.com/elliotxx/kstatus"
//	)
//
//	res, err := kstatus.Compute(resource)
//
// The package also defines a set of new conditions:
//   - InProgress
//   - Failed
//
// These conditions have been chosen to follow the
// "abnormal-true" pattern where conditions should be set to true
// for error/abnormal conditions and the absence of a condition means
// things are normal.
package kstatus
