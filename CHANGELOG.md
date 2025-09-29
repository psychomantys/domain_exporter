# Changelog

## [0.3.0](https://github.com/psychomantys/domain_exporter/compare/v0.2.1...v0.3.0) (2025-07-29)


### Features

* **app:** Added ENV VAR For config file to enable new helm chart ([1693864](https://github.com/psychomantys/domain_exporter/commit/1693864295f80d3b22b7376529e7949fed622031))
* **chart:** Add auto generated schema ([e9615ca](https://github.com/psychomantys/domain_exporter/commit/e9615cad3d191730f845052c57908ee97803bd98))
* **chart:** Start of alert rules for prometheus generation ([1693864](https://github.com/psychomantys/domain_exporter/commit/1693864295f80d3b22b7376529e7949fed622031))
* **chart:** Start of helm chart ([1693864](https://github.com/psychomantys/domain_exporter/commit/1693864295f80d3b22b7376529e7949fed622031))
* **chart:** Start of mixin for grafana dashboard generation ([1693864](https://github.com/psychomantys/domain_exporter/commit/1693864295f80d3b22b7376529e7949fed622031))
* stop yaml truncating trailing 0 in float ([d94f6b2](https://github.com/psychomantys/domain_exporter/commit/d94f6b22223535464b2b8c6d0446f9a1b24e4633))
* using correct kingpin module path ([3c440d2](https://github.com/psychomantys/domain_exporter/commit/3c440d2ef68e3edd4ae164afbae70bbb1e55d87b))


### Bug Fixes

* .cn ([2bf1f73](https://github.com/psychomantys/domain_exporter/commit/2bf1f73829c01bcec915fb522cd40202f010beb6))
* **action:** ensure the directory exists ([45223e4](https://github.com/psychomantys/domain_exporter/commit/45223e46cc51420c13b43b0d1658c1d672627854))
* **action:** quote e-mail address ([8816843](https://github.com/psychomantys/domain_exporter/commit/881684334b9732b25fcb8bb510ccfe241be39863))
* **action:** quotes ([b18565c](https://github.com/psychomantys/domain_exporter/commit/b18565c1e50e2abd0c307cef97bd24b484533061))
* **action:** require to cd into the correct directory first ([1792098](https://github.com/psychomantys/domain_exporter/commit/179209869eb3483b948a7a38cdd29803d822d826))
* **actions:** Use the correct URL for the readme cri label ([54af5bf](https://github.com/psychomantys/domain_exporter/commit/54af5bf69d1a869c4fc7e78739989557ae2018d3))
* **action:** tweaking pgp key loading ([a768602](https://github.com/psychomantys/domain_exporter/commit/a768602039122fed877215f31cedd8b5dbf30c04))
* **action:** typo ([143d355](https://github.com/psychomantys/domain_exporter/commit/143d35518f9b928db3b951a9281d2dd5d153be04))
* **chart:** comment in wrong location ([5b7aa8e](https://github.com/psychomantys/domain_exporter/commit/5b7aa8e13053fa257687436d2c19a1c7f3eeb453))
* **chart:** Config volume mount was causing issues on some CRIs, it now has its own directory. ([487d4ab](https://github.com/psychomantys/domain_exporter/commit/487d4ab4b5b824945077b432b0ce407419d857dc))
* **chart:** Correctly define the probes ([4d53589](https://github.com/psychomantys/domain_exporter/commit/4d535892a2832eb8f901358bf0032955f76223d4))
* **chart:** Fixes dashboards generation path and bump chart version ([26a14e3](https://github.com/psychomantys/domain_exporter/commit/26a14e322f0ede37a5282b59e1318993f684def5))
* **deps:** go mod tidy ([8a01be0](https://github.com/psychomantys/domain_exporter/commit/8a01be08782722c2448abe53469e16afc0dbcef8))
* **deps:** update module github.com/go-kit/kit to v0.13.0 ([#119](https://github.com/psychomantys/domain_exporter/issues/119)) ([5ddb6f5](https://github.com/psychomantys/domain_exporter/commit/5ddb6f5be8e86c0bf1658927b993623b9159fe42))
* **deps:** update module github.com/prometheus/client_golang to v1.12.2 ([f92ad40](https://github.com/psychomantys/domain_exporter/commit/f92ad4019f57c2786c626ecbbf832a39cf6d5704))
* **deps:** update module github.com/prometheus/client_golang to v1.13.0 ([fca0b6f](https://github.com/psychomantys/domain_exporter/commit/fca0b6fa09f34a513248e07ca52cff2bb9175ae8))
* **deps:** update module github.com/prometheus/client_golang to v1.13.1 ([3780839](https://github.com/psychomantys/domain_exporter/commit/37808393864f9a302e032374d829739d9d2383d7))
* **deps:** update module github.com/prometheus/client_golang to v1.14.0 ([5102f0b](https://github.com/psychomantys/domain_exporter/commit/5102f0b557f92c95efc7f4eb95cd7cf68ea9f38e))
* **deps:** update module github.com/prometheus/common to v0.35.0 ([257f8eb](https://github.com/psychomantys/domain_exporter/commit/257f8eb3fcc76e3a44e55e20b1cb2fb3cc46dc2d))
* **deps:** update module github.com/prometheus/common to v0.36.0 ([16f08fe](https://github.com/psychomantys/domain_exporter/commit/16f08fe8b491d1c38ef8b8318cb64761e2d1889d))
* **deps:** update module github.com/prometheus/common to v0.37.0 ([e086b71](https://github.com/psychomantys/domain_exporter/commit/e086b71c4a27dd8398a86fe45d5eb0ffc05aee13))
* **deps:** update module github.com/prometheus/common to v0.38.0 ([d6f4b18](https://github.com/psychomantys/domain_exporter/commit/d6f4b18096d196effbf531820b3e6ccd754b0130))
* **deps:** update module github.com/prometheus/common to v0.39.0 ([c8dbddf](https://github.com/psychomantys/domain_exporter/commit/c8dbddfcc3df9d448d6f1dbf7f72084c13228bbb))
* **deps:** update module github.com/prometheus/common to v0.40.0 ([15326d0](https://github.com/psychomantys/domain_exporter/commit/15326d06788438fe63dbbdfc91971223cf04c116))
* **deps:** update module github.com/prometheus/common to v0.43.0 ([b028c12](https://github.com/psychomantys/domain_exporter/commit/b028c129efdf694a1088e101c8ac30bd9c5513cd))
* **deps:** update module gopkg.in/yaml.v2 to v3 ([2e04e3b](https://github.com/psychomantys/domain_exporter/commit/2e04e3be1dd2ad0f494951b7e34c869b20d39346))

## [0.2.1](https://github.com/shift/domain_exporter/compare/v0.2.0...v0.2.1) (2025-06-13)


### Bug Fixes

* **deps:** update module github.com/go-kit/kit to v0.13.0 ([#119](https://github.com/shift/domain_exporter/issues/119)) ([5ddb6f5](https://github.com/shift/domain_exporter/commit/5ddb6f5be8e86c0bf1658927b993623b9159fe42))

## [0.2.0](https://github.com/shift/domain_exporter/compare/v0.1.25...v0.2.0) (2023-09-04)


### Features

* stop yaml truncating trailing 0 in float ([d94f6b2](https://github.com/shift/domain_exporter/commit/d94f6b22223535464b2b8c6d0446f9a1b24e4633))
* using correct kingpin module path ([3c440d2](https://github.com/shift/domain_exporter/commit/3c440d2ef68e3edd4ae164afbae70bbb1e55d87b))


### Bug Fixes

* **deps:** update module github.com/prometheus/common to v0.43.0 ([b028c12](https://github.com/shift/domain_exporter/commit/b028c129efdf694a1088e101c8ac30bd9c5513cd))

## [0.1.25](https://github.com/shift/domain_exporter/compare/v0.1.24...v0.1.25) (2023-02-20)


### Bug Fixes

* **deps:** update module github.com/prometheus/common to v0.40.0 ([15326d0](https://github.com/shift/domain_exporter/commit/15326d06788438fe63dbbdfc91971223cf04c116))

## [0.1.24](https://github.com/shift/domain_exporter/compare/v0.1.23...v0.1.24) (2022-12-20)


### Bug Fixes

* **deps:** update module gopkg.in/yaml.v2 to v3 ([2e04e3b](https://github.com/shift/domain_exporter/commit/2e04e3be1dd2ad0f494951b7e34c869b20d39346))

## [0.1.23](https://github.com/shift/domain_exporter/compare/v0.1.22...v0.1.23) (2022-12-15)


### Bug Fixes

* **deps:** update module github.com/prometheus/common to v0.39.0 ([c8dbddf](https://github.com/shift/domain_exporter/commit/c8dbddfcc3df9d448d6f1dbf7f72084c13228bbb))

## [0.1.22](https://github.com/shift/domain_exporter/compare/v0.1.21...v0.1.22) (2022-12-08)


### Bug Fixes

* **deps:** update module github.com/prometheus/common to v0.38.0 ([d6f4b18](https://github.com/shift/domain_exporter/commit/d6f4b18096d196effbf531820b3e6ccd754b0130))

## [0.1.21](https://github.com/shift/domain_exporter/compare/v0.1.20...v0.1.21) (2022-11-08)


### Bug Fixes

* **deps:** update module github.com/prometheus/client_golang to v1.14.0 ([5102f0b](https://github.com/shift/domain_exporter/commit/5102f0b557f92c95efc7f4eb95cd7cf68ea9f38e))

## [0.1.20](https://github.com/shift/domain_exporter/compare/v0.1.19...v0.1.20) (2022-11-02)


### Bug Fixes

* **deps:** update module github.com/prometheus/client_golang to v1.13.1 ([3780839](https://github.com/shift/domain_exporter/commit/37808393864f9a302e032374d829739d9d2383d7))

## [0.1.19](https://github.com/shift/domain_exporter/compare/v0.1.18...v0.1.19) (2022-08-06)


### Bug Fixes

* **deps:** update module github.com/prometheus/client_golang to v1.13.0 ([fca0b6f](https://github.com/shift/domain_exporter/commit/fca0b6fa09f34a513248e07ca52cff2bb9175ae8))

## [0.1.18](https://github.com/shift/domain_exporter/compare/v0.1.17...v0.1.18) (2022-08-01)


### Bug Fixes

* **deps:** update module github.com/prometheus/common to v0.37.0 ([e086b71](https://github.com/shift/domain_exporter/commit/e086b71c4a27dd8398a86fe45d5eb0ffc05aee13))

## [0.1.17](https://github.com/shift/domain_exporter/compare/v0.1.16...v0.1.17) (2022-07-21)


### Bug Fixes

* **deps:** go mod tidy ([8a01be0](https://github.com/shift/domain_exporter/commit/8a01be08782722c2448abe53469e16afc0dbcef8))

## [0.1.16](https://github.com/shift/domain_exporter/compare/v0.1.15...v0.1.16) (2022-07-12)


### Bug Fixes

* **deps:** update module github.com/prometheus/common to v0.35.0 ([257f8eb](https://github.com/shift/domain_exporter/commit/257f8eb3fcc76e3a44e55e20b1cb2fb3cc46dc2d))
* **deps:** update module github.com/prometheus/common to v0.36.0 ([16f08fe](https://github.com/shift/domain_exporter/commit/16f08fe8b491d1c38ef8b8318cb64761e2d1889d))

## [0.1.15](https://github.com/shift/domain_exporter/compare/v0.1.14...v0.1.15) (2022-06-16)


### Bug Fixes

* **deps:** update module github.com/prometheus/client_golang to v1.12.2 ([f92ad40](https://github.com/shift/domain_exporter/commit/f92ad4019f57c2786c626ecbbf832a39cf6d5704))
