import 'package:dio/dio.dart';
import 'package:flutter/foundation.dart';
import 'package:my_community/store/conf.dart';

class Service {
  final Dio _dio = Dio(BaseOptions(
    baseUrl: getConf().url,
    connectTimeout: 5000,
    receiveTimeout: 3000,
  ))..interceptors.add(Logging());

  Dio getService() {
    return _dio;
  }

  Dio getAuthService(String jwt) {
    var service = Service().getService();
    service.options.headers['Authorization'] = "Bearer " + jwt;
    return service;
  }
}

class Logging extends Interceptor {
  @override
  void onRequest(RequestOptions options, RequestInterceptorHandler handler) {
    if (kDebugMode) {
      print('REQUEST[${options.method}] => PATH: ${options.path}');
    }
    return super.onRequest(options, handler);
  }

  @override
  void onResponse(Response response, ResponseInterceptorHandler handler) {
    if (kDebugMode) {
      print(
          'RESPONSE[${response.statusCode}] => PATH: ${response.requestOptions.path}');
    }
    return super.onResponse(response, handler);
  }

  @override
  void onError(DioError err, ErrorInterceptorHandler handler) {
    if (kDebugMode) {
      print(
          'ERROR[${err.response?.statusCode}] => PATH: ${err.requestOptions.path}');
    }
    return super.onError(err, handler);
  }
}
