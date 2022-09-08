import 'dart:convert';

import 'package:dio/dio.dart' as dio;
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:my_community/api/user.dart';

class RegisterPage extends StatefulWidget {
  const RegisterPage({Key? key}) : super(key: key);

  @override
  State<RegisterPage> createState() => _RegisterPageState();
}

class _RegisterPageState extends State<RegisterPage> {
  final GlobalKey _formKey = GlobalKey<FormState>();
  late String _username, _password, _key;
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
            const SizedBox(height: 30),
            keyInput(),
            forgetPasswordText(context),
            const SizedBox(height: 50),
            registerButton(context),
            const SizedBox(height: 30),
            loginText(context),
          ],
        ),
      ),
    );
  }


  Widget loginText(context) {
    return Center(
      child: Padding(
        padding: const EdgeInsets.only(top: 10),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text('注册成功?'),
            const SizedBox(
              width: 15,
            ),
            GestureDetector(
              child: const Text('点击登录', style: TextStyle(color: Colors.green)),
              onTap: () {
                Get.offNamed('/login');
              },
            )
          ],
        ),
      ),
    );
  }

  Widget registerButton(BuildContext context) {
    return Align(
      child: SizedBox(
        height: 45,
        width: 270,
        child: ElevatedButton(
          style: ButtonStyle(
              shape: MaterialStateProperty.all(const StadiumBorder(
                  side: BorderSide(style: BorderStyle.none)))),
          child: Text('Register',
              style: Theme
                  .of(context)
                  .primaryTextTheme
                  .headline5),
          onPressed: () async {
            if ((_formKey.currentState as FormState).validate()) {
              (_formKey.currentState as FormState).save();
              dio.FormData _data = dio.FormData.fromMap(
                  {"username": _username, "password": _password, "key": _key});
              dio.Response res = await userRegister(_data);
              if (res.statusCode == 200) {
                print("注册成功");
                Get.back();
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
                      : Theme
                      .of(context)
                      .iconTheme
                      .color)!;
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

  Widget keyInput() {
    return TextFormField(
      decoration: const InputDecoration(labelText: 'Key'),
      validator: (v) {
        var keyReg = RegExp(r"(\w+@\w+(\.\w+)+|\d+)");
        if (!keyReg.hasMatch(v!)) {
          return '请输入邮箱或者手机号';
        }
        return null;
      },
      onSaved: (v) => _key = v!,
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
          'Register',
          style: TextStyle(fontSize: 42),
        ));
  }
}
