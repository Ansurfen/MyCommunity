import 'package:my_community/models/user.dart';

class UserStore {
  String? jwt;
  UserInfo? userInfo;

  void clear() {
   jwt = "" ;
   userInfo = null;
  }
}

UserStore _userStore = UserStore();

UserStore getUserStore() {
  return _userStore;
}

void releaseUserStore() {
  _userStore.clear();
}

void setJWT(String jwt) {
  _userStore.jwt = jwt;
}

String? getJWT() {
  return _userStore.jwt;
}

void setUserInfo(UserInfo userInfo) {
  _userStore.userInfo = userInfo;
}

UserInfo? getUserInfo() {
  return _userStore.userInfo;
}