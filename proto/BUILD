package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["greeter.pb.go"],
    importpath = "github.com/iwaltgen/grpc-gogo-test/proto",
    deps = [
        "@com_github_gogo_protobuf//proto:go_default_library",
        "@com_github_gogo_protobuf//types:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
