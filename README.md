# ONLINE-STORE
> A hands-on project to learn golang

<img src="https://img.shields.io/github/go-mod/go-version/stewie1520/online-store" alt="go version"/>&nbsp;

# :rocket: Quick start
```shell
# running in debug mode
docker-compose -f docker-compose.yml -f docker-compose.debug.yml up --build
```

By default, the application is running in debug mode by [delve](https://github.com/go-delve/delve) under port 8181, so you will need a delve client (which almost implemented in modern IDEs, text editors...)

# :package: Used packages
| Name                                                      | Version      | Type          |
| -------------------------------------------------------   | ------------ | ------------- |
| [gin-gonic/gin](https://github.com/gin-gonic/gin)         | 1.7.7        | Web framework |
| [kyleconroy/sqlc](https://github.com/kyleconroy/sqlc)     | 1.13.0       | Tooling       |
| [go-delve/delve](https://github.com/go-delve/delve)       | 1.8.3        | Tooling       |
| [google/wire](https://github.com/google/wire)             | 0.5.0        | Tooling       |
| [jessevdk/go-flags](https://github.com/jessevdk/go-flags) | 1.5.0        | Library       |
| [spf13/viper](https://github.com/spf13/viper)             | 1.10.0       | Library       |

