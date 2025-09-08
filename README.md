# helloworld-mcp

Training repo

## Docker images

- [math](https://hub.docker.com/r/rvolykh/helloworld-mcp-math)

# Test MCP

```shell
npx @modelcontextprotocol/inspector
```

MCP Server URL: http://localhost:8080

## CI Configuration

The following secrets are required:
- `RELEASE_PLEASE_GH_TOKEN`
- `DOCKERHUB_LOGIN`
- `DOCKERHUB_TOKEN`
