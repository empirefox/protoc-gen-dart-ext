
makefile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
makefile_dir := $(shell dirname ${makefile_path})
dart_path := $(abspath $(makefile_dir)/dart)
test_path := $(abspath $(makefile_dir)/test)
cmd_path := $(abspath $(makefile_dir)/cmd)
pkg_path := $(abspath $(makefile_dir)/pkg)
protos_path := $(abspath $(makefile_dir)/protos)
tools_path := $(abspath $(makefile_dir)/tools)
format_path := $(abspath $(makefile_dir)/protos/format)
units_path := $(abspath $(makefile_dir)/protos/units)
form_path := $(abspath $(makefile_dir)/protos/form)

go_out := ${GOPATH}/src
dart_lib := ${dart_path}/lib
dart_test := ${dart_path}/test

.PHONY: gen_atom
gen_atom:
	@go run ${tools_path}/atom/units_atom.go \
		-proto=${units_path}/atom.proto \
		-dart=${dart_lib}/units/atom.dart \
		-arb=${dart_lib}/units/atom \
		-lang="en,en-US,zh,ar"
	@dartfmt -w ${dart_lib}/units/atom.dart

.PHONY: gen_currency
gen_currency:
	@go run ${tools_path}/currency/currency_iso4217.go \
		-proto=${format_path}/currency.proto \
		-dart=${dart_lib}/units/currency.dart \
		-arb=${dart_lib}/units/currency.arb
	@dartfmt -w ${dart_lib}/units/currency.dart

.PHONY: gen_prefix
gen_prefix:
	@go run ${tools_path}/prefix/units_prefix.go \
		-proto=${units_path}/prefix.proto \
		-dart=${dart_lib}/units/prefix.dart \
		-arb=${dart_lib}/units/prefix.arb
	@dartfmt -w ${dart_lib}/units/prefix.dart

.PHONY: gen_units_gtt_to_dart
gen_units_gtt_to_dart:
	@GTT_DIR=${dart_lib}/units \
		go run ${cmd_path}/gtt_to_dart/*.go \
		-gtt=${dart_lib}/units/gtt.toml \
		-dart_import_path="package:pgde/units/units.l10n.dart" \
		-dart_out=${dart_lib}/units/units.l10n.dart \
		-with_delegate \
		-exports_out=${dart_lib}/units/units.l10n.exports.dart \
		-imports_out=${dart_lib}/units/units.l10n.imports.pbdata
	@dartfmt -w ${dart_lib}/units/units.l10n.dart

.PHONY: gen_proto
gen_proto:
	@echo "no protos to generate"

.PHONY: gen_golang
gen_golang:
	@easyjson ${pkg_path}/arb/arb_attr.go
	@easyjson ${pkg_path}/arb/resolve.go
	@easyjson -no_std_marshalers ${pkg_path}/arb/arb.go
	@go generate ${pkg_path}/dart
	@go generate ${pkg_path}/genshared

.PHONY: gen_plural
gen_plural:
	@go run ${tools_path}/plural/plural.go \
		-dart=${dart_lib}/plural/plural.dart \
		-dart_test=${dart_test}/plural_test.dart
	@dartfmt -w ${dart_lib}/plural/plural.dart
	@dartfmt -w ${dart_test}/plural_test.dart

.PHONY: test_plural
test_plural:
	@cd ${dart_path} && flutter pub pub run test ./test/plural_test.dart

.PHONY: protoc
protoc:
	@protoc -I. --go_out=${go_out} --dart_out=${dart_lib} ./protos/exports/*.proto
	@protoc -I. --go_out=${go_out} ./protos/imports/*.proto
	@protoc -I. --go_out=${go_out} ./protos/l10n/*.proto
	@protoc -I. --go_out=${go_out} ./protos/format/*.proto
	@protoc -I. --go_out=${go_out} ./protos/units/*.proto
	@protoc -I. --go_out=${go_out} --dart_out=${dart_lib} ./protos/form/*.proto
	# @protoc -I. --go_out=${go_out} --dart_out=${dart_lib} ./test/*.proto

.PHONY: protoc_clean
protoc_clean:
	rm -f ${units_path}/*.pb.go
	rm -f ${form_path}/*.pb.go
	rm -f ${test_path}/*.pb.go
	rm -rf ${dart_path}/lib/units
	rm -rf ${dart_path}/lib/form
	rm -rf ${dart_path}/lib/test

.PHONY: protoc_test
protoc_test:
	protoc -I. --dart-ext_out=dart_ext=pkg_l10n_base:${dart_lib} protos/units/*.proto
	dartfmt -w ${dart_lib}/protos/**/*.l10n.base.dart

.PHONY: gen_gtt_to_dart
gen_gtt_to_dart:
	@PROJECT_DIR=${makefile_dir} \
		go run ${cmd_path}/gtt_to_dart/*.go \
		-gtt=${test_path}/gtt_to_dart/gtt.toml \
		-dart_import_path="package:pgde/test/app.l10n.exports.dart" \
		-dart_out=${dart_lib}/test/app.l10n.dart \
		-exports_out=${dart_lib}/test/app.l10n.exports.dart \
		-imports_out=${dart_lib}/test/app.l10n.imports.pbdata
	@dartfmt -w ${dart_lib}/test/app.l10n.dart

.PHONY: gen_exports
gen_exports:
	@cd ${dart_path} && flutter pub pub run tools/exports.dart -o ${dart_path}/pgde.exports.data
	@go run ${tools_path}/test/main.go -exports ${dart_path}/pgde.exports.data

.PHONY: gen_exports_go
gen_exports_go:
	@go run ${tools_path}/goexports/goexports.go -dartExports ${dart_lib}/pgde.exports.dart

.PHONY: gen_fieldmask
gen_fieldmask:
		# -I${go_out}/github.com/TheThingsIndustries/protoc-gen-fieldmask/vendor
	@protoc -I. \
		--fieldmask_out=lang=go:${go_out} \
		./protos/l10n/*.proto

