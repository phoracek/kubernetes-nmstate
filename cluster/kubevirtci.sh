export KUBEVIRT_PROVIDER=${KUBEVIRT_PROVIDER:-'k8s-1.17.0'}

KUBEVIRTCI_VERSION='95096c8189c8b620ddc1310e12388df2190e1cc8'
KUBEVIRTCI_REPO='https://github.com/kubevirt/kubevirtci.git'
KUBEVIRTCI_PATH="${PWD}/_kubevirtci"

function kubevirtci::_get_repo() {
    git --git-dir ${KUBEVIRTCI_PATH} remote get-url origin
}

function kubevirtci::_get_version() {
    git --git-dir ${KUBEVIRTCI_PATH}/.git log --format="%H" -n 1
}

function kubevirtci::install() {
    # Remove cloned kubevirtci repository if it does not match the requested one
    if [ -d ${KUBEVIRTCI_PATH} ]; then
        if [ $(kubevirtci::_get_repo) != ${KUBEVIRTCI_REPO} || $(kubevirtci::_get_version) != ${KUBEVIRTCI_VERSION} ]; then
            rm -rf ${KUBEVIRTCI_PATH}
        fi
    fi

    if [ ! -d ${KUBEVIRTCI_PATH} ]; then
        git clone ${KUBEVIRTCI_REPO} ${KUBEVIRTCI_PATH}
        (
            cd ${KUBEVIRTCI_PATH}
            git checkout ${KUBEVIRTCI_VERSION}
        )
    fi
}

function kubevirtci::path() {
    echo -n ${KUBEVIRTCI_PATH}
}

function kubevirtci::kubeconfig() {
    echo -n ${KUBEVIRTCI_PATH}/_ci-configs/${KUBEVIRT_PROVIDER}/.kubeconfig
}
