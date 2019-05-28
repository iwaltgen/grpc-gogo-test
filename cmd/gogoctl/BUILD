load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/iwaltgen/grpc-gogo-test/cmd/gogoctl",
    visibility = ["//visibility:private"],
    deps = ["//pkg/gogoctl/cmd:go_default_library"],
)

go_binary(
    name = "gogoctl",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
