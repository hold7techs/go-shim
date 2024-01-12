## Example

```go

// client config
appCfg, err := ParseAppConfig(appYamlConfigFile)
if err != nil {
    return nil, err
}

// app prompt
appPromptCfg, err := ParseAppPromptConfig(appPromptConfigFile)
    if err != nil {
    return nil, err
}

// openAI infra
openAIInfra, err := openaix.NewOpenAIHttpProxyClient(appCfg)
    if err != nil {
    return nil, err
}
```