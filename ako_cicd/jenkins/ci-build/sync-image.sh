#!/bin/bash

set -xe

IFS=', ' read -r -a registries <<< "$REMOTE_DOCKER_REGISTRIES"

if [ "${#registries[@]}" == "0" ]; then
	echo "Remote registries count must be non-zero"
    exit 1
fi

version_tag=$($WORKSPACE/hack/jenkins/get_build_version.sh $JOB_NAME $build_num)



VENV_PATH=$HOME/venv
if [ ! -d "$VENV_PATH" ]; then
	virtualenv -p python3 $VENV_PATH
fi

source $VENV_PATH/bin/activate
pip install sshuttle

set +e
pgrep -f sshuttle | xargs -r sudo kill -9



ps -ef | grep sshuttle | grep -v grep
rc=$?

set -e

if [ "$rc" != "0" ]; then
    sshuttle -D -r aviuser@35.209.243.149 10.128.0.0/11 -e "ssh -i  ~/.ssh/id_jumphost_vijay_rsa" -v
fi

###########

source_image=$PVT_DOCKER_REGISTRY/$PVT_DOCKER_REPOSITORY/$DOCKER_AKO_IMAGE_NAME:$version_tag

sudo docker pull $source_image

for registry in "${registries[@]}"
do
    target_image="$registry/$PVT_DOCKER_REPOSITORY/$DOCKER_AKO_IMAGE_NAME:$version_tag"
    echo "Tagging and pushing to registry: $registry"
    sudo docker tag $source_image $target_image
    sudo docker push $target_image
done

##########

source_image=$PVT_DOCKER_REGISTRY/$PVT_DOCKER_REPOSITORY/$DOCKER_AKO_OPERATOR_IMAGE_NAME:$version_tag

sudo docker pull $source_image

for registry in "${registries[@]}"
do
    target_image="$registry/$PVT_DOCKER_REPOSITORY/$DOCKER_AKO_OPERATOR_IMAGE_NAME:$version_tag"
    echo "Tagging and pushing to registry: $registry"
    sudo docker tag $source_image $target_image
    sudo docker push $target_image
done

###########

source_image=$PVT_DOCKER_REGISTRY/$PVT_DOCKER_REPOSITORY/$DOCKER_AKO_GATEWAY_API_IMAGE_NAME:$version_tag

sudo docker pull $source_image

for registry in "${registries[@]}"
do
    target_image="$registry/$PVT_DOCKER_REPOSITORY/$DOCKER_AKO_GATEWAY_API_IMAGE_NAME:$version_tag"
    echo "Tagging and pushing to registry: $registry"
    sudo docker tag $source_image $target_image
    sudo docker push $target_image
done
