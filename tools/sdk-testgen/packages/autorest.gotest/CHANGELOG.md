# Change Log - @autorest/gotest

This log was last generated on Wed, 12 Jan 2022 09:10:46 GMT and should not be manually modified.

## 1.3.0
Wed, 12 Jan 2022 09:10:46 GMT

### Minor changes

- use new api scenario through testmodeler

## 1.2.0
Wed, 12 Jan 2022 02:19:25 GMT

### Minor changes

- Compatible with latest azcore and azidentity.
- Add response check to mock test generation.

### Patches

- Fix result check problem for lro operation with pageable config.
- Fix result log problem for multiple response operation.
- Fix wrong param name for pageable opeation with custom item name.
- Different conversion for choice and sealedchoice.
- Fix wrong generation of null value for object.
- Fix some generated problems including: polymorphism response type, client param, pager response check.
- Fix multiple time format and any-object default value issue.
- Refine log for mock test and fix array item code generate bug.
- Upgrade to latest autorest/core and autorest/go.

## 1.1.3
Mon, 29 Nov 2021 06:10:09 GMT

### Patches

- Replace incomplete response check with just log temporarily.

## 1.1.2
Mon, 15 Nov 2021 09:39:03 GMT

### Patches

- Fix some generation corner case.

## 1.1.1
Tue, 09 Nov 2021 10:20:51 GMT

### Patches

- Remove `go mod tidy` process.

## 1.1.0
Tue, 09 Nov 2021 09:17:24 GMT

### Minor changes

- Refactor structure and fix most of generation problem.

## 1.0.0
Mon, 01 Nov 2021 09:01:05 GMT

### Breaking changes

- Init public version of autorest extension for GO test generation

