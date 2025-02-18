// Copyright 2018 Chris Pearce
// Copyright 2022 Nokia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Modified by Nokia into an importable package.

package arm

import "errors"

var (
	ErrMinSupportOutOfRange    = errors.New("MinSupport value is out of range [0,1.0].")
	ErrMinConfidenceOutOfRange = errors.New("MinConfidence value is out of range [0,1.0].")
	ErrMinLiftOutOfRange       = errors.New("MinLift is out of range [1.0,∞].")
)

type Arguments struct {
	// Input dataset in CSV format.
	Input string
	// File path in which to store Output rules. Format:
	// antecedent -> consequent, confidence, lift, support.
	Output string
	// Minimum itemset support threshold, in range [0,1].
	MinSupport float64
	// Minimum rule confidence threshold, in range [0,1].
	MinConfidence float64
	// Minimum rule lift confidence threshold, in range
	// [1,∞] (optional).
	MinLift float64
	// File path in which to store generated itemsets
	// (optional).
	ItemsetsPath string
}

func (args Arguments) Validate() error {
	if args.MinSupport < 0.0 || args.MinSupport > 1.0 {
		return ErrMinSupportOutOfRange
	}
	if args.MinConfidence < 0.0 || args.MinConfidence > 1.0 {
		return ErrMinConfidenceOutOfRange
	}
	if args.MinLift != 0.0 && args.MinLift < 1.0 {
		return ErrMinLiftOutOfRange
	}
	return nil
}
