package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "greeter.go",
        "root.go",
        "vendor.go",
        "version.go",
    ],
    importpath = "github.com/iwaltgen/grpc-gogo-test/pkg/gogoctl/cmd",
    deps = [
        "//proto:go_default_library",
        "//vendor/github.com/gogo/googleapis/google/api:go_default_library",
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/types:go_default_library",
        "//vendor/github.com/mitchellh/go-homedir:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/github.com/spf13/viper:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
    ],
)
