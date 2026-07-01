## Env

| 变量                        | 描述                                                                      |
| --------------------------- | ------------------------------------------------------------------------- |
| `GATUS_DELAY_START_SECONDS` | 启动延迟，单位：秒                                                        |
| `GATUS_LOG_LEVEL`           | 日志级别，支持 DEBUG/INFO/WARN/ERROR/FATAL，默认为 DEBUG，错误设置为 INFO |
| `GATUS_CONFIG_PATH`         | 配置文件路径                                                              |
| `GATUS_CONFIG_FILE`         | 配置文件名称                                                              |
| `ENVIRONMENT`               | 由 fiber 使用，如果是 dev 的话，则开启 CROS 支持                          |
| `ROUTER_TEST`               | 测试，实际上不启动 Server                                                 |
