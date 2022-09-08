import 'package:dio/dio.dart';
import 'package:my_community/utils/net.dart';

Future<Response> userLogin(String data) {
  return Service().getService().post('/user/login', data: data);
}

Future<Response> userRegister(FormData data) {
  return Service().getService().post('/user/register', data: data);
}

Future<Response> userInfo(String jwt) {
  return Service().getAuthService(jwt).get('/user/info');
}

Future<Response> userEdit(String data, jwt) {
  return Service().getAuthService(jwt).post('/user/edit/info', data: data);
}

Future<Response> imageUpdate(FormData data, String jwt) {
  var service = Service().getAuthService(jwt);
  service.options.headers['Content-Type'] = 'multipart/form-data';
  return service.post('/user/image/update', data: data);
}
