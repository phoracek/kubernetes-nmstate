FROM fedora:32

RUN dnf install -y dnf-plugins-core && \
    dnf copr enable -y nmstate/nmstate-0.2 && \
    dnf install -y nmstate iproute iputils && \
    dnf remove -y dnf-plugins-core && \
    dnf clean all

COPY build/_output/bin/kubernetes-nmstate /usr/local/bin/kubernetes-nmstate
COPY build/bin /usr/local/bin

ENTRYPOINT ["kubernetes-nmstate"]
