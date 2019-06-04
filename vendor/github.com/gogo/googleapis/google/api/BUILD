load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "annotations.pb.go",
        "http.pb.go",
    ],
    importmap = "github.com/iwaltgen/grpc-gogo-test/vendor/github.com/gogo/googleapis/google/api",
    importpath = "github.com/gogo/googleapis/google/api",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/gogo/protobuf/protoc-gen-gogo/descriptor:go_default_library",
    ],
)
