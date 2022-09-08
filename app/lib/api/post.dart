import 'package:dio/dio.dart';
import 'package:my_community/utils/net.dart';

Future<Response> addPost(String data, jwt) {
  return Service().getAuthService(jwt).post('/community/post/add', data: data);
}

Future<Response> delPost(String data, jwt) {
  return Service().getAuthService(jwt).post('/community/post/del', data: data);
}

Future<Response> infoPost(FormData data) {
  return Service().getService().post('/community/post/info', data: data);
}

Future<Response> infoPosts(FormData data) {
  return Service().getService().post('/post/info', data: data);
}

Future<Response> addComment(String data, jwt) {
  return Service().getAuthService(jwt).post('/post/add', data: data);
}

Future<Response> appendComment(String data, jwt) {
  return Service().getAuthService(jwt).post('/post/append', data: data);
}

Future<Response> delComment(String data, jwt) {
  return Service().getAuthService(jwt).post('/post/del', data: data);
}

Future<Response> editPost(String data, jwt) {
  return Service().getAuthService(jwt).post('/post/edit/post', data: data);
}
