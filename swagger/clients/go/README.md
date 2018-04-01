# Go API client for swagger

The LBRY blockchain is read into SQL where important structured information can be extracted through the Chain Query API.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**AddressSummary**](docs/DefaultApi.md#addresssummary) | **Get** /AddressSummary | Returns a summary of Address activity
*DefaultApi* | [**Status**](docs/DefaultApi.md#status) | **Get** /Status | Returns important status information about Chain Query


## Documentation For Models

 - [AddressSummary](docs/AddressSummary.md)
 - [TableSize](docs/TableSize.md)
 - [TableStatus](docs/TableStatus.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Author


