load("//tools:defs.bzl", "go_library")

package(licenses = ["notice"])

go_library(
    name = "v2",
    srcs = [
        "api.go",
        "epoll.go",
        "options.go",
        "service.go",
        "service_linux.go",
        "state.go",
    ],
    visibility = ["//shim:__subpackages__"],
    deps = [
        "//pkg/cleanup",
        "//pkg/shim/runsc",
        "//pkg/shim/v1/proc",
        "//pkg/shim/v1/utils",
        "//pkg/shim/v2/runtimeoptions",
        "//runsc/specutils",
        "@com_github_burntsushi_toml//:go_default_library",
        "@com_github_containerd_cgroups//:go_default_library",
        "@com_github_containerd_console//:go_default_library",
        "@com_github_containerd_containerd//api/events:go_default_library",
        "@com_github_containerd_containerd//api/types/task:go_default_library",
        "@com_github_containerd_containerd//errdefs:go_default_library",
        "@com_github_containerd_containerd//events:go_default_library",
        "@com_github_containerd_containerd//log:go_default_library",
        "@com_github_containerd_containerd//mount:go_default_library",
        "@com_github_containerd_containerd//namespaces:go_default_library",
        "@com_github_containerd_containerd//pkg/process:go_default_library",
        "@com_github_containerd_containerd//pkg/stdio:go_default_library",
        "@com_github_containerd_containerd//runtime:go_default_library",
        "@com_github_containerd_containerd//runtime/linux/runctypes:go_default_library",
        "@com_github_containerd_containerd//runtime/v2/shim:go_default_library",
        "@com_github_containerd_containerd//runtime/v2/task:go_default_library",
        "@com_github_containerd_containerd//sys/reaper:go_default_library",
        "@com_github_containerd_fifo//:go_default_library",
        "@com_github_containerd_typeurl//:go_default_library",
        "@com_github_gogo_protobuf//types:go_default_library",
        "@com_github_sirupsen_logrus//:go_default_library",
        "@org_golang_x_sys//unix:go_default_library",
    ],
)
