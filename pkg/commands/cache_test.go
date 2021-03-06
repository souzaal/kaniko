/*
Copyright 2020 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"testing"
)

func Test_caching(t *testing.T) {
	c := caching{layer: fakeLayer{}, readSuccess: true}

	actual := c.Layer().(fakeLayer)
	expected := fakeLayer{}
	actualLen, expectedLen := len(actual.TarContent), len(expected.TarContent)
	if actualLen != expectedLen {
		t.Errorf("expected layer tar content to be %v but was %v", expectedLen, actualLen)
	}

	if !c.ReadSuccess() {
		t.Errorf("expected ReadSuccess to be %v but was %v", true, c.ReadSuccess())
	}
}
