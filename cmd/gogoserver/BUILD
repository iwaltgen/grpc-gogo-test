load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/iwaltgen/grpc-gogo-test/cmd/gogoserver",
    visibility = ["//visibility:private"],
    deps = ["//pkg/gogoserver/grpc:go_default_library"],
)

go_binary(
    name = "gogoserver",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
