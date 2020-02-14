
makefile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
makefile_dir := $(shell dirname ${makefile_path})
dart_path := $(abspath $(makefile_dir)/dart)
test_path := $(abspath $(makefile_dir)/test)
hybrid_path := $(abspath $(makefile_dir)/hybrid)
google_path := $(abspath $(makefile_dir)/google)
cmd_path := $(abspath $(makefile_dir)/cmd)
pkg_path := $(abspath $(makefile_dir)/pkg)
protos_path := $(abspath $(makefile_dir)/pgde)
tools_path := $(abspath $(makefile_dir)/tools)
pgv_path := $(shell go list -u -m -json github.com/envoyproxy/protoc-gen-validate | jq .Dir)

dart_lib := ${dart_path}/lib
dart_src := ${dart_lib}/src
dart_test := ${dart_path}/test

.PHONY: prepare_protoc_dep
prepare_protoc_dep:
	@rm -rf ${makefile_dir}/validate
	@mkdir -p ${makefile_dir}/validate
	@cp -n ${pgv_path}/validate/validate.proto ${makefile_dir}/validate

.PHONY: gen_golang
gen_golang:
	@easyjson ${pkg_path}/arb/arb_attr.go
	@easyjson -no_std_marshalers ${pkg_path}/arb/arb.go
	@go generate ${pkg_path}/dart
	@go generate ${pkg_path}/genshared
	@cd ${cmd_path}/arb_rewrite_langs && go generate

.PHONY: gen_plural
gen_plural:
	@go run ${tools_path}/plural/plural.go \
		-dart=${dart_src}/plural/plural.dart \
		-dart_test=${dart_test}/plural_test.dart
	@dartfmt -w ${dart_src}/plural/plural.dart
	@dartfmt -w ${dart_test}/plural_test.dart

.PHONY: gen_atom
gen_atom:
	@go run ${tools_path}/atom/*.go \
		-proto=${protos_path}/units/atom.proto \
		-dart=${dart_src}/units/atom.dart \
		-arb=${dart_src}/units/atom.arb
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

.PHONY: gen_dart_lib
gen_dart_lib:
	@rm -rf ${pkg_path}/dart/dart_lib.go
	@MAKEFILE_DIR=${makefile_dir} go run ${tools_path}/dart_lib/*.go \
		-go ${pkg_path}/dart/dart_lib.go
	@gofmt -w ${pkg_path}/dart/dart_lib.go

.PHONY: gen_protoc
gen_protoc:
	@protoc -I${makefile_dir} --go_out=paths=source_relative:${makefile_dir}/pkg ${protos_path}/l10n/*.proto
	@protoc -I${makefile_dir} --go_out=paths=source_relative:${makefile_dir}/pkg ${protos_path}/zero/*.proto
	@protoc -I${makefile_dir} --go_out=paths=source_relative:${makefile_dir}/pkg ${protos_path}/form/*.proto
	@protoc -I${makefile_dir} --go_out=paths=source_relative:${makefile_dir}/pkg ${protos_path}/units/*.proto
	@protoc -I${makefile_dir} --go_out=paths=source_relative:${makefile_dir}/pkg ${protos_path}/format/*.proto
	@protoc -I${makefile_dir} --go_out=paths=source_relative:${makefile_dir}/pkg --dart_out=:${dart_src} ${protos_path}/error/*.proto

.PHONY: gen_pgde_arbs
gen_pgde_arbs:
	@cd ${dart_path} && \
		flutter pub run intl_translation:extract_to_arb \
		--locale=en \
		--output-dir=lib/src/validate \
		--output-file=validate.arb \
		lib/src/validate/validate.l10n.dart
	@go run ${cmd_path}/arb_rewrite_langs/*.go \
		-arb=${dart_src}/validate/validate.arb \
		-langs=${dart_src}/validate/validate.arb.toml
	# validate done.
	@cd ${dart_path} && \
		flutter pub run intl_translation:extract_to_arb \
		--locale=en \
		--output-dir=lib/src/form \
		--output-file=form.arb \
		lib/src/form/form.l10n.dart
	# form done.

.PHONY: gen_pgde_l10n_zip
gen_pgde_l10n_zip:
	@rm -rf ${dart_src}/l10n/archive
	@rm -f ${dart_src}/l10n/archive.zip

	# atom
	@go run ${cmd_path}/arb_variant/*.go \
		-input=${dart_src}/units/atom.arb \
		-lang=ar,en-US,zh \
		-output=${dart_src}/l10n/archive/%V/%N%E \
		-force=0
	# prefix
	@go run ${cmd_path}/arb_variant/*.go \
		-input=${dart_src}/units/prefix.arb \
		-lang=ar,en-US,zh \
		-output=${dart_src}/l10n/archive/%V/%N%E
	# currency
	@go run ${cmd_path}/arb_variant/*.go \
		-input=${dart_src}/format/currency.arb \
		-lang=ar,en-US,zh \
		-output=${dart_src}/l10n/archive/%V/%N%E
	# validate
	@go run ${cmd_path}/arb_variant/*.go \
		-input=${dart_src}/validate/validate.arb \
		-lang=ar,en-US,zh \
		-output=${dart_src}/l10n/archive/%V/%N%E
	# form
	@go run ${cmd_path}/arb_variant/*.go \
		-input=${dart_src}/form/form.arb \
		-lang=ar,en-US,zh \
		-output=${dart_src}/l10n/archive/%V/%N%E

	# archive
	@cd ${dart_src}/l10n && zip -r archive.zip archive

# generate PgdeLocalizations
.PHONY: gen_pgde_gtt_to_dart
gen_pgde_gtt_to_dart:
	@cd ${dart_src} && \
		SRC_DIR=${dart_src} \
		go run ${cmd_path}/gtt_to_dart/*.go \
		-gtt=l10n/gtt.toml \
		-dart_out=l10n/pgde.l10n.dart
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

.PHONY: protoc_hybrid_base
protoc_hybrid_base:
	@protoc -I${makefile_dir} -I${pgv_path} --dart_out=:${dart_test}/pgde ${google_path}/protobuf/duration.proto
	@protoc -I${makefile_dir} -I${pgv_path} --dart_out=:${dart_test}/pgde ${google_path}/protobuf/timestamp.proto
	@protoc -I${makefile_dir} -I${pgv_path} --dart_out=:${dart_test}/pgde ${google_path}/protobuf/empty.proto
	@protoc -I${makefile_dir} -I${pgv_path} --dart_out=:${dart_test}/pgde ${google_path}/protobuf/wrappers.proto
	@protoc -I${makefile_dir} -I${pgv_path} --dart_out=:${dart_test}/pgde ${protos_path}/error/error.proto
	@protoc -I${makefile_dir} -I${pgv_path} --dart_out=grpc:${dart_test}/pgde ${hybrid_path}/*.proto

.PHONY: protoc_hybrid_static
protoc_hybrid_static:
	@protoc -I${makefile_dir} \
		--static_out=google/protobuf/timestamp.fmt.dart:${dart_test}/pgde \
		${google_path}/protobuf/timestamp.proto
	@protoc -I${makefile_dir} \
		--static_out=google/protobuf/duration.fmt.dart:${dart_test}/pgde \
		${google_path}/protobuf/duration.proto

.PHONY: protoc_hybrid_l10n_zip
protoc_hybrid_l10n_zip:
	@rm -rf ${dart_test}/pgde/hybrid/archive
	@rm -f ${dart_test}/pgde/hybrid/archive.zip

	@protoc -I${makefile_dir} -I${pgv_path} \
		--dart-ext_out=arb:${dart_test}/pgde \
		${hybrid_path}/config.proto

	@go run ${cmd_path}/arb_variant/*.go \
		-input=${dart_test}/pgde/hybrid/config.arb \
		-lang=ar,en,zh \
		-output=${dart_test}/pgde/hybrid/archive/%V/%N%E
	@cd ${dart_test}/pgde/hybrid && zip -r archive.zip archive

.PHONY: protoc_hybrid_l10n
protoc_hybrid_l10n:
	@GTT_DIR=${dart_test}/pgde/hybrid protoc -I${makefile_dir} -I${pgv_path} \
		--dart-ext_out=l10n=${dart_test}/pgde/%P.gtt.toml:${dart_test}/pgde \
		${hybrid_path}/config.proto
	@dartfmt -w ${dart_test}/pgde/hybrid/*.l10n.dart


.PHONY: protoc_hybrid_zero
protoc_hybrid_zero:
	@protoc -I${makefile_dir} -I${pgv_path} \
		--dart-ext_out=zero=all:${dart_test}/pgde ${hybrid_path}/config.proto
	@dartfmt -w ${dart_test}/pgde/hybrid/*.zero.dart

.PHONY: protoc_hybrid_validate
protoc_hybrid_validate:
	@protoc -I${makefile_dir} -I${pgv_path} \
		--dart-ext_out=validate:${dart_test}/pgde ${hybrid_path}/config.proto
	@dartfmt -w ${dart_test}/pgde/hybrid/*.validate.dart

.PHONY: protoc_test_x
protoc_test_x:
	protoc -I. \
		--dart-ext_out=arb,l10n=%P.gtt.toml,validate,form:${dart_lib} \
		protos/units/*.proto
	dartfmt -w ${dart_lib}/pgde/**/*.l10n.base.dart

.PHONY: gen_gtt_to_dart
gen_gtt_to_dart:
	@PROJECT_DIR=${makefile_dir} \
		go run ${cmd_path}/gtt_to_dart/*.go \
		-gtt=${test_path}/gtt_to_dart/gtt.toml \
		-dart_out=${dart_lib}/test/app.l10n.dart \
	@dartfmt -w ${dart_lib}/test/app.l10n.dart
