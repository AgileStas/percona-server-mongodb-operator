FROM goboring/golang:1.16.7b7 AS go_builder
WORKDIR /go/src/github.com/percona/percona-server-mongodb-operator

COPY . .

ARG GIT_COMMIT
ARG GIT_BRANCH
ARG GO_LDFLAGS
ARG GOOS=linux
ARG GOARCH=amd64
ARG CGO_ENABLED=0

RUN ls -la && mkdir -p build/_output/bin \
    && GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=$CGO_ENABLED GO_LDFLAGS=$GO_LDFLAGS \
       go build -mod=vendor -ldflags "-w -s -X main.GitCommit=$GIT_COMMIT -X main.GitBranch=$GIT_BRANCH" \
           -o build/_output/bin/percona-server-mongodb-operator \
               cmd/manager/main.go \
    && cp -r build/_output/bin/percona-server-mongodb-operator /usr/local/bin/percona-server-mongodb-operator \
    && GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=$CGO_ENABLED GO_LDFLAGS=$GO_LDFLAGS \
       go build -mod=vendor -ldflags "-w -s -X main.GitCommit=$GIT_COMMIT -X main.GitBranch=$GIT_BRANCH" \
           -o build/_output/bin/mongodb-healthcheck \
               cmd/mongodb-healthcheck/main.go \
    && cp -r build/_output/bin/mongodb-healthcheck /usr/local/bin/mongodb-healthcheck

# Looking for all possible License/Notice files and copying them to the image
RUN find vendor -regextype posix-extended -iregex '.*(license|notice)(\.md|\.txt|$)' \
         -exec \
            sh -c 'mkdir -pv /licenses/$(echo "$0" | sed -E "s/\/(license|notice).*$//gi") \
                   && cp -v "$0" /licenses/$(echo "$0" | sed -E "s/\/(license|notice).*$//gi")' {} \;

FROM registry.access.redhat.com/ubi7/ubi-minimal AS ubi7
# sqlite, libsolv and libmodulemd are from the base system (needed by microdnf itself)
#RUN microdnf update && microdnf install -y dracut-fips && microdnf clean all
RUN microdnf update \
    && export GNUPGHOME="$(mktemp -d)" \
    && gpg --batch --keyserver keyserver.ubuntu.com --recv-keys 6341AB2753D78A78A7C27BB124C6A8A7F4A80EB5 \
    && gpg --batch --export --armor 6341AB2753D78A78A7C27BB124C6A8A7F4A80EB5 > ${GNUPGHOME}/RPM-GPG-KEY-CentOS-7 \
    && rpmkeys --import ${GNUPGHOME}/RPM-GPG-KEY-CentOS-7 \
    && microdnf -y install dracut \
    && curl -o /tmp/hmaccalc.rpm http://mirror.centos.org/centos/7/os/x86_64/Packages/hmaccalc-0.9.13-4.el7.x86_64.rpm \
    && curl -o /tmp/dracut-fips.rpm http://mirror.centos.org/centos/7/os/x86_64/Packages/dracut-fips-033-572.el7.x86_64.rpm \
    && rpmkeys --checksig /tmp/hmaccalc.rpm /tmp/dracut-fips.rpm \
    && rpm -i /tmp/hmaccalc.rpm /tmp/dracut-fips.rpm \
    && rm -rf /tmp/hmaccalc.rpm /tmp/dracut-fips.rpm \
    && microdnf clean all

LABEL name="Percona Server for MongoDB Operator" \
      vendor="Percona" \
      summary="Percona Server for MongoDB Operator simplifies the deployment and management of Percona Server for MongoDB in a Kubernetes or OpenShift environment" \
      description="Percona Server for MongoDB is our free and open-source drop-in replacement for MongoDB Community Edition. It offers all the features and benefits of MongoDB Community Edition, plus additional enterprise-grade functionality." \
      maintainer="Percona Development <info@percona.com>"

COPY --from=go_builder /usr/local/bin/percona-server-mongodb-operator /usr/local/bin/percona-server-mongodb-operator
COPY --from=go_builder /licenses/* /licenses
COPY LICENSE /licenses/
COPY --from=go_builder /usr/local/bin/mongodb-healthcheck /mongodb-healthcheck
COPY build/init-entrypoint.sh /init-entrypoint.sh
COPY build/ps-entry.sh /ps-entry.sh
# Mozilla licensed source code should be included
COPY vendor/github.com/hashicorp /lib

USER 2
