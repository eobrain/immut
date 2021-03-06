# Copyright 2013 Eamonn O'Brien-Strain
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

ENV=GOPATH=`pwd`/../../../..

dotest:
	$(ENV) go test github.com/eobrain/immut/test

bench:
	$(ENV) go test --bench=. --benchmem --benchtime=0.01s github.com/eobrain/immut/test
