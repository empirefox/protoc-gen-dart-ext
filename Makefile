
makefile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
makefile_dir := $(shell dirname ${makefile_path})
dart_path := $(abspath $(makefile_dir)/dart)
test_path := $(abspath $(makefile_dir)/test)
cmd_path := $(abspath $(makefile_dir)/cmd)
pkg_path := $(abspath $(makefile_dir)/pkg)
protos_path := $(abspath $(makefile_dir)/pgde)
tools_path := $(abspath $(makefile_dir)/tools)

dart_lib := ${dart_path}/lib
dart_src := ${dart_lib}/src
dart_test := ${dart_path}/test

.PHONY: gen_validate_arb
gen_validate_arb:
	@cd ${dart_path} && \
		flutter pub pub run intl_translation:extract_to_arb \
		--locale=en \
		--output-dir=lib/src/validate \
		--output-file=validate.arb \
		lib/src/validate/validate.l10n.dart

.PHONY: rewrite_validate_arb
rewrite_validate_arb:
	@go run ${cmd_path}/rewrite_arb_langs/*.go \
		-arb=${dart_src}/validate/validate.arb \
		-langs=${dart_src}/validate/validate.arb.toml

.PHONY: gen_golang
gen_golang:
	@easyjson ${pkg_path}/arb/arb_attr.go
	@easyjson -no_std_marshalers ${pkg_path}/arb/arb.go
	@go generate ${pkg_path}/dart
	@go generate ${pkg_path}/genshared
	@cd ${cmd_path}/rewrite_arb_langs && go generate

.PHONY: gen_plural
gen_plural:
	@go run ${tools_path}/plural/plural.go \
		-dart=${dart_src}/plural/plural.dart \
		-dart_test=${dart_test}/plural_test.dart
	@dartfmt -w ${dart_src}/plural/plural.dart
	@dartfmt -w ${dart_test}/plural_test.dart

.PHONY: gen_atom
gen_atom:
	@go run ${tools_path}/atom/units_atom.go \
		-proto=${protos_path}/units/atom.proto \
		-dart=${dart_src}/units/atom.dart \
		-arb=${dart_src}/units/atom \
		-lang="en,en-US,zh,ar"
	@dartfmt -w ${dart_src}/units/atom.dart

.PHONY: gen_currency
gen_currency:
	@go run ${tools_path}/currency/currency_iso4217.go \
		-proto=${protos_path}/format/currency.proto \
		-dart=${dart_src}/format/currency.dart \
		-arb=${dart_src}/format/currency.arb
	@dartfmt -w ${dart_src}/format/currency.dart

.PHONY: gen_prefix
gen_prefix:
	@go run ${tools_path}/prefix/units_prefix.go \
		-proto=${protos_path}/units/prefix.proto \
		-dart=${dart_src}/units/prefix.dart \
		-arb=${dart_src}/units/prefix.arb
	@dartfmt -w ${dart_src}/units/prefix.dart

.PHONY: gen_protoc
gen_protoc:
	@protoc -I${makefile_dir} --go_out=paths=source_relative:${makefile_dir}/pkg ${protos_path}/l10n/*.proto
	@protoc -I${makefile_dir} --go_out=paths=source_relative:${makefile_dir}/pkg ${protos_path}/format/*.proto
	@protoc -I${makefile_dir} --go_out=paths=source_relative:${makefile_dir}/pkg ${protos_path}/units/*.proto
	@protoc -I${makefile_dir} --go_out=paths=source_relative:${makefile_dir}/pkg ${protos_path}/zero/*.proto
	@protoc -I${makefile_dir} --go_out=paths=source_relative:${makefile_dir}/pkg ${protos_path}/form/*.proto
	@protoc -I${makefile_dir} --go_out=paths=source_relative:${makefile_dir} ${test_path}/*.proto

# generate PgdeLocalizations
.PHONY: gen_pgde_gtt_to_dart
gen_pgde_gtt_to_dart:
	@SRC_DIR=${dart_src} \
		go run ${cmd_path}/gtt_to_dart/*.go \
		-gtt=${dart_src}/l10n/gtt.toml \
		-dart_out=${dart_src}/l10n/pgde.l10n.dart
	@dartfmt -w ${dart_src}/l10n/pgde.l10n.dart

.PHONY: gen_proto
gen_proto:
	@echo "no protos to generate"

.PHONY: test_plural
test_plural:
	@cd ${dart_path} && flutter pub test ./test/plural_test.dart

.PHONY: protoc_clean
protoc_clean:
	rm -f ${protos_path}/units/*.pb.go
	rm -f ${protos_path}/form/*.pb.go
	rm -f ${test_path}/*.pb.go
	rm -rf ${dart_path}/lib/units
	rm -rf ${dart_path}/lib/form
	rm -rf ${dart_path}/lib/test

.PHONY: protoc_test
protoc_test:
	protoc -I. \
		--dart-ext_out=arb,l10n=${dart_test}/gtt.toml,validate,form:${dart_lib} \
		protos/units/*.proto
	dartfmt -w ${dart_lib}/pgde/**/*.l10n.base.dart

.PHONY: gen_gtt_to_dart
gen_gtt_to_dart:
	@PROJECT_DIR=${makefile_dir} \
		go run ${cmd_path}/gtt_to_dart/*.go \
		-gtt=${test_path}/gtt_to_dart/gtt.toml \
		-dart_out=${dart_lib}/test/app.l10n.dart \
	@dartfmt -w ${dart_lib}/test/app.l10n.dart
