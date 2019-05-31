package(default_visibility = ["//visibility:public"])

load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_docker//container:container.bzl", "container_image")

# gazelle:prefix github.com/iwaltgen/grpc-gogo-test
# gazelle:build_file_name BUILD,BUILD.bazel
# gazelle:proto disable_global
gazelle(name = "gazelle")

container_image(
    name = "genproto",
    base = "@prototool//image",
    cmd = ["prototool", "generate"],
    legacy_run_behavior = False,
    docker_run_flags = "--rm -v $(pwd):/work",
)
