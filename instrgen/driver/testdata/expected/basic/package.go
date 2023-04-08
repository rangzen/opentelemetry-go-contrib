// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//nolint:all
package main

import (
	"os"
	__atel_context "context"

	"go.opentelemetry.io/contrib/instrgen/rtlib"
	__atel_otel "go.opentelemetry.io/otel"
)

func Close() error {
	return nil
}

func pack(__atel_tracing_ctx __atel_context.Context,) {
	__atel_child_tracing_ctx, __atel_span := __atel_otel.Tracer("pack").Start(__atel_tracing_ctx, "pack")
	_ = __atel_child_tracing_ctx
	defer __atel_span.End()
	f, e := os.Create("temp")
	defer f.Close()
	if e != nil {

	}
}