import 'dart:math';

class Conf {
  String url = 'http://172.22.42.42:9090';
  bool login = false;

  bool isLogin() {
    return login;
  }

  void unLogin() {
    login = false;
  }

  void loggedIn() {
    login = true;
  }
}

Conf _conf = Conf();

Conf getConf() {
  return _conf;
}
