[![Docker Pulls](https://img.shields.io/docker/pulls/ystyle/drone-plugin-rancher2)](https://hub.docker.com/r/ystyle/drone-plugin-rancher2 "view docker")
[![](https://images.microbadger.com/badges/version/ystyle/drone-plugin-rancher2.svg)](https://hub.docker.com/r/ystyle/drone-plugin-rancher2/tags "view tags")

# drone-plugin-rancher2
>NOTE: only can update image name

### Usage
- rancher 2.6.3 以上
  ```yaml
  steps:
    - name: deploy-uat
      image: ystyle/drone-plugin-rancher2:2.6
      when:
        branch: master
        event: [push]
      environment:
        SERVER:
          from_secret: SERVER
        TOKEN:
          from_secret: TOKEN
      commands:
        - /setup.sh
        - kubectl set image deployment/${DRONE_REPO_NAME} ${DRONE_REPO_NAME}=registry.cn-hangzhou.aliyuncs.com/dexdev/${DRONE_REPO_NAME}:${DRONE_COMMIT_BRANCH}-${DRONE_BUILD_NUMBER}

  ```
  >server和token取自集群-右上搜索图标旁边kubectl config的server和token
  ![image](https://user-images.githubusercontent.com/4478635/154631389-70fe7290-5768-4bd1-8476-240330514de9.png)w



- drone 1.x
```yml
steps:
  - name: update-rancher
    image: ystyle/drone-plugin-rancher2
    settings:
      api: https://rancher.youdomain.com/v3/project/c-27tt1:p-wqtk3/workloads/default:default:nginx
      access_key: token-dfi8s
      secret_key: mhgb9w2csrvs6h9nm4fldl4z8f7h4sznx4m6ggvv4p54sffhjlkwvx
      data: [{
        "name": "nginx",
        "image": "nginx:alpine",
        "environment": {
          "tag": "${DRONE_COMMIT_SHA:0:8}"
          }
      }]
```


- drone 0.8
```yml
pipeline:
  update-rancher:
    image: ystyle/drone-plugin-rancher2
    api: https://rancher.youdomain.com/v3/project/c-27tt1:p-wqtk3/workloads/default:default:nginx
    access_key: token-dfi8s
    secret_key: mhgb9w2csrvs6h9nm4fldl4z8f7h4sznx4m6ggvv4p54sffhjlkwvx
    data: [{
      "name": "nginx",
      "image": "nginx:alpine",
      "environment": {
        "tag": "${DRONE_COMMIT_SHA:0:8}"
        }
      }]
```

### Parameter Reference
- `api` workload api url
- `access_key` rancher user api keys
- `secret_key` rancher user api keys
- `data` api result's containers field, only support `name` `image` and `environment` field
