package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "gogotest_proto",
    srcs = ["greeter.proto"],
    deps = [
        "@com_google_protobuf//:empty_proto",
        "@com_google_protobuf//:timestamp_proto",
    ],
)

go_proto_library(
    name = "gogotest_go_proto",
    compilers = ["@io_bazel_rules_go//proto:gofast_grpc"],
    importpath = "github.com/iwaltgen/grpc-gogo-test/proto",
    proto = ":gogotest_proto",
)

go_library(
    name = "go_default_library",
    embed = [":gogotest_go_proto"],
)
