package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["greeter.pb.go"],
    importpath = "github.com/iwaltgen/grpc-gogo-test/proto",
    deps = [
        "//vendor/github.com/gogo/googleapis/google/api:go_default_library",
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/gogo/protobuf/types:go_default_library",
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
    ],
)
