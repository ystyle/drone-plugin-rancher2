[![Docker Pulls](https://img.shields.io/docker/pulls/ystyle/drone-plugin-rancher2)](https://hub.docker.com/r/ystyle/drone-plugin-rancher2 "view docker")
[![](https://images.microbadger.com/badges/version/ystyle/drone-plugin-rancher2.svg)](https://hub.docker.com/r/ystyle/drone-plugin-rancher2/tags "view tags")

# drone-plugin-rancher2
>NOTE: only can update image name


### Usage
- rancher 2.6.3
  ```yaml
  steps:
    - name: update-rancher
      image: ystyle/drone-plugin-rancher2:2.6
      when:
        branch: master
        event: [push]
      environment:
        SERVER: https://rancher.youdomain.com/k8s/clusters/c-m-w7b37mq1
        TOKEN: kubeconfig-user-d71qwk98dq:d45jfxrdx672r26xmf8ltsj46dcbw2p7lf76dck3vt5nr24gzlpbk9
      commands:
        - /setup.sh
        - kubectl set image deployment/${DRONE_REPO_NAME} ${DRONE_REPO_NAME}=registry.cn-hangzhou.aliyuncs.com/dexdev/${DRONE_REPO_NAME}:${DRONE_COMMIT_BRANCH}-${DRONE_BUILD_NUMBER}

  ```
  >server和token取自集群-右上搜索图标旁边kubectl config的server和token

  ![image](https://user-images.githubusercontent.com/4478635/154631389-70fe7290-5768-4bd1-8476-240330514de9.png)

