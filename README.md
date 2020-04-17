# grpc-gogo-test
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fiwaltgen%2Fgrpc-gogo-test.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fiwaltgen%2Fgrpc-gogo-test?ref=badge_shield)


## /cmd

모든 실행 가능한 파일

## /pkg

모든 로직 라이브러리

## /proto

프로토콜 정의

## build protocol (with protobuf)

프로토콜이 변경 될경우만 한번 실행

```sh
./gen_proto.sh
```

## build all (with bazel)

```sh
bazel run //:gazelle -- update-repos -from_file=go.mod # go 의존성 변경시 마다 실행 필요
bazel build //...
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fiwaltgen%2Fgrpc-gogo-test.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fiwaltgen%2Fgrpc-gogo-test?ref=badge_large)