![Docker Pulls](https://img.shields.io/docker/pulls/ystyle/drone-plugin-rancher2)
[![](https://images.microbadger.com/badges/version/ystyle/drone-plugin-rancher2.svg)](https://microbadger.com/images/ystyle/drone-plugin-rancher2 "Get your own version badge on microbadger.com")
[![](https://images.microbadger.com/badges/image/ystyle/drone-plugin-rancher2.svg)](https://microbadger.com/images/ystyle/drone-plugin-rancher2 "Get your own image badge on microbadger.com")

# drone-plugin-rancher2
>NOTE: only can update image name

### Usage
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
