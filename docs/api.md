## API

使用 fiber 框架。

| 请求方法 | 路径                                                      | 用途                    |
|------|---------------------------------------------------------|-----------------------|
| GET  | `/metrics`                                              | 暴露 prometheus 采集指标的端点 |
| GET  | `/api/v1/config`                                        | 读取配置                  |
| GET  | `/v1/endpoints/:key/health/badge.svg`                   |                       |
| GET  | `/v1/endpoints/:key/health/badge.shields`               |                       |
| GET  | `/v1/endpoints/:key/uptimes/:duration`                  |                       |
| GET  | `/v1/endpoints/:key/uptimes/:duration/badge.svg`        |                       |
| GET  | `/v1/endpoints/:key/response-times/:duration`           |                       |
| GET  | `/v1/endpoints/:key/response-times/:duration/badge.svg` |                       |
| GET  | `/v1/endpoints/:key/response-times/:duration/chart.svg` |                       |
| GET  | `/v1/endpoints/:key/response-times/:duration/history`   |                       |
| GET  | `/v1/endpoints/:key/external`                           |                       |
| GET  | `/`                                                     | 首页                    |
| GET  | `/endpints/:key`                                        | 首页                    |
| GET  | `/suites/:key`                                          | 首页                    |
| GET  | `/health`                                               | 健康端点                  |
| GET  | `/css/custom.css`                                       | CSS                   |
| -    | `/index.html`                                           | 重定向到 /                |
| -    | `/`                                                     | 重定向到 index.html       |
| GET  | `/v1/endpoints/statuses`                                |                       |
| GET  | `/v1/endpoints/:key/statuses`                           |                       |
| GET  | `/v1/suites/statuses`                                   |                       |
| GET  | `/v1/suites/:key/statuses`                              |                       |