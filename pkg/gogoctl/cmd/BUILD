package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "root.go",
        "version.go",
    ],
    importpath = "github.com/iwaltgen/grpc-gogo-test/pkg/gogoctl/cmd",
    deps = [
        "@com_github_mitchellh_go_homedir//:go_default_library",
        "@com_github_spf13_cobra//:go_default_library",
        "@com_github_spf13_viper//:go_default_library",
    ],
)
