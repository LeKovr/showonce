# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [1.2.1] - 2025-03-12

* build for go 1.23.6
* update go-kit/config v0.5.1 features
* use github.com/dopos/narra v0.27.1
* update other deps

## [1.2.0] - 2025-02-21

### Devops

* use github.com/LeKovr/go-kit/config v0.5.1 (new config arch)
* update linters

### Bump build(deps)

* google.golang.org/protobuf from 1.36.4 to 1.36.5 (#89)
* github.com/grpc-ecosystem/grpc-gateway/v2 (#88)

## [1.1.7] - 2025-01-27

### Bump build(deps)

* google.golang.org/protobuf from 1.36.1 to 1.36.4 (#87)
* google.golang.org/grpc from 1.69.2 to 1.70.0 (#86)
* github.com/grpc-ecosystem/grpc-gateway/v2 (#85)

## [1.1.6] - 2025-01-04

* add: --version support
* html linting
* golang 1.22.3
* golangci v1.60.3
* upd: codeql v3
* upd: activate gosec
* upd: github workflows
* fix: linter notices
* fix: match FROM & AS case

### Bump build(deps)

* actions/checkout from 3 to 4
* actions/setup-go from 3 to 5
* docker/build-push-action from 4 to 6
* docker/login-action from 2 to 3
* docker/metadata-action from 4 to 5
* docker/setup-buildx-action from 2 to 3
* docker/setup-qemu-action from 2 to 3
* github.com/dopos/narra from 0.26.4 to 0.26.6
* github.com/grpc-ecosystem/grpc-gateway/v2
* github.com/stretchr/testify from 1.9.0 to 1.10.0
* google.golang.org/grpc from 1.67.1 to 1.68.0
* google.golang.org/grpc from 1.68.0 to 1.69.2
* google.golang.org/protobuf from 1.35.1 to 1.35.2
* google.golang.org/protobuf from 1.35.2 to 1.36.1

## [1.1.5] - 2024-11-08

* add: confirmation dialog
* fix: increase bottom padding
* upd: bump deps

## [1.1.4] - 2024-06-01

* update golang-ci version

## [1.1.3] - 2024-06-01

### Changed

* go 1.22.3
* closes #29, add HTTP_PROTO to .env
* build(deps): bump github.com/alecthomas/assert/v2 from 2.9.0 to 2.10.0
* build(deps): bump github.com/grpc-ecosystem/grpc-gateway/v2 to v2.20.0

## [1.1.2] - 2024-05-15

### Bump build(deps)

* github.com/LeKovr/go-kit/server v0.12.24 -> v0.13.0
* github.com/LeKovr/go-kit/slogger v0.12.24 -> v0.13.0
* google.golang.org/genproto/googleapis/api v0.0.0-20240125205218-1f4bbc51befe -> v0.0.0-20240227224415-6ceb2ff114de
* google.golang.org/grpc v1.62.1 -> v1.63.2
* google.golang.org/protobuf v1.33.0 -> v1.34.1

## [1.1.1] - 2024-03-21

### Changed

* add robots.txt
* js deps fix: Inclusion of functionality from an untrusted source
* build(deps): bump google.golang.org/protobuf from 1.32.0 to 1.33.0
* build(deps): bump github.com/stretchr/testify from 1.8.4 to 1.9.0
* build(deps): bump github.com/alecthomas/assert/v2 from 2.5.0 to 2.6.0
* build(deps): bump google.golang.org/grpc from 1.61.0 to 1.62.1
* bump go-kit/config v0.2.2

## [1.1.0] - 2024-01-02

### Changed

* move to [go-kit/server](https://github.com/LeKovr/go-kit/tree/main/server), [go-kit/slogger](https://github.com/LeKovr/go-kit/tree/main/slogger)
* fix test script
* move static files to subdir

## [1.0.7] - 2023-12-23

### Changed

* add: calc app version on static pages
* fix: static build

## [1.0.6] - 2023-12-22

### Changed

* move old changelog to CHANGELOG.md
* add docker image for linux/arm64
* update dcape v3 support
* update go ver to 1.21
* resolve linter warnings

### Dependencies updated

* build(deps): bump github.com/felixge/httpsnoop from 1.0.3 to 1.0.4
* build(deps): bump golang.org/x/sync from 0.4.0 to 0.5.0
* build(deps): bump google.golang.org/grpc from 1.59.0 to 1.60.0
* build(deps): bump github.com/dopos/narra from 0.26.0 to 0.26.1
* build(deps): bump github.com/grpc-ecosystem/grpc-gateway/v2 to v2.18.1

## [1.0.5] - 2023-11-03

### Changed

* fix: Dependabot alerts #4, #5 (bump golang.org/x/net from v0.15.0 to v0.17.0)
* fix: add build check for go.*
* fix: Status READ replace with EXPIRED

### Dependencies updated

* build(deps): bump github.com/go-logr/logr from 1.2.4 to 1.3.0

## [1.0.4] - 2023-10-29

### Dependencies updated

* build(deps): bump google.golang.org/grpc from 1.57.0 to 1.59.0
* build(deps): bump golang.org/x/sync from 0.3.0 to 0.4.0
* build(deps): bump github.com/grpc-ecosystem/grpc-gateway/v2

## [1.0.3] - 2023-08-02

### Added

* static: use chota.css
* code: use github.com/LeKovr/go-kit/ver
* code: use MDUserKey
* code: improve names,docs,comments,tests
* README: cloc & help parts
* README: add mermaid chart, collapse chart legend
* Makefile: cloc target

### Changed

* code: rename storage.Iface -> StorageIface
* CI/CD: replace .drone.yml with .woodpecker.yml

## [1.0.0] - 2023-07-09

*  рефакторинг с изменением архитектуры на GRPC

## [0.2.0] - 2022-12-09

* незначительный рефакторинг
* предварительная доработка стилей
* переезд на dopos/narra

## [0.1.1] - 2022-07-07

* релиз, доработан ввод срока жизни, добавлен деплой в dcape

## [0.0.2] - 2022-07-06

* MVP, начало тестирования

## [0.0.1] - 2022-06-29

* начало работ, предварительный вариант ТЗ
