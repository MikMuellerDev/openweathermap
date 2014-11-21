// Copyright 2014 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openweathermap

import (
	owm "github.com/briandowns/openweathermap"
	"reflect"
	"testing"
)

func TestNewCurrent(t *testing.T) {
	t.Parallel()
	for d, _ := range owm.DataUnits {
		t.Logf("Data unit: %s", d)
		if owm.ValidDataUnit(d) {
			c, err := owm.NewCurrent(d)
			if err != nil {
				t.Error(err)
			}
			if reflect.TypeOf(c).String() != "*openweathermap.CurrentWeatherData" {
				t.Error("ERROR: incorrect data type returned")
			}
		} else {
			t.Errorf("ERROR: unusable data unit - %s", d)
		}
	}
	_, err := owm.NewCurrent("asdf")
	if err == nil {
		t.Error("ERROR: created instance when it shouldn't have")
	}
}

func TestCurrentByName(t *testing.T) {
	testCities := []string{"Philadelphia", "Newark", "Helena"}
	c, err := owm.NewCurrent("imperial")
	if err != nil {
		t.Error(err)
	}
	for _, city := range testCities {
		c.CurrentByName(city)
		if c.Name != city {
			t.Error("ERROR: Incorrect city returned")
		}
	}
}

func TestCurrentByCoordinates(t *testing.T) {
	c, err := owm.NewCurrent("imperial")
	if err != nil {
		t.Error("Error creating instance of CurrentWeatherData")
	}
	c.CurrentByCoordinates(
		&owm.Coordinates{
			Longitude: -112.07,
			Latitude:  33.45,
		},
	)
}

func TestCurrentByID(t *testing.T) {
	c, err := owm.NewCurrent("metric")
	if err != nil {
		t.Error("Error creating instance of CurrentWeatherData")
	}
	c.CurrentByID(5344157)
}

func TestCurrentByArea(t *testing.T) {}