load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["server.go"],
    importpath = "github.com/iwaltgen/grpc-gogo-test/pkg/gogoserver/grpc",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/gogoserver/grpc/service:go_default_library",
        "//proto/v1:go_default_library",
        "//vendor/github.com/grpc-ecosystem/go-grpc-middleware:go_default_library",
        "//vendor/github.com/grpc-ecosystem/go-grpc-middleware/recovery:go_default_library",
        "//vendor/github.com/grpc-ecosystem/go-grpc-middleware/validator:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
        "//vendor/google.golang.org/grpc/status:go_default_library",
    ],
)
