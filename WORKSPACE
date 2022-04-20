workspace(name = "goC8")

################################################################################
#  Bazel tools
################################################################################
# These rules are built-into Bazel but we need to load them first to download more rules
# https://docs.bazel.build/versions/master/be/workspace.html#http_archive
# https://docs.bazel.build/versions/master/repo/git.html
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

################################################################################
#  Go rules
################################################################################
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "d6b2513456fe2229811da7eb67a444be7785f5323c6708b38d851d2b51e54d83",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.30.0/rules_go-v0.30.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.30.0/rules_go-v0.30.0.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.18")

################################################################################
#  gazelle
################################################################################
http_archive(
    name = "bazel_gazelle",
    sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

################################################################################
# Protobuf dependencies
# https://github.com/bazelbuild/rules_proto
# https://github.com/protocolbuffers/protobuf/releases
################################################################################
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "rules_proto",
    sha256 = "66bfdf8782796239d3875d37e7de19b1d94301e8972b3cbd2446b332429b4df1",
    strip_prefix = "rules_proto-4.0.0",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

http_archive(
    name = "com_google_protobuf",
    sha256 = "769bb7b97f89fda3d9cadcfd68e6f388d976ba8f0ea9be57601eef839a461898",
    strip_prefix = "protobuf-3.19.3",
    urls = ["https://github.com/protocolbuffers/protobuf/releases/download/v3.19.3/protobuf-all-3.19.3.zip"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

################################################################################
#  Buildifier
#  https://github.com/bazelbuild/buildtools/tree/master/buildifier#readme
################################################################################
http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "ae34c344514e08c23e90da0e2d6cb700fcd28e80c02e23e4d5715dddcb42f7b3",
    strip_prefix = "buildtools-4.2.2",
    urls = [
        "https://github.com/bazelbuild/buildtools/archive/refs/tags/4.2.2.tar.gz",
    ],
)

################################################################################
# golink | https://medium.com/goc0de/a-cute-bazel-proto-hack-for-golang-ides-2a4ef0415a7f
# https://github.com/nikunjy/golink
################################################################################
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "golink",
    sha256 = "ea728cfc9cb6e2ae024e1d5fbff185224592bbd4dad6516f3cc96d5155b69f0d",
    strip_prefix = "golink-1.0.0",
    urls = ["https://github.com/nikunjy/golink/archive/v1.0.0.tar.gz"],
)

################################################################################
# Docker dependencies
# https://github.com/bazelbuild/rules_docker
# https://github.com/bazelbuild/rules_docker/releases
################################################################################
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "59536e6ae64359b716ba9c46c39183403b01eabfbd57578e84398b4829ca499a",
    strip_prefix = "rules_docker-0.22.0",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.22.0/rules_docker-v0.22.0.tar.gz"],
)

################################################################################
# Docker authentication to push into container registry
################################################################################
# Load the macro that allows you to customize the docker toolchain configuration.
load(
    "@io_bazel_rules_docker//toolchains/docker:toolchain.bzl",
    docker_toolchain_configure = "toolchain_configure",
)

# For authentication, run
# gcloud auth login
# gcloud auth configure-docker
# https://cloud.google.com/container-registry/docs/advanced-authentication#gcloud-helper
# docker custom configuration
# https://github.com/bazelbuild/rules_docker#container_push-custom-client-configuration
docker_toolchain_configure(
    name = "docker_config",
    # see https://github.com/bazelbuild/rules_docker
)

load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)

################################################################################
# Containers to pull required to build.
################################################################################
# Digest must be updated in case of a repo change OR a push
# docker pull gcr.io/future-309012/scratch:latest
container_pull(
    name = "scratch",
    # tag = "latest", # Avoid tag as much as possible to ascertain hermetic build
    digest = "sha256:1bbcd9993f5dbfa4f9b84740a37c20d2f98c620caef54d305651abf763367bf2",
    registry = "gcr.io",
    repository = "future-309012/scratch",
)

################################################################################
# Kubernetes rules
# https://github.com/bazelbuild/rules_k8s/releases/
################################################################################
http_archive(
    name = "io_bazel_rules_k8s",
    sha256 = "51f0977294699cd547e139ceff2396c32588575588678d2054da167691a227ef",
    strip_prefix = "rules_k8s-0.6",
    urls = ["https://github.com/bazelbuild/rules_k8s/archive/v0.6.tar.gz"],
)

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories")

k8s_repositories()

load("@io_bazel_rules_k8s//k8s:k8s_go_deps.bzl", k8s_go_deps = "deps")

k8s_go_deps()

################################################################################
#  go_repositories
################################################################################
