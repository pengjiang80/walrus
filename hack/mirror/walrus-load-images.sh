#!/usr/bin/env bash

list="walrus-images.txt"
images="walrus-images.tar.gz"
repository_type="harbor"
harbor_https="true"
harbor_versions="v2.0"
harbor_user=""
harbor_password=""

usage () {
    echo "USAGE: $0 [--images walrus-images.tar.gz] --registry my.registry.com:5000"
    echo "  [-l|--image-list path] text file with list of images; one image per line."
    echo "  [-i|--images path] tar.gz generated by docker save."
    echo "  [-r|--registry registry:port] target private registry in the registry:port format."
    echo "  [--repository-type type] Repository type. enum. e.g \"harbor\". (Default \"\")"
    echo "  [--harbor-https type] Use https by default when create harbor project. enum. e.g \"true\" || \"false\". (Default \"true\")"
    echo "  [--harbor-versions version] Harbor Version. enum. e.g \"v2.0\" || \"v1.0\". (Default \"v2.0\")"
    echo "  [--harbor-user string] Harbor User. (Default \"\")"
    echo "  [--harbor-password string] Harbor Password. (Default \"\")"
    echo "  [-h|--help] Usage message"
}

push_manifest () {
    export DOCKER_CLI_EXPERIMENTAL=enabled
    manifest_list=()
    for i in "${arch_list[@]}"
    do
        manifest_list+=("$1-${i}")
    done

    echo "Preparing manifest $1, list[${arch_list[@]}]"
    docker manifest create "$1" "${manifest_list[@]}" --amend
    docker manifest push "$1" --purge
}

create_harbor_project () {
    if [[ -n "${harbor_user}" && -n "${harbor_password}" ]]; then
        project="sealio"
        echo "Creating harbor project ${project}"
        url="${target_registry}/api/v2.0/projects"
        if [[ "${harbor_versions}" == "v1.0" ]]; then
            url="${target_registry}/api/projects"
        fi
        if [[ "${harbor_https}" == "true" ]]; then
            url="https://${url}"
        fi
        result_code=$(curl -k -s -u "${harbor_user}:${harbor_password}" -X POST -H "Content-type:application/json" -d '{"project_name":"'"${project}"'","metadata":{"public":"true"}}' $url)

        if [[ -z "${result_code}" ]]; then
            echo "Created porject ${project}."
        elif [[ "${result_code}" =~ "already exists" || "${result_code}" =~ "conflict project" ]]; then
            echo "Porject ${project} already exists."
        else
            echo "Failed to create Porject ${project}."
        fi
    fi
}

while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
        -r|--registry)
        target_registry="$2"
        shift # past argument
        shift # past value
        ;;
        -l|--image-list)
        list="$2"
        shift # past argument
        shift # past value
        ;;
        -i|--images)
        images="$2"
        shift # past argument
        shift # past value
        ;;
        --windows-image-list)
        windows_image_list="$2"
        shift # past argument
        shift # past value
        ;;
        --windows-versions)
        windows_versions="$2"
        shift # past argument
        shift # past value
        ;;
        --harbor-versions)
        harbor_versions="$2"
        shift # past argument
        shift # past value
        ;;
        --harbor-user)
        harbor_user="$2"
        shift # past argument
        shift # past value
        ;;
        --harbor-password)
        harbor_password="$2"
        shift # past argument
        shift # past value
        ;;
        --repository-type)
        repository_type="$2"
        shift # past argument
        shift # past value
        ;;
         --harbor-https)
        harbor_https="$2"
        shift # past argument
        shift # past value
        ;;
        -h|--help)
        help="true"
        shift
        ;;
        *)
        usage
        exit 1
        ;;
    esac
done
if [[ -z "${target_registry}" ]]; then
    usage
    exit 1
fi
if [[ $help ]]; then
    usage
    exit 0
fi

#docker load --input ${images}

linux_images=()
linux_err_images=()
while IFS= read -r i; do
    [ -z "${i}" ] && continue
    linux_images+=("${i}");
done < "${list}"

if [[ "${repository_type}" == "harbor" ]]; then
    create_harbor_project
elif [[ -n "${repository_type}" && "${repository_type}" != "harbor" ]]; then
    echo "Repository type error"
fi

arch_list=()
arch_list+=("linux-amd64")
for i in "${linux_images[@]}"; do
    [ -z "${i}" ] && continue
    target_image_name="${i#docker.io/}"
    arch_suffix=""
    use_manifest=false

    case $target_image_name in
    */*)
        image_name="${target_registry}/${target_image_name}"
        ;;
    *)
        image_name="${target_registry}/sealio/${target_image_name}"
        ;;
    esac

    docker tag "${i}" "${image_name}${arch_suffix}"
    docker push "${image_name}${arch_suffix}"

    if [ $? -ne 0 ]; then
        linux_err_images+=("${image_name}");
    fi

    if $use_manifest; then
        push_manifest "${image_name}"
    fi
done

if [[ "${#linux_err_images[@]}" -gt 0 ]]; then
    echo "Some images failed to upload. Generate failed-upload-images.txt to log the failed images"
    if [ -f "./failed-upload-images.txt" ]; then
        rm -rf ./failed-upload-images.txt
    fi
    touch "failed-upload-images.txt"

    for i in "${linux_err_images[@]}"; do
        echo "docker push ${i} error"
        echo "${i}" >> "failed-upload-images.txt"
    done
else
    echo "Successfully push all images"
    if [ -f "./failed-upload-images.txt" ]; then
        rm -rf ./failed-upload-images.txt
    fi
fi
