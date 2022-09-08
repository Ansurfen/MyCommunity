import 'package:dio/dio.dart';
import 'package:my_community/utils/net.dart';

Future<Response> passCommit(String data, jwt) {
  return Service().getAuthService(jwt).post('/commit/pass', data: data);
}

Future<Response> rejectCommit(String data, jwt) {
  return Service().getAuthService(jwt).post('/commit/reject', data: data);
}

Future<Response> communityInfoCommit(String data, jwt) {
  return Service().getAuthService(jwt).post('/commit/info', data: data);
}

Future<Response> userInfoCommit(String data, jwt) {
  return Service().getAuthService(jwt).post('/commit/info/user', data: data);
}

Future<Response> changeCommit(String data, jwt) {
  return Service().getAuthService(jwt).post('/commit/change', data: data);
}