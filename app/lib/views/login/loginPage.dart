import 'dart:convert';

import 'package:dio/dio.dart' as dio;
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:my_community/api/user.dart';
import 'package:my_community/models/user.dart';
import 'package:my_community/store/conf.dart';

import '../../store/user.dart';

class LoginPage extends StatefulWidget {
  const LoginPage({Key? key}) : super(key: key);

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final GlobalKey _formKey = GlobalKey<FormState>();
  late String _username, _password;
  bool _isObscure = true;
  Color _eyeColor = Colors.grey;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: Form(
        key: _formKey,
        autovalidateMode: AutovalidateMode.onUserInteraction,
        child: ListView(
          padding: const EdgeInsets.symmetric(horizontal: 20),
          children: [
            const SizedBox(height: kToolbarHeight),
            showTitle(),
            showTitleLine(),
            const SizedBox(height: 50),
            userNameInput(),
            const SizedBox(height: 30),
            passwordInput(context),
            forgetPasswordText(context),
            const SizedBox(height: 50),
            loginButton(context),
            const SizedBox(height: 30),
            registerText(context),
          ],
        ),
      ),
    );
  }

  Widget registerText(context) {
    return Center(
      child: Padding(
        padding: const EdgeInsets.only(top: 10),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text('没有账号?'),
            const SizedBox(
              width: 15,
            ),
            GestureDetector(
              child: const Text('点击注册', style: TextStyle(color: Colors.green)),
              onTap: () {
                Get.toNamed('/register');
              },
            )
          ],
        ),
      ),
    );
  }

  Widget loginButton(BuildContext context) {
    return Align(
      child: SizedBox(
        height: 45,
        width: 270,
        child: ElevatedButton(
          style: ButtonStyle(
              // 设置圆角
              shape: MaterialStateProperty.all(const StadiumBorder(
                  side: BorderSide(style: BorderStyle.none)))),
          child: Text('Login',
              style: Theme.of(context).primaryTextTheme.headline5),
          onPressed: () async {
            if ((_formKey.currentState as FormState).validate()) {
              (_formKey.currentState as FormState).save();
              User user = User(username: _username, password: _password);
              dio.Response res = await userLogin(json.encode(user.toJson()));
              if (res.statusCode == 200) {
                setJWT(res.data['data']['token']);
                getConf().loggedIn();
                res = await userInfo(getJWT()!);
                UserInfo _userInfo = UserInfo.fromJson(res.data['data']);
                setUserInfo(_userInfo);
                Get.offAllNamed('/home');
              }
            }
          },
        ),
      ),
    );
  }

  Widget forgetPasswordText(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(top: 8),
      child: Align(
        alignment: Alignment.centerRight,
        child: TextButton(
          onPressed: () {
            // Navigator.pop(context);
            print("忘记密码");
          },
          child: const Text("忘记密码？",
              style: TextStyle(fontSize: 14, color: Colors.grey)),
        ),
      ),
    );
  }

  Widget passwordInput(BuildContext context) {
    return TextFormField(
        obscureText: _isObscure, // 是否显示文字
        onSaved: (v) => _password = v!,
        validator: (v) {
          if (v!.isEmpty) {
            return '请输入密码';
          }
        },
        decoration: InputDecoration(
            labelText: "Password",
            suffixIcon: IconButton(
              icon: Icon(
                Icons.remove_red_eye,
                color: _eyeColor,
              ),
              onPressed: () {
                // 修改 state 内部变量, 且需要界面内容更新, 需要使用 setState()
                setState(() {
                  _isObscure = !_isObscure;
                  _eyeColor = (_isObscure
                      ? Colors.grey
                      : Theme.of(context).iconTheme.color)!;
                });
              },
            )));
  }

  Widget userNameInput() {
    return TextFormField(
      decoration: const InputDecoration(labelText: 'Username'),
      validator: (v) {
        var usernameReg = RegExp(r".+");
        if (!usernameReg.hasMatch(v!)) {
          return '请输入用户名';
        }
        return null;
      },
      onSaved: (v) => _username = v!,
    );
  }

  Widget showTitleLine() {
    return Padding(
        padding: const EdgeInsets.only(left: 12.0, top: 4.0),
        child: Align(
          alignment: Alignment.bottomLeft,
          child: Container(
            color: Colors.black,
            width: 40,
            height: 2,
          ),
        ));
  }

  Widget showTitle() {
    return const Padding(
        padding: EdgeInsets.all(8),
        child: Text(
          'Login',
          style: TextStyle(fontSize: 42),
        ));
  }
}
