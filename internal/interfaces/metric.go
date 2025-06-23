/*
 * HXE - Host-based Process Execution Engine
 * Copyright (C) 2025 Rangertaha <rangertaha@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package interfaces

import (
	"time"
)

// ValueType is an enumeration of metric types that represent a simple value.
type ValueType int

// // Possible values for the ValueType enum.
// const (
// 	_ ValueType = iota
// 	Counter
// 	Gauge
// 	Untyped
// 	Summary
// 	Histogram
// )

// Tag represents a single tag key and value.
type Tag struct {
	Key   string
	Value string
}

// Field represents a single field key and value.
type Field struct {
	Key   string
	Value interface{}
}

// Metric is the type of data that is processed by ares.  Input plugins,
// and to a lesser degree, Processor and Aggregator plugins create new Metrics
// and Output plugins write them.
// nolint:interfacebloat // conditionally allow to contain more methods
type Metric interface {
	// Name is the primary identifier for the Metric and corresponds to the
	// measurement in the InfluxDB data model.
	Name() string

	// Tags returns the tags as a map.  This method is deprecated, use TagList instead.
	Tags() []Tag

	// // TagList returns the tags as a slice ordered by the tag key in lexical
	// // bytewise ascending order.  The returned value should not be modified,
	// // use the AddTag or RemoveTag methods instead.
	// TagList() []*Tag

	// 	// GetTag returns the value of a tag and a boolean to indicate if it was set.
	// 	GetTag(key string) (string, bool)

	// 	// HasTag returns true if the tag is set on the Metric.
	// 	HasTag(key string) bool

	// Fields returns the fields as a map.  This method is deprecated, use FieldList instead.
	Fields() []Field

	// // FieldList returns the fields as a slice in an undefined order.  The
	// // returned value should not be modified, use the AddField or RemoveField
	// // methods instead.
	// FieldList() []*Field

	// Time returns the timestamp of the metric.
	Time() time.Time

	// IsOrder() bool

	// // AddTag sets the tag on the Metric.  If the Metric already has the tag
	// // set then the current value is replaced.
	// AddTag(key, value string)

	// // RemoveTag removes the tag if it is set.
	// RemoveTag(key string)

	// // GetField returns the value of a field and a boolean to indicate if it was set.
	// GetField(key string) (interface{}, bool)

	// // GetFloat64 returns the float64 value of a field and a boolean to indicate if it was set.
	// GetFloat64(key string) (float64, bool)

	// // HasField returns true if the field is set on the Metric.
	// HasField(key string) bool

	// // AddField sets the field on the Metric.  If the Metric already has the field
	// // set then the current value is replaced.
	// AddField(key string, value interface{})

	// // RemoveField removes the tag if it is set.
	// RemoveField(key string)

	// // SetTime sets the timestamp of the Metric.
	// SetTime(t time.Time)

	// // HashID returns an unique identifier for the series.
	// HashID() uint64

	// // Copy returns a deep copy of the Metric.
	// Copy() Metric

	// // Accept marks the metric as processed successfully and written to an
	// // output.
	// Accept()

	// // Reject marks the metric as processed unsuccessfully.
	// Reject()

	// // Drop marks the metric as processed successfully without being written
	// // to any output.
	// Drop()
}
