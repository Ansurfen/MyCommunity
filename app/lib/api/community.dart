import 'package:dio/dio.dart';
import 'package:my_community/utils/net.dart';

Future<Response> newCommunity(String data, jwt) {
  return Service().getAuthService(jwt).post('/community/new', data: data);
}

Future<Response> searchCommunity(FormData data) {
  return Service().getService().post('/community/search', data: data);
}

Future<Response> infoCommunity(FormData data, jwt) {
  return Service().getAuthService(jwt).post('/community/new', data: data);
}

Future<Response> addCommunity(String data, jwt) {
  return Service().getAuthService(jwt).post('/community/add', data: data);
}

Future<Response> editCommunity(String data, jwt) {
  return Service().getAuthService(jwt).post('/community/edit', data: data);
}

Future<Response> addNote(String data, jwt) {
  return Service().getAuthService(jwt).post('/community/edit', data: data);
}

Future<Response> infoNote(FormData data) {
  return Service().getService().post('/community/edit', data: data);
}