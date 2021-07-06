/*
Copyright 2021 Adobe. All rights reserved.
This file is licensed to you under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License. You may obtain a copy
of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR REPRESENTATIONS
OF ANY KIND, either express or implied. See the License for the specific language
governing permissions and limitations under the License.
*/

package domain

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type ChartValues struct {
	Values map[interface{}]interface{}
}

func (v ChartValues) loadValuesFile(valuesFilePath string) ChartValues {
	d := Datastore{}
	valuesBuffer := d.ReadFile(valuesFilePath)
	return v.loadValues(valuesBuffer)
}

func (v ChartValues) loadValues(byteData []byte) ChartValues {
	err := yaml.Unmarshal(byteData, &v.Values)
	if err != nil {
		log.Fatalf("Error loading values: %v", err)
	}
	return v
}