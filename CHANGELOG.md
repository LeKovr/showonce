# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

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
* build(deps): bump github.com/grpc-ecosystem/grpc-gateway/v2
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
