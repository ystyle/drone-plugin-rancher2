# drone-plugin-rancher2
>NOTE: only can update image name

### Usage
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
