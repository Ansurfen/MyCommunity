import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:my_community/store/conf.dart';
import 'package:my_community/store/user.dart';

import '../models/user.dart';

class LoggedInDrawer extends StatelessWidget {
  LoggedInDrawer({Key? key}) : super(key: key);
  UserInfo? userInfo = getUserInfo();

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: MediaQuery.removePadding(
          context: context,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: <Widget>[
              Row(
                children: [
                  Padding(
                    padding: const EdgeInsets.only(top: 60.0,left: 10),
                    child: ClipOval(
                      child: Image.network(
                        getUserInfo()!.profile == 1
                            ? getConf().url +
                            "/images/user/" +
                            getUserInfo()!.username! +
                            ".png"
                            : 'https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png',
                        width: 80,
                      ),
                    ),
                  ),
                  Text(
                    getUserInfo()!.username!,
                    style: const TextStyle(fontWeight: FontWeight.bold),
                  ),
                ],
              ),
              Expanded(
                  child: ListView(
                children: <Widget>[
                  const ListTile(
                    leading: Icon(Icons.account_circle),
                    title: Text('个人信息'),
                  ),
                  ListTile(
                    leading: const Icon(Icons.clear),
                    title: const Text('退出登录'),
                    onTap: () {
                      getConf().unLogin();
                      Get.offAndToNamed('/home');
                    },
                  )
                ],
              ))
            ],
          )),
    );
  }
}

class UnLoginDrawer extends StatelessWidget {
  const UnLoginDrawer({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: MediaQuery.removePadding(
          context: context,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: <Widget>[
              const Padding(padding: EdgeInsets.only(top: 40)),
              ListTile(
                leading: const Icon(Icons.login),
                title: const Text('登录'),
                onTap: () {
                  Get.toNamed('/login');
                },
              ),
              ListTile(
                leading: const Icon(Icons.add),
                title: const Text('注册'),
                onTap: () {
                  Get.toNamed('/register');
                },
              ),
            ],
          )),
    );
  }
}
